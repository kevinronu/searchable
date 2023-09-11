package main

import (
	"net/http"

	"github.com/kevinronu/email-indexer/server/utils"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, map[string]string{"status": "ok"})
}

func handlerErr(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 400, "Something went wrong")
}
