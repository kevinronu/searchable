package zinc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/kevinronu/email-indexer/server/models"
	"github.com/kevinronu/email-indexer/server/query"
	"github.com/kevinronu/email-indexer/server/utils"
)

func (zincService ZincService) HandlerDocumentsGet(w http.ResponseWriter, r *http.Request, id string) {
	const queryTemplate = `{
    "query": {
      "term": {
				"_id": "%s"
			}
    },
		"_source": []
	}`
	query := fmt.Sprintf(queryTemplate, id)

	req, err := http.NewRequest("POST", zincService.BaseUrl+"/es/"+zincService.IndexName+"/_search", strings.NewReader(query))
	if err != nil {
		fmt.Println("WARNING: Couldn't create a request:", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create a request")
		return
	}
	req.SetBasicAuth(zincService.User, zincService.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("WARNING: Couldn't send request to ZincSearch", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to send a request to ZincSearch")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("WARNING: ZincSearch for id: %s, responded with code %v and can't read response body: %v\n", id, resp.StatusCode, err)
			utils.RespondWithError(w, resp.StatusCode, "An error occurred while processing the request")
		} else {
			fmt.Printf("WARNING: ZincSearch for id: %s, responded with code %v: %v\n", id, resp.StatusCode, string(body))
			utils.RespondWithError(w, resp.StatusCode, string(body))
		}
		return
	}

	var responseJSON models.Response
	// var responseJSON map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseJSON)
	if err != nil {
		fmt.Println("WARNING: Couldn't decode response body from ZincSearch", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to decode response body from ZincSearch")
		return
	}

	documents := utils.ConvertResponseToDocuments(responseJSON)
	searchResult := models.SearchResult{
		TotalFound: responseJSON.Hits.Total.Value,
		Documents:  documents,
	}
	utils.RespondWithJSON(w, http.StatusOK, searchResult)
}

func (zincService ZincService) HandlerDocumentsDelete(w http.ResponseWriter, r *http.Request, id string) {
	req, err := http.NewRequest("DELETE", zincService.BaseUrl+"/es/"+zincService.IndexName+"/_doc/"+id, nil)
	if err != nil {
		fmt.Println("WARNING: Couldn't create a request:", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create a request")
		return
	}
	req.SetBasicAuth(zincService.User, zincService.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("WARNING: Couldn't send request to ZincSearch", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to send a request to ZincSearch")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("WARNING: ZincSearch for id: %s, responded with code %v and can't read response body: %v\n", id, resp.StatusCode, err)
			utils.RespondWithError(w, resp.StatusCode, "An error occurred while processing the request")
		} else {
			fmt.Printf("WARNING: ZincSearch for id: %s, responded with code %v: %v\n", id, resp.StatusCode, string(body))
			utils.RespondWithError(w, resp.StatusCode, string(body))
		}
		return
	}

	responseJSON := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&responseJSON)
	if err != nil {
		fmt.Println("WARNING: Couldn't decode response body from ZincSearch", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to decode response body from ZincSearch")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, responseJSON)
}

func (zincService ZincService) HandlerDocumentsPost(w http.ResponseWriter, r *http.Request, querySettings query.QuerySettings) {
	query := querySettings.GenerateForDocumentsPost()

	req, err := http.NewRequest("POST", zincService.BaseUrl+"/es/"+zincService.IndexName+"/_search", strings.NewReader(query))
	if err != nil {
		fmt.Println("WARNING: Couldn't create a request:", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create a request")
		return
	}
	req.SetBasicAuth(zincService.User, zincService.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("WARNING: Couldn't send request to ZincSearch", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to send a request to ZincSearch")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("WARNING: ZincSearch responded with code %v and can't read response body: %v\n", resp.StatusCode, err)
			utils.RespondWithError(w, resp.StatusCode, "An error occurred while processing the request")
		} else {
			fmt.Printf("WARNING: ZincSearch responded with code %v: %v\n", resp.StatusCode, string(body))
			utils.RespondWithError(w, resp.StatusCode, string(body))
		}
		return
	}

	var responseJSON models.Response
	err = json.NewDecoder(resp.Body).Decode(&responseJSON)
	if err != nil {
		fmt.Println("WARNING: Couldn't decode response body from ZincSearch", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to decode response body from ZincSearch")
		return
	}

	documents := utils.ConvertResponseToDocuments(responseJSON)
	searchResult := models.SearchResult{
		TotalFound: responseJSON.Hits.Total.Value,
		Documents:  documents,
	}
	utils.RespondWithJSON(w, http.StatusOK, searchResult)
}

func (zincService ZincService) HandlerSearchDocumentPost(w http.ResponseWriter, r *http.Request, searchQuerySettings query.SearchQuerySettings) {
	query := searchQuerySettings.GenerateForSearchDocumentPost()

	req, err := http.NewRequest("POST", zincService.BaseUrl+"/es/"+zincService.IndexName+"/_search", strings.NewReader(query))
	if err != nil {
		fmt.Println("WARNING: Couldn't create a request:", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create a request")
		return
	}
	req.SetBasicAuth(zincService.User, zincService.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("WARNING: Couldn't send request to ZincSearch", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to send a request to ZincSearch")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("WARNING: ZincSearch responded with code %v and can't read response body: %v\n", resp.StatusCode, err)
			utils.RespondWithError(w, resp.StatusCode, "An error occurred while processing the request")
		} else {
			fmt.Printf("WARNING: ZincSearch responded with code %v: %v\n", resp.StatusCode, string(body))
			utils.RespondWithError(w, resp.StatusCode, string(body))
		}
		return
	}

	var responseJSON models.Response
	// var responseJSON map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseJSON)
	if err != nil {
		fmt.Println("WARNING: Couldn't decode response body from ZincSearch", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to decode response body from ZincSearch")
		return
	}

	documents := utils.ConvertResponseToDocuments(responseJSON)
	searchResult := models.SearchResult{
		TotalFound: responseJSON.Hits.Total.Value,
		Documents:  documents,
	}

	utils.RespondWithJSON(w, http.StatusOK, searchResult)
}
