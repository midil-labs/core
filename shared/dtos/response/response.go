// Following the JSON API specification, this file contains the structs that represent the response objects.
// The JSON API specification is a standard for building APIs in JSON format. It defines the structure of the response objects and the relationships between them.
// Visit https://jsonapi.org/format/#document-structure to learn more about the JSON API specification.

package response

import (
	"fmt"
	"encoding/json"
	"strings"
	"github.com/midil-labs/core/shared/dtos"
)


type PaginationLinks struct {
	Self  string `json:"self,omitempty"`
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
}


type Meta struct {
	Pagination Pagination `json:"pagination,omitempty"`
}


type Pagination struct {
	CurrentPage int64 `json:"current_page"`
	PrevPage int64 `json:"prev_page"`
	NextPage int64 `json:"next_page"`
	TotalPages int64 `json:"total_pages"`
	TotalCount int64 `json:"total_count"`
}


type RelatedLink struct {
	Href        string            `json:"href,omitempty"`
	Title       string            `json:"title,omitempty"`
	DescribedBy string            `json:"describedby,omitempty"`
	Meta        map[string]interface{} `json:"meta,omitempty"`
}

type Links struct {
	Self    string      `json:"self,omitempty"`
	Related *RelatedLink `json:"related,omitempty"`
}

type ResourceIdentifier struct {
	ID   string `json:"id"`
	Type string `json:"type"`
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


type Resource[T dtos.DTOInterface] struct {
	ResourceIdentifier
	Attributes    T                      `json:"attributes,omitempty"`
	Relationships map[string]Relationship `json:"relationships,omitempty"`
	Links         *Links                 `json:"links,omitempty"`
	Meta          map[string]any         `json:"meta,omitempty"`
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


type RelationshipData struct {
	Resource   *ResourceIdentifier
	Resources []ResourceIdentifier
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


type Relationship struct {
	Data  RelationshipData `json:"data"`
	Links *Links            `json:"links,omitempty"`
	Meta  map[string]interface{}    `json:"meta,omitempty"`
}

func (r Relationship) Validate() error {
	return r.Data.Validate()
}

type SingleResourceResponse[T dtos.DTOInterface] struct {
	Data    *Resource[T]    `json:"data,omitempty"`
	Meta    *Meta           `json:"meta,omitempty"`
	Included []Resource[dtos.DTOInterface] `json:"included,omitempty"`
}


type MultipleResourcesResponse[T dtos.DTOInterface] struct {
	Data    []Resource[T]    `json:"data,omitempty"`
	Links   *PaginationLinks `json:"links,omitempty"`
	Meta    *Meta            `json:"meta,omitempty"`
	Included []Resource[dtos.DTOInterface] `json:"included,omitempty"`
}

type ResourceResponse[T dtos.DTOInterface] struct {
	resource   *Resource[T] 
	resources []Resource[T]
	Links    *Links     `json:"links,omitempty"`
	Meta     *Meta      `json:"meta,omitempty"`
	Included []any      `json:"included,omitempty"`
}

func (r *ResourceResponse[T]) UnmarshalJSON(data []byte) error {
	var singleResponse struct {
		Data *Resource[T] `json:"data"`
		Links *Links      `json:"links,omitempty"`
		Meta  *Meta       `json:"meta,omitempty"`
	}
	if err := json.Unmarshal(data, &singleResponse); err == nil && singleResponse.Data != nil {
		r.resource = singleResponse.Data
		r.Links = singleResponse.Links
		r.Meta = singleResponse.Meta
		return nil
	}

	var multiResponse struct {
		Data []Resource[T] `json:"data"`
		Links *Links       `json:"links,omitempty"`
		Meta  *Meta        `json:"meta,omitempty"`
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

type ErrorSource struct {
	Pointer   string `json:"pointer,omitempty"`
	Parameter string `json:"parameter,omitempty"`
}

type ErrorObject struct {
	Status string            `json:"status,omitempty"`
	Title  string            `json:"title,omitempty"`
	Detail string            `json:"detail,omitempty"`
	Source *ErrorSource      `json:"source,omitempty"`
}


func (e ErrorObject) Validate() error {
	if strings.TrimSpace(e.Status) == "" {
		return fmt.Errorf("error status cannot be empty")
	}
	if strings.TrimSpace(e.Title) == "" {
		return fmt.Errorf("error title cannot be empty")
	}
	return nil
}


type ErrorResponse struct {
	Errors  []ErrorObject `json:"errors"`
}

func (e ErrorResponse) Validate() error {
	if len(e.Errors) == 0 {
		return fmt.Errorf("errors cannot be empty")
	}
	for i, errObj := range e.Errors {
		if err := errObj.Validate(); err != nil {
			return fmt.Errorf("error object at index %d validation failed: %v", i, err)
		}
	}
	return nil
}

