package zinc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/kevinronu/email-indexer/server/query"
	"github.com/kevinronu/email-indexer/server/utils"
)

func (zincService ZincService) HandlerDocumentsGet(w http.ResponseWriter, r *http.Request) {
	querySettings := query.GenerateDefaultQuerySettings()
	const queryTemplate = `{
    "query": {
      "match_all": {}
    },
		"sort": ["%s"],
		"from": %d,
    "size": %d
	}`
	query := fmt.Sprintf(queryTemplate, querySettings.Sort, querySettings.Pagination.From, querySettings.Pagination.Size)
	req, err := http.NewRequest("POST", zincService.BaseUrl+"/es/"+zincService.IndexName+"/_search", strings.NewReader(query))
	if err != nil {
		fmt.Println("WARNING: Couldn't create a request:", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create a request")
		return
	}
	req.SetBasicAuth(zincService.User, zincService.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("WARNING: Couldn't send request", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to send a request")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("WARNING: Received non-200 status code:", resp.Status)
		utils.RespondWithError(w, resp.StatusCode, "Non-200 status code received")
		return
	}

	var responseJSON map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseJSON)
	if err != nil {
		fmt.Println("WARNING: Couldn't decode response body", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to decode response body")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, responseJSON)
}

func (zincService ZincService) HandlerDocumentsPost(w http.ResponseWriter, r *http.Request, querySettings query.QuerySettings) {
	const queryTemplate = `{
    "query": {
      "match_all": {}
    },
		"sort": ["%s"],
		"from": %d,
    "size": %d
	}`
	query := fmt.Sprintf(queryTemplate, querySettings.Sort, querySettings.Pagination.From, querySettings.Pagination.Size)
	req, err := http.NewRequest("POST", zincService.BaseUrl+"/es/"+zincService.IndexName+"/_search", strings.NewReader(query))
	if err != nil {
		fmt.Println("WARNING: Couldn't create a request:", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create a request")
		return
	}
	req.SetBasicAuth(zincService.User, zincService.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("WARNING: Couldn't send request", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to send a request")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("WARNING: Received non-200 status code:", resp.Status)
		utils.RespondWithError(w, resp.StatusCode, "Non-200 status code received")
		return
	}

	responseJSON := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&responseJSON)
	if err != nil {
		fmt.Println("WARNING: Couldn't decode response body", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to decode response body")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, responseJSON)
}
