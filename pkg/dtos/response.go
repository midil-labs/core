package dtos

type JSONAPI struct {
	Version string `json:"version"`
}

type Links struct {
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

type Resource[T any] struct {
	ID            string                 `json:"id"`
	Type          string                 `json:"type"`
	Attributes    T                      `json:"attributes,omitempty"`
	Relationships map[string]Relationship `json:"relationships,omitempty"`
	Links         *Links                 `json:"links,omitempty"`
}

type RelationshipData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Relationship struct {
	Data  interface{} `json:"data,omitempty"`
	Links *Links      `json:"links,omitempty"`
	Meta  *Meta       `json:"meta,omitempty"`
}

type SingleResourceResponse[T any] struct {
	JSONAPI *JSONAPI        `json:"jsonapi,omitempty"`
	Data    *Resource[T]    `json:"data,omitempty"`
	Links   *Links          `json:"links,omitempty"`
	Meta    *Meta           `json:"meta,omitempty"`
	Included []Resource[any] `json:"included,omitempty"`
}

type MultipleResourcesResponse[T any] struct {
	JSONAPI *JSONAPI        `json:"jsonapi,omitempty"`
	Data    []Resource[T]    `json:"data,omitempty"`
	Links   *Links           `json:"links,omitempty"`
	Meta    *Meta            `json:"meta,omitempty"`
	Included []Resource[any] `json:"included,omitempty"`
}

type ErrorObject struct {
	ID     string            `json:"id,omitempty"`
	Links  *Links            `json:"links,omitempty"`
	Status string            `json:"status,omitempty"`
	Code   string            `json:"code,omitempty"`
	Title  string            `json:"title,omitempty"`
	Detail string            `json:"detail,omitempty"`
	Source *ErrorSource      `json:"source,omitempty"`
	Meta   map[string]string `json:"meta,omitempty"`
}

type ErrorSource struct {
	Pointer   string `json:"pointer,omitempty"`
	Parameter string `json:"parameter,omitempty"`
}

type ErrorResponse struct {
	JSONAPI *JSONAPI      `json:"jsonapi,omitempty"`
	Errors  []ErrorObject `json:"errors"`
	Meta    *Meta         `json:"meta,omitempty"`
}
