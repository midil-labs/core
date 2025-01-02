package request

type Filter struct {
	Fields map[string][]string 
}

type Sort struct {
	Fields []string
}

type PaginationQuery struct {
	PageSize   int
	PageNumber int
}

type Fields map[string][]string

type Include []string

type QueryParams struct {
	Filter     Filter
	Sort       Sort
	Page       PaginationQuery
	Fields     Fields
	Include    Include
}