package response

import (
	error "github.com/midil-labs/core/shared/dtos/error"
	"github.com/midil-labs/core/shared/dtos/common"
)

// PaginationLinks is a map of pagination links for the response. The keys are the link relation type.
type PaginationLinks struct {
	Self  string `json:"self,omitempty"`
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
}


type Pagination struct {
	CurrentPage int64 		`json:"current_page"`
	PrevPage int64 			`json:"prev_page"`
	NextPage int64 			`json:"next_page"`
	TotalPages int64 		`json:"total_pages"`
	TotalCount int64 		`json:"total_count"`
}


//Resource is a JSON:API resource object. It contains the resource identifier, attributes, relationships, links, and meta-information.
//The attributes field is an interface{} type to allow for any type of data to be passed in. 
//The relationships field is a map of relationship objects.
//The links field is a map of links objects. 
//The meta field is a map of non-standard meta-information.
type Resource struct {
	common.ResourceIdentifier
	Attributes    map[string]interface{}                   `json:"attributes,omitempty"`
	Relationships map[string]Relationship 		`json:"relationships,omitempty"`
	Links         *common.Links                 		`json:"links,omitempty"`
	Meta          common.NonStandardMeta        		`json:"meta,omitempty"`
}

type RelationshipData struct {
	Resource   *common.ResourceIdentifier
	Resources []common.ResourceIdentifier
}


type Relationship struct {
	Data  RelationshipData 			`json:"data"`
	Links *common.Links            	`json:"links,omitempty"`
	Meta  common.NonStandardMeta    `json:"meta,omitempty"`
}

type ListResource = []Resource

type DataType interface {
    ~*Resource | ~ListResource
}

type ErrorResponse struct {
	Errors []error.ErrorObject   			`json:"errors" validate:"required"`
	Meta   common.NonStandardMeta 			`json:"meta,omitempty"`
}


// JSONAPIResponse is a generic JSON:API response object. It contains the data, errors, meta-information, pagination links, and included resources.
type JSONAPIResponse[T DataType] struct {
	Data   		T               				`json:"data"`
	Meta    	common.NonStandardMeta      	`json:"meta,omitempty"`
	Links   	*PaginationLinks     			`json:"links,omitempty"`
	Included 	[]Resource 						`json:"included,omitempty"`
}
