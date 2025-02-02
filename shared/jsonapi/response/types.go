package response

import (
	"github.com/midil-labs/core/shared/jsonapi/common"
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
	CurrentPage int64 `json:"current_page"`
	PrevPage    int64 `json:"prev_page"`
	NextPage    int64 `json:"next_page"`
	TotalPages  int64 `json:"total_pages"`
	TotalCount  int64 `json:"total_count"`
}

// Data is a JSON:API resource object. It contains the resource identifier, attributes, relationships, links, and meta-information.
// The attributes field is an interface{} type to allow for any type of data to be passed in.
// The relationships field is a map of relationship objects.
// The links field is a map of links objects.
// The meta field is a map of non-standard meta-information.
type Data struct {
	common.ResourceIdentifier
	Attributes    map[string]interface{}          `json:"attributes,omitempty"`
	Relationships map[string]*common.Relationship `json:"relationships,omitempty"`
	Links         *common.Links                   `json:"links,omitempty"`
	Meta          common.NonStandardMeta          `json:"meta,omitempty"`
}

// example of Data
// Data{
// 	ResourceIdentifier: common.ResourceIdentifier{
// 		Type: "users",
// 		ID:   "1",
// 	},
// 	Attributes: map[string]interface{}{
// 		"name": "John Doe",
// 		"email": "
// 	},
// 	Relationships: map[string]*common.Relationship{
// 		"articles": &common.Relationship{
// 			Links: &common.Links{
// 				Self: "/users/1/relationships/articles",
// 				Related: "/users/1/articles",
// 			},
// 			Data: []*common.ResourceIdentifier{
// 				&common.ResourceIdentifier{
// 					Type: "articles",
// 					ID:   "1",
// 				},
// 				&common.ResourceIdentifier{
// 					Type: "articles",
// 					ID:   "2",
// 				},
// 			},
// 		},
// 	},
// 	Links: &common.Links{
// 		Self: "/users/1",
// 	},

type ListData = []*Data

type DataType interface {
	~*Data | ~ListData
}

type JSONAPIResponse[T DataType] struct {
	Data     T                      `json:"data"`
	Meta     common.NonStandardMeta `json:"meta,omitempty"`
	Links    *PaginationLinks       `json:"links,omitempty"`
	Included ListData               `json:"included,omitempty"`
}
