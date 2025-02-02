// Following the JSON API specification, this file contains the structs that represent the response objects.
// The JSON API specification is a standard for building APIs in JSON format. It defines the structure of the response objects and the relationships between them.
// Visit https://jsonapi.org/format/#document-structure to learn more about the JSON API specification.

package response

import (
	"github.com/midil-labs/core/shared/jsonapi/common"
)

func NewData(id, resourceType string, attributes map[string]interface{}, opts ...ResourceOption) *Data {
	id = common.ID(id)
	resource := &Data{
		ResourceIdentifier: common.ResourceIdentifier{ID: &id, Type: resourceType},
		Attributes:         attributes,
	}
	for _, opt := range opts {
		opt(resource)
	}
	return resource
}

func NewJsonAPIResponse[T DataType](data T, opts ...ResponseOption[T]) *JSONAPIResponse[T] {
	builder := &JSONAPIResponse[T]{Data: data}
	for _, opt := range opts {
		opt(builder)
	}
	return builder
}
