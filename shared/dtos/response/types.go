package response

import (
	"github.com/midil-labs/core/shared/dtos"
	jsonApiError "github.com/midil-labs/core/shared/dtos/error"
)


type PaginationLinks struct {
	Self  string `json:"self,omitempty"`
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
}

type NonStandardMeta map[string]interface{}

type Pagination struct {
	CurrentPage int64 		`json:"current_page"`
	PrevPage int64 			`json:"prev_page"`
	NextPage int64 			`json:"next_page"`
	TotalPages int64 		`json:"total_pages"`
	TotalCount int64 		`json:"total_count"`
}

type RelatedLink struct {
	Href        string            `json:"href,omitempty"`
	Title       string            `json:"title,omitempty"`
	DescribedBy string            `json:"describedby,omitempty"`
	Meta        NonStandardMeta   `json:"meta,omitempty"`
}

type Links struct {
	Self    string      		`json:"self,omitempty"`
	Related *RelatedLink 		`json:"related,omitempty"`
}

type ResourceIdentifier struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Resource[T dtos.DTOInterface] struct {
	ResourceIdentifier
	Attributes    T                      		`json:"attributes,omitempty"`
	Relationships map[string]Relationship 		`json:"relationships,omitempty"`
	Links         *Links                 		`json:"links,omitempty"`
	Meta          NonStandardMeta        		`json:"meta,omitempty"`
}

type RelationshipData struct {
	Resource   *ResourceIdentifier
	Resources []ResourceIdentifier
}


type Relationship struct {
	Data  RelationshipData 		`json:"data"`
	Links *Links            	`json:"links,omitempty"`
	Meta  NonStandardMeta    	`json:"meta,omitempty"`
}

type SingleResourceResponse[T dtos.DTOInterface] struct {
	Data    *Resource[T]    					`json:"data,omitempty"`
	Meta    NonStandardMeta           			`json:"meta,omitempty"`
	Included []Resource[dtos.DTOInterface] 		`json:"included,omitempty"`
}


type MultipleResourcesResponse[T dtos.DTOInterface] struct {
	Data    []Resource[T]   						`json:"data,omitempty"`
	Links   *PaginationLinks 						`json:"links,omitempty"`
	Meta    NonStandardMeta            			    `json:"meta,omitempty"`
	Included []Resource[dtos.DTOInterface] 			`json:"included,omitempty"`
}


type ResourceResponse[T dtos.DTOInterface] struct {
	resource   *Resource[T] 
	resources []Resource[T]
	Links    *PaginationLinks     		   `json:"links,omitempty"`
	Meta     NonStandardMeta      		   `json:"meta,omitempty"`
	Included []Resource[dtos.DTOInterface] `json:"included,omitempty"`
}

type ErrorResponse struct {
	Errors []jsonApiError.ErrorObject   `json:"errors" validate:"required"`
	Meta   NonStandardMeta 				`json:"meta,omitempty"`
}



