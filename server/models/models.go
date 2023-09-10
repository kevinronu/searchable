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

type MappingProperty struct {
	Type          string `json:"type"`
	Index         bool   `json:"index"`
	Store         bool   `json:"store"`
	Sortable      bool   `json:"sortable"`
	Aggregatable  bool   `json:"aggregatable"`
	Highlightable bool   `json:"highlightable"`
}

type DateMappingProperty struct {
	Type          string `json:"type"`
	Format        string `json:"format"`
	Index         bool   `json:"index"`
	Store         bool   `json:"store"`
	Sortable      bool   `json:"sortable"`
	Aggregatable  bool   `json:"aggregatable"`
	Highlightable bool   `json:"highlightable"`
}

type MappingProperties struct {
	MessageID MappingProperty     `json:"message_id"`
	Date      DateMappingProperty `json:"date"`
	From      MappingProperty     `json:"from"`
	To        MappingProperty     `json:"to"`
	CC        MappingProperty     `json:"cc"`
	BCC       MappingProperty     `json:"bcc"`
	Subject   MappingProperty     `json:"subject"`
	Body      MappingProperty     `json:"body"`
}

type Properties struct {
	Properties MappingProperties `json:"properties"`
}

type IndexMapping struct {
	Name        string     `json:"name"`
	StorageType string     `json:"storage_type"`
	Mappings    Properties `json:"mappings"`
}

type QueryPaginationSettings struct {
	From int // Default: 0
	Size int // Default: 20
}

type QuerySettings struct {
	Sort       string                   // Default: "-date"
	Pagination *QueryPaginationSettings // Default: {From: 0, Size: 20}
}
