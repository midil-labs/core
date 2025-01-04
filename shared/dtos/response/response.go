// Following the JSON API specification, this file contains the structs that represent the response objects.
// The JSON API specification is a standard for building APIs in JSON format. It defines the structure of the response objects and the relationships between them.
// Visit https://jsonapi.org/format/#document-structure to learn more about the JSON API specification.

package response

import (
	"fmt"
	"encoding/json"
	"strings"
	"github.com/midil-labs/core/shared/dtos"
	jsonAPIError "github.com/midil-labs/core/shared/dtos/error"
)


func NewResource(id string, resourceType string, attributes any, opts ...ResourceOption) *Resource {
	r := &Resource{
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


func NewJsonAPIResponse(data Data, opts ...ResponseOption) *JSONAPIResponse {
	builder := &JSONAPIResponse{Data: data}
	for _, opt := range opts {
		opt(builder)
	}
	return builder
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


func (r Resource) Validate() error {
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

func (r *JSONAPIResponse) UnmarshalJSON(data []byte) error {
	var singleResponse struct {
		Data *Resource `json:"data"`
		Meta  NonStandardMeta       `json:"meta,omitempty"`
	}
	if err := json.Unmarshal(data, &singleResponse); err == nil && singleResponse.Data != nil {
		r.Data.resource = singleResponse.Data
		r.Meta = singleResponse.Meta
		return nil
	}

	var multiResponse struct {
		Data []Resource `json:"data"`
		Links *PaginationLinks       `json:"links,omitempty"`
		Meta  NonStandardMeta        `json:"meta,omitempty"`
	}
	if err := json.Unmarshal(data, &multiResponse); err == nil {
		r.Data.resources = multiResponse.Data
		r.Links = multiResponse.Links
		r.Meta = multiResponse.Meta
		return nil
	}

	return fmt.Errorf("invalid resource response format")
}


func (r JSONAPIResponse) MarshalJSON() ([]byte, error) {
	response := make(map[string]interface{})

	if r.Data.resource != nil {
		response["data"] = r.Data.resource
	} else if len(r.Data.resources) > 0 {
		response["data"] = r.Data.resources
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

func (r JSONAPIResponse) Validate() error {
	if r.Data.resource != nil {
		if r.Links != nil {
			return fmt.Errorf("single resource response cannot have pagination links")
		}

		if err := r.Data.resource.Validate(); err != nil {
			return fmt.Errorf("single resource validation failed: %v", err)
		}
	}

	for i, res := range r.Data.resources {
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

