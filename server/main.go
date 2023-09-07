package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/kevinronu/email-indexer/server/models"
	"github.com/kevinronu/email-indexer/server/routines"
	"github.com/kevinronu/email-indexer/server/utils"
)

func main() {
	portString := utils.GetEnv("SERVER_PORT")
	emailsDir := utils.GetEnv("EMAILS_DIR")
	zincBaseUrl := utils.GetEnv("ZINC_BASEURL")
	zincPort := utils.GetEnv("ZINC_PORT")
	zincAdminUser := utils.GetEnv("ZINC_ADMIN_USER")
	zincAdminPassword := utils.GetEnv("ZINC_ADMIN_PASSWORD")
	numUploaderWorkers, _ := strconv.Atoi(utils.GetEnv("NUM_UPLOADER_WORKERS"))
	numParserWorkers, _ := strconv.Atoi(utils.GetEnv("NUM_PARSER_WORKERS"))
	bulkUploadQuantity, _ := strconv.Atoi(utils.GetEnv("BULK_UPLOAD_QUANTITY"))

	zincAuth := &models.ZincAuth{
		BaseUrl:  fmt.Sprintf("http://%s:%s", zincBaseUrl, zincPort),
		User:     zincAdminUser,
		Password: zincAdminPassword,
	}

	log.Println("INFO: starting to parse and upload emails at dir:", emailsDir)
	start := time.Now()
	routines.ParseAndUploadEmails(emailsDir, numUploaderWorkers, numParserWorkers, bulkUploadQuantity, zincAuth)
	log.Printf("INFO: finished reading and write in %v\n", time.Since(start))

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting and port %s", portString)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
