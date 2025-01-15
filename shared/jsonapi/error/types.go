// types.go
package error

import (
	"github.com/midil-labs/core/shared/jsonapi/common"
)

// Supported languages for localization
type Language string

// ErrCode represents the error code
type ErrCode string

// ErrorObject follows the JSON:API specification
type ErrorObject struct {
	ID        string                    `json:"id,omitempty" validate:"omitempty,uuid"`
	Links     *ErrorLinks               `json:"links,omitempty" validate:"omitempty"`
	Status    string                    `json:"status,omitempty" validate:"required,numeric"`
	Code      ErrCode                   `json:"code,omitempty" validate:"required"`
	Title     string                    `json:"title,omitempty" validate:"required"`
	Detail    string                    `json:"detail,omitempty"`
	Source    *ErrorSource              `json:"source,omitempty"`
	Meta      map[string]interface{}    `json:"meta,omitempty"`
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

type JSONAPIError struct {
	Errors []ErrorObject          `json:"errors" validate:"required"`
	Meta   common.NonStandardMeta `json:"meta,omitempty"`
}

