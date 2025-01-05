// Following the JSON API specification, this file contains the structs that represent the response objects.
// The JSON API specification is a standard for building APIs in JSON format. It defines the structure of the response objects and the relationships between them.
// Visit https://jsonapi.org/format/#document-structure to learn more about the JSON API specification.

package response

import (
	"fmt"
	"encoding/json"
	"strings"
	jsonAPIError "github.com/midil-labs/core/shared/dtos/error"
	"github.com/midil-labs/core/shared/dtos/common"
)


func NewResource(id string, resourceType string, attributes map[string]any, opts ...ResourceOption) *Resource {
	r := &Resource{
		ResourceIdentifier: common.ResourceIdentifier{
			ID:   &id,
			Type: resourceType,
		},
		Attributes:    attributes,
		Relationships: make(map[string]Relationship),
		Meta:         make(map[string]any),
	}

	for _, opt := range opts {
		opt(r)
	}
	return r
}


func NewJsonAPIResponse[T DataType](data T, opts ...ResponseOption[T]) *JSONAPIResponse[T] {
	builder := &JSONAPIResponse[T]{Data: data}
	for _, opt := range opts {
		opt(builder)
	}
	return builder
}


func NewSingleAPIResponse(data *Resource, opts ...ResponseOption[*Resource]) *JSONAPIResponse[*Resource] {
	return NewJsonAPIResponse[*Resource](data, opts...)
}


func NewListAPIResponse(data []Resource, opts ...ResponseOption[ListResource]) *JSONAPIResponse[ListResource] {
	return NewJsonAPIResponse(data, opts...)
}


func (c *RelationshipData) UnmarshalJSON(b []byte) error {
	var resource common.ResourceIdentifier
	if err := json.Unmarshal(b, &resource); err == nil && resource.ID != "" {
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


func (r *ErrorResponse) Validate() error {
	if len(r.Errors) == 0 {
		return fmt.Errorf("errors array cannot be empty")
	}
	return nil
}


func (v *ErrorResponse) Add(status int, code, title, detail string, opts ...jsonAPIError.Option) *ErrorResponse {
	err := jsonAPIError.New(status, code, title, detail, opts...)
	v.Errors = append(v.Errors, *err)
	return v
}


func (r *ErrorResponse) MarshalJSON() ([]byte, error) {
	if err := r.Validate(); err != nil {
		return nil, fmt.Errorf("invalid error response: %w", err)
	}
	type Alias ErrorResponse
	return json.Marshal((*Alias)(r))
}


func (r *ErrorResponse) UnmarshalJSON(data []byte) error {
	type Alias ErrorResponse
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return r.Validate()
}


func NewErrorResponse(meta map[string]interface{}, errs ...jsonAPIError.ErrorObject) *ErrorResponse {
	return &ErrorResponse{
		Errors: errs,
		Meta: meta,
	}
}

func (r *ErrorResponse) FilterByCodePrefix(prefix string) []jsonAPIError.ErrorObject {
	var filtered []jsonAPIError.ErrorObject
	for _, err := range r.Errors {
		if strings.HasPrefix(string(err.Code), prefix) {
			filtered = append(filtered, err)
		}
	}
	return filtered
}

