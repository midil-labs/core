package request

import (
	"github.com/midil-labs/core/shared/jsonapi/common"
)

// Page holds pagination data for queries.
type Page struct {
	Size   int `json:"page[size]"`
	Number int `json:"page[number]"`
}

// Filter represents filter criteria, mapping field names to their filter values.
type Filter map[string][]string

// Fields represents sparse fieldsets, mapping resource types to their requested fields.
type Fields map[string][]string

// Include represents relationships to include in the response.
type Include []string

// Sort represents sorting criteria as a slice of field names, with optional direction prefixes.
type Sort []string

// QueryParams holds all query parameters that can be passed to the API.
type Query struct {
	Filter  Filter  `json:"filter,omitempty"`
	Sort    Sort    `json:"sort,omitempty"`
	Page    *Page   `json:"page,omitempty"`
	Fields  Fields  `json:"fields,omitempty"`
	Include Include `json:"include,omitempty"`
}

// Resource is a JSON:API resource object for inbound requests.
type Resource struct {
	common.ResourceIdentifier
	LID           *common.ID                     `json:"id,omitempty"`
	Attributes    map[string]interface{}         `json:"attributes,omitempty"`
	Relationships map[string]common.Relationship `json:"relationships,omitempty"`
}

// ListResource is a slice of Resource objects.
type ListResource = []Resource


// RequestType is a type constraint that allows either *Resource or ListResource.
type RequestType interface {
	~*Resource | ~ListResource
}


// Header is a map of HTTP headers.
type Header map[string][]string

// JSONAPIRequest is the top-level request structure, containing either a single resource or multiple resources in "data".
type JSONAPIRequest[T RequestType] struct {
	Path   string `json:"-"`
	Method string `json:"-"`
	Body   T      `json:"data"`
	Query  *Query `json:"-"`
	Header Header `json:"-"`
}
