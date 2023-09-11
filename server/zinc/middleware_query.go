package zinc

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (zincService ZincService) MiddlewareSearchDocuments(handler searchDocumentHandler) http.HandlerFunc {
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
