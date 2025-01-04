package request

import "github.com/midil-labs/core/shared/dtos/common"

type Filter struct {
	Fields
}

type Sort struct {
	Fields []string
}

type PaginationQuery struct {
	PageSize   int
	PageNumber int
}

// Fields is a map of fields to include in the response.
type Fields map[string][]string


// Include is a slice of strings that represent the relationships to include in the response.
type Include []string


// QueryParams is a struct that holds all the query parameters that can be passed to the API.
type QueryParams struct {
	Filter  *Filter
	Sort    *Sort
	Page    *PaginationQuery
	Fields  *Fields
	Include *Include
}

// ResourceRequest is a JSON:API resource object for inbound requests.
type Resource struct {
	common.ResourceIdentifier
	LID           *common.ID                     `json:"id,omitempty"`
	Attributes    map[string]interface{}         `json:"attributes,omitempty"`
	Relationships map[string]common.Relationship `json:"relationships,omitempty"`
}

// JSONAPIRequest top-level. Usually, you either have single resource or multiple resources in "data".
type JSONAPIRequest struct {
	Data *Resource `json:"data,omitempty"`
}
