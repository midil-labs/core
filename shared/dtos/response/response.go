// Following the JSON API specification, this file contains the structs that represent the response objects.
// The JSON API specification is a standard for building APIs in JSON format. It defines the structure of the response objects and the relationships between them.
// Visit https://jsonapi.org/format/#document-structure to learn more about the JSON API specification.

package response

import (
	"fmt"
	"encoding/json"
	"strings"
	"github.com/midil-labs/core/shared/dtos"
	jsonApiError "github.com/midil-labs/core/shared/dtos/error"
)


func NewResource[T dtos.DTOInterface](id string, resourceType string, attributes T, opts ...ResourceOption[T]) *Resource[T] {
	r := &Resource[T]{
		ResourceIdentifier: ResourceIdentifier{
			ID:   id,
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


func NewSingleResourceResponse[T dtos.DTOInterface](opts ...ResponseOption[T]) *SingleResourceResponse[T] {
    builder := &ResourceResponse[T]{}
    for _, opt := range opts {
        opt(builder)
    }

    return &SingleResourceResponse[T]{
        Data:     builder.resource,
        Meta:     builder.Meta,
        Included: builder.Included,
    }
}

func NewMultipleResourcesResponse[T dtos.DTOInterface](opts ...ResponseOption[T]) *MultipleResourcesResponse[T] {
    builder := &ResourceResponse[T]{}
    for _, opt := range opts {
        opt(builder)
    }

    if builder.resources == nil {
        builder.resources = []Resource[T]{}
    }

    return &MultipleResourcesResponse[T]{
        Data:     builder.resources,
        Links:    builder.Links,
        Meta:     builder.Meta,
        Included: builder.Included,
    }
}


func (r ResourceIdentifier) Validate() error {
	if strings.TrimSpace(r.ID) == "" {
		return fmt.Errorf("resource ID cannot be empty")
	}
	if strings.TrimSpace(r.Type) == "" {
		return fmt.Errorf("resource type cannot be empty")
	}
	return nil
}


func (r Resource[T]) Validate() error {
	if err := r.ResourceIdentifier.Validate(); err != nil {
		return err
	}

	if attrs, ok := any(r.Attributes).(dtos.DTOInterface); ok {
		if err := attrs.Validate(); err != nil {
			return fmt.Errorf("attributes validation failed: %v", err)
		}
	}

	for name, rel := range r.Relationships {
		if err := rel.Validate(); err != nil {
			return fmt.Errorf("relationship '%s' validation failed: %v", name, err)
		}
	}

	return nil
}


func (c *RelationshipData) UnmarshalJSON(b []byte) error {
	var resource ResourceIdentifier
	if err := json.Unmarshal(b, &resource); err == nil && resource.ID != "" {
		c.Resource = &resource
		return nil
	}
	
	var resources []ResourceIdentifier
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


func (r RelationshipData) Validate() error {
	if r.Resource != nil {
		return r.Resource.Validate()
	}
	for _, res := range r.Resources {
		if err := res.Validate(); err != nil {
			return err
		}
	}
	return nil
}


func (r Relationship) Validate() error {
	return r.Data.Validate()
}

func (r *ResourceResponse[T]) UnmarshalJSON(data []byte) error {
	var singleResponse struct {
		Data *Resource[T] `json:"data"`
		Meta  NonStandardMeta       `json:"meta,omitempty"`
	}
	if err := json.Unmarshal(data, &singleResponse); err == nil && singleResponse.Data != nil {
		r.resource = singleResponse.Data
		r.Meta = singleResponse.Meta
		return nil
	}

	var multiResponse struct {
		Data []Resource[T] `json:"data"`
		Links *PaginationLinks       `json:"links,omitempty"`
		Meta  NonStandardMeta        `json:"meta,omitempty"`
	}
	if err := json.Unmarshal(data, &multiResponse); err == nil {
		r.resources = multiResponse.Data
		r.Links = multiResponse.Links
		r.Meta = multiResponse.Meta
		return nil
	}

	return fmt.Errorf("invalid resource response format")
}


func (r ResourceResponse[T]) MarshalJSON() ([]byte, error) {
	response := make(map[string]interface{})

	if r.resource != nil {
		response["data"] = r.resource
	} else if len(r.resources) > 0 {
		response["data"] = r.resources
	}

	if r.Links != nil {
		response["links"] = r.Links
	}
	if r.Meta != nil {
		response["meta"] = r.Meta
	}
	if len(r.Included) > 0 {
		response["included"] = r.Included
	}

	return json.Marshal(response)
}

func (r ResourceResponse[T]) Validate() error {
	if r.resource != nil {
		if r.Links != nil {
			return fmt.Errorf("single resource response cannot have pagination links")
		}

		if err := r.resource.Validate(); err != nil {
			return fmt.Errorf("single resource validation failed: %v", err)
		}
	}

	for i, res := range r.resources {
		if err := res.Validate(); err != nil {
			return fmt.Errorf("resource at index %d validation failed: %v", i, err)
		}
	}

	return nil
}


func (r *ErrorResponse) Validate() error {
	if len(r.Errors) == 0 {
		return fmt.Errorf("errors array cannot be empty")
	}
	return nil
}

func (v *ErrorResponse) Add(status int, code, title, detail string, opts ...jsonApiError.Option) *ErrorResponse {
	err := jsonApiError.New(status, code, title, detail, opts...)
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

func NewErrorResponse(meta map[string]interface{}, errs ...jsonApiError.ErrorObject) *ErrorResponse {
	return &ErrorResponse{
		Errors: errs,
		Meta: meta,
	}
}

func (r *ErrorResponse) FilterByCodePrefix(prefix string) []jsonApiError.ErrorObject {
	var filtered []jsonApiError.ErrorObject
	for _, err := range r.Errors {
		if strings.HasPrefix(string(err.Code), prefix) {
			filtered = append(filtered, err)
		}
	}
	return filtered
}
