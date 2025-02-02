// Package jsonapierror provides a standardized way to handle and communicate errors
// in JSON API responses. It includes predefined error codes, HTTP status mappings,
// and multilingual error messages for consistent error handling across the application.
//
// Key Features:
// - Predefined error codes for common scenarios (e.g., authentication, validation, resources).
// - Mapping of error codes to appropriate HTTP status codes.
// - Support for multilingual error messages (e.g., English, Spanish, French).
// - Easy-to-use functions for retrieving HTTP status codes and error messages.


package jsonapierror

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"net/http"
)

func New(status int, code ErrCode, title, detail string, opts ...Option) *ErrorObject {
	err := &ErrorObject{
		Status:    strconv.Itoa(status),
		Code:      code,
		Title:     title,
		Detail:    detail,
		localizer: make(map[Language]Localization),
	}
	for _, opt := range opts {
		opt(err)
	}

	return err
}

func NewAppError(code ErrCode, details string, lang Language) *ErrorObject {
	return &ErrorObject{
		Code:   code,
		Title:  GetErrorMessage(code, lang),
		Detail: details,
		Status: strconv.Itoa(GetHTTPStatus(code)),
	}
}


func (e *ErrorObject) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.Status, e.Title, e.Detail)
}

func (e *ErrorObject) Unwrap() error {
	return e.internal
}

func (e *ErrorObject) GetLocalizedError(lang Language) *ErrorObject {
	if loc, exists := e.localizer[lang]; exists {
		localized := *e
		localized.Title = loc.Title
		localized.Detail = loc.Detail
		return &localized
	}
	return e
}

// GetHTTPStatus returns the HTTP status code for a given error code.
func GetHTTPStatus(code ErrCode) int {
	if status, exists := ErrorCodeToHTTPStatus[code]; exists {
		return status
	}

	return http.StatusInternalServerError
}

// GetErrorMessage returns the error message for a given error code and language.
func GetErrorMessage(code ErrCode, lang Language) string {
	if msg, exists := ErrorMessages[code]; exists {
		switch lang {
		case EN:
			return msg.EN
		case ES:
			return msg.ES
		case FR:
			return msg.FR
		}
	}
	return "Internal server error" // Default message
}


func (e *ErrorObject) MarshalJSON() ([]byte, error) {
	type Alias ErrorObject
	return json.Marshal(&struct {
		*Alias
		Meta map[string]interface{} `json:"meta,omitempty"`
	}{
		Alias: (*Alias)(e),
		Meta:  e.enrichMeta(),
	})
}

func (e *ErrorObject) enrichMeta() map[string]interface{} {
	if e.Meta == nil {
		e.Meta = make(map[string]interface{})
	}

	if len(e.localizer) > 0 {
		languages := make([]string, 0, len(e.localizer))
		for lang := range e.localizer {
			languages = append(languages, string(lang))
		}
		e.Meta["available_languages"] = languages
	}

	return e.Meta
}

func (r *JSONAPIError) Validate() error {
	if len(r.Errors) == 0 {
		return fmt.Errorf("errors array cannot be empty")
	}
	return nil
}

func (r *JSONAPIError) MarshalJSON() ([]byte, error) {
	if err := r.Validate(); err != nil {
		return nil, fmt.Errorf("invalid error response: %w", err)
	}
	type Alias JSONAPIError
	return json.Marshal((*Alias)(r))
}

func (r *JSONAPIError) UnmarshalJSON(data []byte) error {
	type Alias JSONAPIError
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

func NewJSONAPIError(meta map[string]interface{}, errs ...*ErrorObject) *JSONAPIError {
	return &JSONAPIError{
		Errors: errs,
		Meta:   meta,
	}
}

func (j *JSONAPIError) FilterByCodePrefix(prefix string) ErrorObjects {
	var filtered ErrorObjects
	for _, err := range j.Errors {
		if strings.HasPrefix(string(err.Code), prefix) {
			filtered = append(filtered, err)
		}
	}
	return filtered
}

func (j *JSONAPIError) AddError(err *ErrorObject) {
	j.Errors = append(j.Errors, err)
}
