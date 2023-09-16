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

type Response struct {
	Shards struct {
		Failed     int `json:"failed"`
		Skipped    int `json:"skipped"`
		Successful int `json:"successful"`
		Total      int `json:"total"`
	} `json:"_shards"`
	Hits struct {
		Hits []struct {
			Timestamp time.Time `json:"@timestamp"`
			ID        string    `json:"_id"`
			Index     string    `json:"_index"`
			Score     float64   `json:"_score"`
			Source    struct {
				Timestamp time.Time `json:"@timestamp"`
				Bcc       []string  `json:"bcc,omitempty"`
				Body      string    `json:"body"`
				Cc        []string  `json:"cc,omitempty"`
				Date      string    `json:"date"`
				From      string    `json:"from"`
				MessageID string    `json:"message_id"`
				Subject   string    `json:"subject"`
				To        []string  `json:"to"`
			} `json:"_source"`
			Type string `json:"_type"`
		} `json:"hits"`
		MaxScore float64 `json:"max_score"`
		Total    struct {
			Value int `json:"value"`
		} `json:"total"`
	} `json:"hits"`
	TimedOut bool `json:"timed_out"`
	Took     int  `json:"took"`
}

type Document struct {
	Id      string    `json:"id"`
	Subject string    `json:"subject"`
	From    string    `json:"from"`
	To      []string  `json:"to"`
	Cc      []string  `json:"cc"`
	Body    string    `json:"body"`
	Date    time.Time `json:"date"`
}

type SearchResult struct {
	TotalFound int        `json:"total_found"`
	Documents  []Document `json:"documents"`
}
