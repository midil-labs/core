package error

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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

func (j *JSONAPIError) AddError(err *ErrorObject) *JSONAPIError {
	j.Errors = append(j.Errors, err)
	return j
}
