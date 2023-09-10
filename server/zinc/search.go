package zinc

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/kevinronu/email-indexer/server/models"
)

func GetAllDocuments(querySettings models.QuerySettings, zincCredentials models.ZincCredentials, indexName string) ([]byte, error) {
	const queryTemplate = `{
    "query": {
      "match_all": {}
    },
		"sort": ["%s"],
		"from": %d,
    "size": %d
	}`
	query := fmt.Sprintf(queryTemplate, querySettings.Sort, querySettings.Pagination.From, querySettings.Pagination.Size)
	req, err := http.NewRequest("POST", zincCredentials.BaseUrl+"/api/"+indexName+"/_search", strings.NewReader(query))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(zincCredentials.User, zincCredentials.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
