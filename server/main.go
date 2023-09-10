package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/kevinronu/email-indexer/server/models"
	"github.com/kevinronu/email-indexer/server/routines"
	"github.com/kevinronu/email-indexer/server/utils"
	"github.com/kevinronu/email-indexer/server/zinc"
)

func main() {
	enableProfiling, _ := strconv.ParseBool(utils.GetEnv("ENABLE_PROFILING"))
	profilingPortString := utils.GetEnv("PROFILING_PORT")
	if enableProfiling {
		go func() {
			log.Println(http.ListenAndServe("localhost:"+profilingPortString, nil))
		}()
	}

	portString := utils.GetEnv("SERVER_PORT")
	emailsDir := utils.GetEnv("EMAILS_DIR")
	indexName := utils.GetEnv("INDEX_NAME")
	zincHost := utils.GetEnv("ZINC_HOST")
	zincPort := utils.GetEnv("ZINC_PORT")
	zincAdminUser := utils.GetEnv("ZINC_ADMIN_USER")
	zincAdminPassword := utils.GetEnv("ZINC_ADMIN_PASSWORD")
	removeIndexIfExists, _ := strconv.ParseBool(utils.GetEnv("REMOVE_INDEX_If_EXISTS"))
	numUploaderWorkers, _ := strconv.Atoi(utils.GetEnv("NUM_UPLOADER_WORKERS"))
	numParserWorkers, _ := strconv.Atoi(utils.GetEnv("NUM_PARSER_WORKERS"))
	bulkUploadQuantity, _ := strconv.Atoi(utils.GetEnv("BULK_UPLOAD_QUANTITY"))

	zincCredentials := models.ZincCredentials{
		BaseUrl:  fmt.Sprintf("http://%s:%s", zincHost, zincPort),
		User:     zincAdminUser,
		Password: zincAdminPassword,
	}

	indexExists, err := zinc.CheckIfIndexExists(indexName, zincCredentials)
	if err != nil {
		log.Fatal("FATAL: ", err)
	}

	if indexExists {
		log.Printf("INFO: index name %s already exists.\n", indexName)
		if removeIndexIfExists {
			log.Printf("INFO: deleting %s index\n", indexName)
			if err := zinc.DeleteIndex(indexName, zincCredentials); err != nil {
				log.Fatal("FATAL: ", err)
			}
			// Update indexExists value after delete
			indexExists, err = zinc.CheckIfIndexExists(indexName, zincCredentials)
			if err != nil {
				log.Fatal("FATAL: ", err)
			}
		}
	}

	if !indexExists {
		log.Println("INFO: creating index for:", indexName)
		err = zinc.CreateIndex(indexName, zincCredentials)
		if err != nil {
			log.Fatal("FATAL: ", err)
		}
		log.Println("INFO: starting to parse and upload emails at dir:", emailsDir)
		start := time.Now()
		routines.ParseAndUploadEmails(emailsDir, numUploaderWorkers, numParserWorkers, bulkUploadQuantity, zincCredentials)
		log.Printf("INFO: file indexing finished in %v\n", time.Since(start))
	}

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

	v1Router.Get("/emails", handlerReadiness)
	v1Router.Delete("/emails", handlerReadiness)
	// v1Router.Post("/emails", handlerReadiness)

	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting and port %s", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
