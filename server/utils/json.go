package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(responseWriter http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with %d error: %s", code, msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJSON(responseWriter, code, errResponse{
		Error: msg,
	})
}

func RespondWithJSON(responseWriter http.ResponseWriter, code int, payload interface{}) {
	serializedPayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		responseWriter.WriteHeader(500)
		return
	}

	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(code)
	responseWriter.Write(serializedPayload)
}
