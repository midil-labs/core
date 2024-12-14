package schemas

type PaginationSettings struct {
	Limit    int64           `json:"limit"`
	Skip     int64           `json:"skip"`
	Sort     map[string]int  `json:"sort"`
	PageSize int64           `json:"page_size"`
}

type ResourceFilter struct {
	Filter map[string]interface{} `json:"filter"`
}

type ResourceQuery struct {
	*ResourceFilter 		  `json:",inline"`	
	*PaginationSettings       `json:",inline"`
}
