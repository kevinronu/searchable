package models

import "time"

type Email struct {
	MessageId string    `json:"message_id"`
	Date      time.Time `json:"date"`
	From      string    `json:"from"`
	To        []string  `json:"to"`
	Cc        []string  `json:"cc"`
	Bcc       []string  `json:"bcc"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
}

type BulkEmails struct {
	Index   string  `json:"index"`
	Records []Email `json:"records"`
}

type ZincCredentials struct {
	BaseUrl  string
	User     string
	Password string
}
