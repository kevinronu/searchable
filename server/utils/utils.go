package utils

import (
	"bytes"
	"fmt"
	"log"
	"net/mail"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kevinronu/email-indexer/server/models"
)

func FileToEmail(path string) (models.Email, error) {
	file, err := os.Open(path)
	if err != nil {
		return models.Email{}, err
	}
	defer file.Close()

	msg, err := mail.ReadMessage(file)
	if err != nil {
		return models.Email{}, err
	}

	email := models.Email{
		MessageId: msg.Header.Get("Message-ID"),
		From:      msg.Header.Get("From"),
		Subject:   msg.Header.Get("Subject"),
	}

	date, err := msg.Header.Date()
	if err != nil {
		return models.Email{}, err
	}
	email.Date = date

	if msg.Header.Get("To") != "" {
		to, err := msg.Header.AddressList("To")
		if err != nil {
			return models.Email{}, err
		}
		for _, addr := range to {
			email.To = append(email.To, addr.Address)
		}
	}

	if msg.Header.Get("Cc") != "" {
		cc, err := msg.Header.AddressList("Cc")
		if err != nil {
			return models.Email{}, err
		}
		for _, addr := range cc {
			email.Cc = append(email.Cc, addr.Address)
		}
	}

	if msg.Header.Get("Bcc") != "" {
		bcc, err := msg.Header.AddressList("Bcc")
		if err != nil {
			return models.Email{}, err
		}
		for _, addr := range bcc {
			email.Bcc = append(email.Bcc, addr.Address)
		}
	}

	// body, err := io.ReadAll(msg.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// email.Body = string(body)
	buf := new(bytes.Buffer)
	buf.ReadFrom(msg.Body)
	email.Body = buf.String()

	return email, nil
}

func GetEnv(key string) string {
	godotenv.Load(".env")

	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("FATAL: %s environment variable is not set", key)
	}

	return value
}

func ConvertResponseToDocuments(response models.Response) []models.Document {
	documents := make([]models.Document, 0, len(response.Hits.Hits))

	for _, hit := range response.Hits.Hits {
		date, err := time.Parse(time.RFC3339, hit.Source.Date)
		if err != nil {
			fmt.Printf("WARNING: Couldn't parse Date for with ID %s: %s", hit.ID, err)
			date = time.Now()
		}

		document := models.Document{
			Id:      hit.ID,
			Subject: hit.Source.Subject,
			From:    hit.Source.From,
			To:      hit.Source.To,
			Cc:      hit.Source.Cc,
			Body:    hit.Source.Body,
			Date:    date,
		}

		documents = append(documents, document)
	}

	return documents
}
