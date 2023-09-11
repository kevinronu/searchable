package zinc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kevinronu/email-indexer/server/query"
	"github.com/kevinronu/email-indexer/server/utils"
)

type queryHandler func(http.ResponseWriter, *http.Request, query.QuerySettings)

func (zincService ZincService) MiddlewareQuery(handler queryHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		querySettings := query.QuerySettings{}
		err := json.NewDecoder(r.Body).Decode(&querySettings)
		if err != nil {
			fmt.Println("WARNING: Couldn't decode request body", err)
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to decode request body")
			return
		}

		handler(w, r, querySettings)
	}
}
