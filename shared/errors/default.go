package errors

import (
	jsonAPIError "github.com/midil-labs/core/shared/dtos/error"
)


type ValidationErrors struct {
	errors []jsonAPIError.ErrorObject
	meta   map[string]interface{}
}

func NewValidationErrors() *ValidationErrors {
	return &ValidationErrors{
		errors: make([]jsonAPIError.ErrorObject, 0),
		meta:   make(map[string]interface{}),
	}
}

func (v *ValidationErrors) Add(field, detail string, opts ...jsonAPIError.Option) *ValidationErrors {
	err := jsonAPIError.NewValidationError(detail, field, opts...)
	v.errors = append(v.errors, *err)
	return v
}

func (v *ValidationErrors) WithMeta(key string, value interface{}) *ValidationErrors {
	v.meta[key] = value
	return v
}

