package utils

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/kevinronu/email-indexer/server/models"
)

func TestFileToEmail(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test_email_*.eml")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	date := time.Now()
	customFormat := "Mon, 02 Jan 2006 15:04:05.999999999 -0700"
	headers := make(map[string]string)
	headers["Message-ID"] = "<test-message-id>"
	headers["From"] = "sender@example.com"
	headers["To"] = "recipient1@example.com, recipient2@example.com"
	headers["CC"] = "recipient1@example.com, recipient2@example.com"
	headers["Subject"] = "Test Email"
	// headers["Date"] = date.Format(time.RFC1123Z)
	headers["Date"] = date.Format(customFormat)

	body := "This is the body of the email."

	message := ""
	for header, value := range headers {
		message += fmt.Sprintf("%s: %s\r\n", header, value)
	}
	message += "\r\n" + body

	if _, err := tmpFile.WriteString(message); err != nil {
		t.Fatalf("Error writing to temporary file: %v", err)
	}

	email, err := FileToEmail(tmpFile.Name())
	if err != nil {
		t.Fatalf("Error in FileToEmail: %v", err)
	}

	expectedEmail := models.Email{
		MessageId: "<test-message-id>",
		// Date:      date.Truncate(time.Second),
		Date:    date,
		From:    "sender@example.com",
		Subject: "Test Email",
		To:      []string{"recipient1@example.com", "recipient2@example.com"},
		Cc:      []string{"recipient1@example.com", "recipient2@example.com"},
		Body:    "This is the body of the email.",
	}

	if email.MessageId != expectedEmail.MessageId {
		t.Errorf("Expected ID email: %v, but got: %v", expectedEmail.MessageId, email.MessageId)
	}

	if !email.Date.Equal(expectedEmail.Date) {
		t.Errorf("Expected Date email: %v, but got: %v", expectedEmail.Date, email.Date)
	}

	if email.From != expectedEmail.From {
		t.Errorf("Expected From email: %v, but got: %v", expectedEmail.From, email.From)
	}

	if email.Subject != expectedEmail.Subject {
		t.Errorf("Expected Subject email: %v, but got: %v", expectedEmail.Subject, email.Subject)
	}

	if !reflect.DeepEqual(email.To, expectedEmail.To) {
		t.Errorf("Expected To email: %v, but got: %v", expectedEmail.To, email.To)
	}

	if !reflect.DeepEqual(email.Cc, expectedEmail.Cc) {
		t.Errorf("Expected Cc email: %v, but got: %v", expectedEmail.Cc, email.Cc)
	}

	if email.Body != expectedEmail.Body {
		t.Errorf("Expected Body email: %v, but got: %v", expectedEmail.Body, email.Body)
	}

	// if !reflect.DeepEqual(email, expectedEmail) {
	// 	t.Errorf("Expected email:\n%+v\n\nBut got:\n%+v", expectedEmail, email)
	// }
}
