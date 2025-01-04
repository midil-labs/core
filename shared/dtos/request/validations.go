package request

import (
	"fmt"
	"slices"
	"strings"

	jsonApiError "github.com/midil-labs/core/shared/dtos/error"
)

func (f Filter) Validate(config Filter) []jsonApiError.ErrorObject {
	var errors []jsonApiError.ErrorObject
	if len(config.Fields) > 0 {
		for field := range f.Fields {
			if !slices.Contains(config.Fields[field], field) {
				errors = append(errors, jsonApiError.ErrorObject{
					Source: &jsonApiError.ErrorSource{
						Pointer: "filter",
					},
					Detail: fmt.Sprintf("invalid filter field: %s", field),
				})
			}
		}
	}
	return errors
}


func (s Sort) Validate(config Sort) []jsonApiError.ErrorObject {
	var errors []jsonApiError.ErrorObject
	if len(config.Fields) > 0 {
		for _, field := range s.Fields {
			cleanField := strings.TrimPrefix(field, "-")
			if !slices.Contains(config.Fields, cleanField) {
				errors = append(errors, jsonApiError.ErrorObject{
					Source: &jsonApiError.ErrorSource{
						Pointer: "sort",
					},
					Detail: fmt.Sprintf("invalid sort field: %s", cleanField),
				})
			}
		}
	}
	return errors
}


func (p PaginationQuery) Validate(config PaginationQuery) []jsonApiError.ErrorObject {
	var errors []jsonApiError.ErrorObject
	if p.PageSize <= 0 {
		errors = append(errors, 
			jsonApiError.ErrorObject{
				Source: &jsonApiError.ErrorSource{
					Pointer:"page[size]",
				},
				Detail: "PageSize must be greater than 0",
			},
		)
	}

	if p.PageSize >= config.PageSize{
		errors = append(errors, 
		jsonApiError.ErrorObject{
			Source: &jsonApiError.ErrorSource{
				Pointer: "page[size]",
			},
			Detail: fmt.Sprintf("PageSize must be less than or equal to %s", config.PageSize),
		})
	}

	if p.PageNumber < 0 {
		errors = append(errors, 
			jsonApiError.ErrorObject{
				Source: &jsonApiError.ErrorSource{
					Pointer: "page[number]",
				},
				Detail: "PageNumber cannot be negative",
			},
		)
	}
	return errors
}


func (f Fields) Validate(config Fields) []jsonApiError.ErrorObject {
	var errors []jsonApiError.ErrorObject
	if len(config) > 0 {
		for resourceType, fields := range f {
			allowedFields, exists := config[resourceType]
			if !exists {
				errors = append(errors, jsonApiError.ErrorObject{
					Source: &jsonApiError.ErrorSource{
						Pointer: "fields",
					},
					Detail: fmt.Sprintf("invalid resource type: %s", resourceType),
				})
				continue
			}

            allowedFieldMap := make(map[string]bool)
            for _, f := range allowedFields {
                allowedFieldMap[f] = true
            }

			for _, field := range fields {
				if !allowedFieldMap[field] {
					errors = append(errors, 
						jsonApiError.ErrorObject{
							Source: &jsonApiError.ErrorSource{Pointer: fmt.Sprintf("fields[%s]", resourceType)},
							Detail: fmt.Sprintf("invalid field: %s", field),
						})
				}
			}
		}
	}
	return errors
}

func (i Include) Validate(config Include) []jsonApiError.ErrorObject {
	var errors []jsonApiError.ErrorObject

	if len(config) > 0 {
		for _, include := range i {
			if !slices.Contains(config, include) {
				errors = append(errors, jsonApiError.ErrorObject{
					Source: &jsonApiError.ErrorSource{Pointer: "include"},
					Detail: fmt.Sprintf("invalid include: %s", include),
				})
			}
		}
	}
	return errors
}


func (q QueryParams) Validate(config QueryParams) []jsonApiError.ErrorObject {
    var errObjects []jsonApiError.ErrorObject
    errObjects = append(errObjects, q.Filter.Validate(config.Filter)...)
    errObjects = append(errObjects, q.Sort.Validate(config.Sort)...)
    errObjects = append(errObjects, q.Page.Validate(config.Page)...)
    errObjects = append(errObjects, q.Fields.Validate(config.Fields)...)
    errObjects = append(errObjects, q.Include.Validate(config.Include)...)

    return errObjects
}
