package zinc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kevinronu/email-indexer/server/query"
	"github.com/kevinronu/email-indexer/server/utils"
)

type getAllDocumentsHandler func(http.ResponseWriter, *http.Request, query.QuerySettings)

func (zincService ZincService) MiddlewareGetAllDocuments(handler getAllDocumentsHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		querySettings := query.QuerySettings{}
		err := json.NewDecoder(r.Body).Decode(&querySettings)
		if err != nil {
			fmt.Println("WARNING: Couldn't decode request body", err)
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to decode request body")
			return
		}

		if !querySettings.IsValid() {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Query Settings")
			return
		}

		handler(w, r, querySettings)
	}
}

type searchDocumentHandler func(http.ResponseWriter, *http.Request, query.SearchQuerySettings)

func (zincService ZincService) MiddlewareSearchDocument(handler searchDocumentHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		querySettings := query.SearchQuerySettings{}
		err := json.NewDecoder(r.Body).Decode(&querySettings)
		if err != nil {
			fmt.Println("WARNING: Couldn't decode request body", err)
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to decode request body")
			return
		}

		if !querySettings.IsValidSearch() {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Query Settings")
			return
		}

		handler(w, r, querySettings)
	}
}

type deleteDocumentHandler func(http.ResponseWriter, *http.Request, string)

func (zincService ZincService) MiddlewareDeleteDocument(handler deleteDocumentHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid or missing 'id' parameter in the URL")
			return
		}

		handler(w, r, id)
	}
}
