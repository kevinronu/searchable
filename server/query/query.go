package query

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

func GenerateDefaultQuerySettings() QuerySettings {
	defaultQuerySettings := QuerySettings{}
	defaultQuerySettings.Sort = "-date"
	defaultQuerySettings.Pagination = QueryPaginationSettings{
		From: 0,
		Size: 20,
	}

	return defaultQuerySettings
}
