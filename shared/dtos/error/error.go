package error

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func New(status int, code, title, detail string, opts ...Option) *ErrorObject {
	err := &ErrorObject{
		Status:    strconv.Itoa(status),
		Code:     code,
		Title:    title,
		Detail:   detail,
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