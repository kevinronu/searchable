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

// For Zinc

type SearchQueryZinc struct {
	Query  *QueryZinc `json:"query,omitempty"`
	Source []string   `json:"_source"`
	Sort   []string   `json:"sort,omitempty"`
	From   int        `json:"from,omitempty"`
	Size   int        `json:"size,omitempty"`
}

type QueryZinc struct {
	QueryString *QueryStringZinc `json:"query_string,omitempty"`
	Range       *RangeZinc       `json:"range,omitempty"`
	MatchAll    *MatchAll        `json:"match_all,omitempty"`
	Bool        *BoolZinc        `json:"bool,omitempty"`
	Term        *TermZinc        `json:"term,omitempty"`
}

type MatchAll struct {
}

type QueryStringZinc struct {
	Query string `json:"query"`
}

type RangeZinc struct {
	Date DateZinc `json:"date"`
}

type DateZinc struct {
	Gte      time.Time `json:"gte,omitempty"`
	Lte      time.Time `json:"lte,omitempty"`
	Format   string    `json:"format"`
	TimeZone string    `json:"time_zone"`
}

type BoolZinc struct {
	Must   []QueryZinc  `json:"must"`
	Filter *[]QueryZinc `json:"filter,omitempty"`
}

type TermZinc struct {
	Id string `json:"_id"`
}
