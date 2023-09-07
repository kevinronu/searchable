package zinc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kevinronu/email-indexer/server/models"
)

func UploadEmails(bulk *models.BulkEmails, auth *models.ZincAuth) error {
	const uploadEndpoint = "/api/_bulkv2"

	jsonBytes, err := json.Marshal(*bulk)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", auth.BaseUrl+uploadEndpoint, bytes.NewReader(jsonBytes))
	if err != nil {
		return err
	}
	req.SetBasicAuth(auth.User, auth.Password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return fmt.Errorf("ZincSearch replied with code %v: %v", resp.StatusCode, string(body))
	}

	return nil
}
