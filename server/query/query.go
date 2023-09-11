package query

func GenerateDefaultQuerySettings() QuerySettings {
	defaultQuerySettings := QuerySettings{}
	defaultQuerySettings.Sort = "-date"
	defaultQuerySettings.Pagination = QueryPaginationSettings{
		From: 0,
		Size: 20,
	}

	return defaultQuerySettings
}
