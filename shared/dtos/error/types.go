// types.go
package error

// Supported languages for localization
type Language string

const (
	EN Language = "en"
	ES Language = "es"
	FR Language = "fr"
)

// ErrorObject follows the JSON:API specification
type ErrorObject struct {
	ID     string                 `json:"id,omitempty" validate:"omitempty,uuid"`
	Links  *ErrorLinks            `json:"links,omitempty" validate:"omitempty"`
	Status string                 `json:"status,omitempty" validate:"required,numeric"`
	Code   string                 `json:"code,omitempty" validate:"required"`
	Title  string                 `json:"title,omitempty" validate:"required"`
	Detail string                 `json:"detail,omitempty"`
	Source *ErrorSource           `json:"source,omitempty"`
	Meta   map[string]interface{} `json:"meta,omitempty"`

	// Internal fields
	internal  error                     `json:"-"`
	localizer map[Language]Localization `json:"-"`
}

// ErrorLinks contains reference links related to the error
type ErrorLinks struct {
	About string `json:"about,omitempty" validate:"omitempty,url"`
	Type  string `json:"type,omitempty" validate:"omitempty,url"`
}

// ErrorSource contains references to the source of the error
type ErrorSource struct {
	Pointer   string `json:"pointer,omitempty" validate:"omitempty,startswith=/"`
	Parameter string `json:"parameter,omitempty"`
	Header    string `json:"header,omitempty"`
}

// Localization contains localized versions of error messages
type Localization struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}
