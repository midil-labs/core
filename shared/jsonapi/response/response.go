// Following the JSON API specification, this file contains the structs that represent the response objects.
// The JSON API specification is a standard for building APIs in JSON format. It defines the structure of the response objects and the relationships between them.
// Visit https://jsonapi.org/format/#document-structure to learn more about the JSON API specification.

package response

import (
	"encoding/json"
	"fmt"
	"github.com/midil-labs/core/shared/jsonapi/common"
)

func NewResource(id, resourceType string, attributes map[string]interface{}, opts ...ResourceOption) *Resource {
	id = common.ID(id)
	resource := &Resource{
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

func (c *RelationshipData) UnmarshalJSON(b []byte) error {
	var resource common.ResourceIdentifier
	if err := json.Unmarshal(b, &resource); err == nil && resource.ID != nil && *resource.ID != "" {
		c.Resource = &resource
		return nil
	}

	var resources []common.ResourceIdentifier
	if err := json.Unmarshal(b, &resources); err == nil {
		c.Resources = resources
		return nil
	}

	return fmt.Errorf("data field is neither a resource object nor a valid array of objects")
}

func (c RelationshipData) MarshalJSON() ([]byte, error) {
	if c.Resource != nil {
		return json.Marshal(c.Resource)
	}
	return json.Marshal(c.Resources)
}
