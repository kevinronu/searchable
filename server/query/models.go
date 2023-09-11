package query

type QueryPaginationSettings struct {
	From int
	Size int
}

type QuerySettings struct {
	Sort       string
	Pagination QueryPaginationSettings
}

type SearchQuerySettings struct {
	Query string
	QuerySettings
}
