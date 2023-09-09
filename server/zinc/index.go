package zinc

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/kevinronu/email-indexer/server/models"
)

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
