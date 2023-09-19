package query

import "fmt"

func (q QuerySettings) IsValid() bool {
	if q.Sort == "" || q.Pagination.From < 0 || q.Pagination.Size < 0 {
		return false
	}
	return true
}

func (q SearchQuerySettings) IsValidSearch() bool {
	if q.Query == "" || q.Sort == "" || q.Pagination.From < 0 || q.Pagination.Size < 0 {
		return false
	}
	return true
}

func (q QuerySettings) GenerateForDocumentsPost() string {
	const queryTemplate = `{
    "query": {
      "match_all": {}
    },
		"_source": [],
		"sort": ["%s"],
		"from": %d,
    "size": %d
	}`

	return fmt.Sprintf(queryTemplate, q.Sort, q.Pagination.From, q.Pagination.Size)
}

func (q SearchQuerySettings) GenerateForSearchDocumentPost() string {
	const queryTemplate = `{
		"query": {
			"bool": {
				"must": [
					{
						"query_string": {
							"query": "%s"
						}
					}
				]%s
			}
		},
		"_source": [],
		"sort": ["%s"],
		"from": %d,
		"size": %d
	}`

	dateRangeFilter := ""
	if q.DateRange != nil {
		dateRangeFilter = fmt.Sprintf(`,
		"filter": [
			{
				"range": {
					"date": {
						"gte": "%v",
						"lte": "%v",
						"format": "2006-01-02 15:04:05 -0700 MST",
						"time_zone": "UTC"
					}
				}
			}
		]`, q.DateRange.From, q.DateRange.To)
	}

	return fmt.Sprintf(queryTemplate, q.Query, dateRangeFilter, q.Sort, q.Pagination.From, q.Pagination.Size)
}
