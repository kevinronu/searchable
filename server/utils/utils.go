package utils

import (
	"bytes"
	"log"
	"net/mail"
	"os"

	"github.com/joho/godotenv"
	"github.com/kevinronu/email-indexer/server/models"
)

func FileToEmail(path string) (*models.Email, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	msg, err := mail.ReadMessage(file)
	if err != nil {
		return nil, err
	}

	email := &models.Email{
		MessageId: msg.Header.Get("Message-ID"),
		From:      msg.Header.Get("From"),
		Subject:   msg.Header.Get("Subject"),
	}

	date, err := msg.Header.Date()
	if err != nil {
		return nil, err
	}
	email.Date = date

	if msg.Header.Get("To") != "" {
		to, err := msg.Header.AddressList("To")
		if err != nil {
			return nil, err
		}
		for _, addr := range to {
			email.To = append(email.To, addr.Address)
		}
	}

	if msg.Header.Get("Cc") != "" {
		cc, err := msg.Header.AddressList("Cc")
		if err != nil {
			return nil, err
		}
		for _, addr := range cc {
			email.Cc = append(email.Cc, addr.Address)
		}
	}

	if msg.Header.Get("Bcc") != "" {
		bcc, err := msg.Header.AddressList("Bcc")
		if err != nil {
			return nil, err
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
