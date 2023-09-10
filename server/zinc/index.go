package zinc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/kevinronu/email-indexer/server/models"
)

func CreateIndex(indexName string, zincCredentials models.ZincCredentials) error {
	indexMapping := models.IndexMapping{
		Name:        indexName,
		StorageType: "disk",
		Mappings: models.Properties{
			Properties: models.MappingProperties{
				MessageID: models.MappingProperty{
					Type:          "keyword",
					Index:         true,
					Store:         true,
					Sortable:      false,
					Aggregatable:  false,
					Highlightable: false,
				},
				Date: models.DateMappingProperty{
					Type:          "date",
					Format:        "2006-01-02T15:04:05Z07:00",
					Index:         true,
					Store:         false,
					Sortable:      true,
					Aggregatable:  true,
					Highlightable: false,
				},
				From: models.MappingProperty{
					Type:          "keyword",
					Index:         true,
					Store:         false,
					Sortable:      false,
					Aggregatable:  false,
					Highlightable: false,
				},
				To: models.MappingProperty{
					Type:          "keyword",
					Index:         true,
					Store:         false,
					Sortable:      false,
					Aggregatable:  false,
					Highlightable: false,
				},
				CC: models.MappingProperty{
					Type:          "keyword",
					Index:         true,
					Store:         false,
					Sortable:      false,
					Aggregatable:  false,
					Highlightable: false,
				},
				BCC: models.MappingProperty{
					Type:          "keyword",
					Index:         true,
					Store:         false,
					Sortable:      false,
					Aggregatable:  false,
					Highlightable: false,
				},
				Subject: models.MappingProperty{
					Type:          "text",
					Index:         true,
					Store:         false,
					Sortable:      false,
					Aggregatable:  false,
					Highlightable: false,
				},
				Body: models.MappingProperty{
					Type:          "text",
					Index:         true,
					Store:         false,
					Sortable:      false,
					Aggregatable:  false,
					Highlightable: false,
				},
			},
		},
	}

	jsonIndexMapping, err := json.Marshal(indexMapping)
	if err != nil {
		return fmt.Errorf("Error marshaling indexMapping: %s", err)
	}

	req, err := http.NewRequest("POST", zincCredentials.BaseUrl+"/api/index", bytes.NewReader(jsonIndexMapping))
	if err != nil {
		return err
	}
	req.SetBasicAuth(zincCredentials.User, zincCredentials.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == 401 {
			return errors.New("Unauthorized, invalid credentials")
		}
		return errors.New(string(body))
	}

	return nil
}

func DeleteIndex(indexName string, zincCredentials models.ZincCredentials) error {
	req, err := http.NewRequest("DELETE", zincCredentials.BaseUrl+"/api/index/"+indexName, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(zincCredentials.User, zincCredentials.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ZincSearch responded with code %v: %v", resp.StatusCode, string(body))
	}

	return nil
}

func CheckIfIndexExists(indexName string, zincCredentials models.ZincCredentials) (bool, error) {
	req, err := http.NewRequest("HEAD", zincCredentials.BaseUrl+"/api/index/"+indexName, nil)
	if err != nil {
		return false, err
	}
	req.SetBasicAuth(zincCredentials.User, zincCredentials.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == 401 {
			return false, errors.New("Unauthorized, invalid credentials")
		}
		fmt.Println(resp)
		return false, nil
	}

	return true, nil
}
