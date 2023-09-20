package query

import "time"

type QueryPaginationSettings struct {
	From int `json:"from"`
	Size int `json:"size"`
}

type QueryDateRangeSettings struct {
	From time.Time `json:"from,omitempty"`
	To   time.Time `json:"to,omitempty"`
}

type QuerySettings struct {
	Sort       string                  `json:"sort"`
	Pagination QueryPaginationSettings `json:"pagination"`
	DateRange  *QueryDateRangeSettings `json:"date_range,omitempty"`
}

type SearchQuerySettings struct {
	Query string `json:"query"`
	QuerySettings
}
