package zinc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (zincService ZincService) CreateIndex() error {
	indexMapping := IndexMapping{
		Name:        zincService.IndexName,
		StorageType: "disk",
		Mappings: Properties{
			Properties: MappingProperties{
				MessageID: MappingProperty{
					Type:          "keyword",
					Index:         true,
					Store:         true,
					Sortable:      false,
					Aggregatable:  false,
					Highlightable: false,
				},
				Date: DateMappingProperty{
					Type:          "date",
					Format:        "2006-01-02T15:04:05Z07:00",
					Index:         true,
					Store:         false,
					Sortable:      true,
					Aggregatable:  true,
					Highlightable: false,
				},
				From: MappingProperty{
					Type:          "keyword",
					Index:         true,
					Store:         false,
					Sortable:      false,
					Aggregatable:  false,
					Highlightable: false,
				},
				To: MappingProperty{
					Type:          "keyword",
					Index:         true,
					Store:         false,
					Sortable:      false,
					Aggregatable:  false,
					Highlightable: false,
				},
				CC: MappingProperty{
					Type:          "keyword",
					Index:         true,
					Store:         false,
					Sortable:      false,
					Aggregatable:  false,
					Highlightable: false,
				},
				BCC: MappingProperty{
					Type:          "keyword",
					Index:         true,
					Store:         false,
					Sortable:      false,
					Aggregatable:  false,
					Highlightable: false,
				},
				Subject: MappingProperty{
					Type:          "text",
					Index:         true,
					Store:         false,
					Sortable:      false,
					Aggregatable:  false,
					Highlightable: false,
				},
				Body: MappingProperty{
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

	req, err := http.NewRequest("POST", zincService.BaseUrl+"/api/index", bytes.NewReader(jsonIndexMapping))
	if err != nil {
		return err
	}
	req.SetBasicAuth(zincService.User, zincService.Password)

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

func (zincService ZincService) DeleteIndex() error {
	req, err := http.NewRequest("DELETE", zincService.BaseUrl+"/api/index/"+zincService.IndexName, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(zincService.User, zincService.Password)

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

func (zincService ZincService) CheckIfIndexExists() (bool, error) {
	req, err := http.NewRequest("HEAD", zincService.BaseUrl+"/api/index/"+zincService.IndexName, nil)
	if err != nil {
		return false, err
	}
	req.SetBasicAuth(zincService.User, zincService.Password)

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
