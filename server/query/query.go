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

func (q QuerySettings) GenerateForDocumentsPost() SearchQueryZinc {
	searchQuery := SearchQueryZinc{
		Query: &QueryZinc{
			MatchAll: &MatchAll{},
		},
		Source: []string{},
		Sort:   []string{q.Sort},
		From:   q.Pagination.From,
		Size:   q.Pagination.Size,
	}

	return searchQuery
}

func (q SearchQuerySettings) GenerateForSearchDocumentPost() SearchQueryZinc {
	var filter *[]QueryZinc
	if q.DateRange != nil {
		filter = &[]QueryZinc{
			{
				Range: &RangeZinc{
					Date: DateZinc{
						Gte:      q.DateRange.From,
						Lte:      q.DateRange.To,
						Format:   "2006-01-02T15:04:05Z",
						TimeZone: "UTC",
					},
				},
			},
		}
	}

	searchQuery := SearchQueryZinc{
		Query: &QueryZinc{
			Bool: &BoolZinc{
				Must: []QueryZinc{
					{
						QueryString: &QueryStringZinc{
							Query: q.Query,
						},
					},
				},
				Filter: filter,
			},
		},
		Source: []string{},
		Sort:   []string{q.Sort},
		From:   q.Pagination.From,
		Size:   q.Pagination.Size,
	}

	return searchQuery
}
