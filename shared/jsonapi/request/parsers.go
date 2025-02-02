package request

import (
	"encoding/json"
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/midil-labs/core/shared/jsonapi/jsonapierror"
)

// ParseFilter parses query parameters of the form filter[foo]=bar and returns a Filter.
func ParseFilter(values url.Values) (Filter, jsonapierror.ErrorObjects) {
	filter := make(Filter)
	var errs jsonapierror.ErrorObjects

	for key, vals := range values {
		if strings.HasPrefix(key, "filter[") && strings.HasSuffix(key, "]") {
			field := key[len("filter[") : len(key)-1] // Extract "foo" from "filter[foo]"
			if field == "" {
				errs = append(errs, jsonapierror.NewValidationError(
					jsonapierror.ErrorSource{
						Parameter: "filter",
					},
					"filter field cannot be empty",
				))
				continue
			}
			filter[field] = append(filter[field], vals...)
		}
	}

	return filter, errs
}

// ParseSort parses the "sort" query parameter and returns a Sort slice.
func ParseSort(values url.Values) (Sort, jsonapierror.ErrorObjects) {
	sortParam := values.Get("sort")
	sortSlice := Sort{}
	var errs jsonapierror.ErrorObjects

	if sortParam == "" {
		return sortSlice, errs
	}

	fields := strings.Split(sortParam, ",")
	for _, field := range fields {
		field = strings.TrimSpace(field)
		if field == "" {
			errs = append(errs, jsonapierror.NewValidationError(
				jsonapierror.ErrorSource{
					Parameter: "sort",
				},
				"sort field cannot be empty",
			))
			continue
		}
		sortSlice = append(sortSlice, field)
	}

	return sortSlice, errs
}

// ParsePagination parses "page[size]" and "page[number]" query parameters and returns a PaginationQuery.
func ParsePagination(values url.Values) (Page, jsonapierror.ErrorObjects) {
	pagination := Page{
		Size:   MaximumPaginationSize,
		Number: DefaultPageNumber,
	}
	var errs jsonapierror.ErrorObjects
	for key, _ := range values {
		if strings.HasPrefix(key, "page[") && strings.HasSuffix(key, "]") {
			field := key[len("page[") : len(key)-1]
			if field != "size" && field != "number" {
				errs = append(errs, jsonapierror.NewValidationError(
					jsonapierror.ErrorSource{
						Parameter: "page",
					},
					"page parameter must be either 'size' or 'number'",
				))
				continue
			}
		}
	}

	if pageSizeStr := values.Get("page[size]"); pageSizeStr != "" {
		size, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			errs = append(errs, jsonapierror.NewValidationError(
				jsonapierror.ErrorSource{
					Parameter: "page[size]",
				},
				"invalid page[size] value",
			))
		} else if size <= 0 {
			errs = append(errs, jsonapierror.NewValidationError(
				jsonapierror.ErrorSource{
					Parameter: "page[size]",
				},
				"page[size] must be greater than 0",
			))
		} else {
			pagination.Size = size
		}
	}

	if pageNumberStr := values.Get("page[number]"); pageNumberStr != "" {
		number, err := strconv.Atoi(pageNumberStr)
		if err != nil {
			errs = append(errs, jsonapierror.NewValidationError(
				jsonapierror.ErrorSource{
					Parameter: "page[number]",
				},
				"invalid page[number] value",
			))
		} else if number <= 0 {
			errs = append(errs, jsonapierror.NewValidationError(
				jsonapierror.ErrorSource{
					Parameter: "page[number]",
				},
				"page[number] must be greater than 0",
			))
		} else {
			pagination.Number = number
		}
	}

	return pagination, errs
}

// ParseFields parses the "fields[resourceType]=field1,field2" query parameters and returns a Fields map.
func ParseFields(values url.Values) (Fields, jsonapierror.ErrorObjects) {
	fieldsMap := make(Fields)
	var errs jsonapierror.ErrorObjects

	for key, vals := range values {
		if strings.HasPrefix(key, "fields[") && strings.HasSuffix(key, "]") {
			resourceType := key[len("fields[") : len(key)-1] // Extract "articles" from "fields[articles]"
			if resourceType == "" {
				errs = append(errs, jsonapierror.NewValidationError(
					jsonapierror.ErrorSource{
						Parameter: "fields",
					},
					"resource type cannot be empty",
				))
				continue
			}

			for _, val := range vals {
				fieldParts := strings.Split(val, ",")
				for _, field := range fieldParts {
					trimmedField := strings.TrimSpace(field)
					if trimmedField == "" {
						errs = append(errs, jsonapierror.NewValidationError(
							jsonapierror.ErrorSource{
								Parameter: "fields",
							},
							"field name cannot be empty",
						))
						continue
					}
					fieldsMap[resourceType] = append(fieldsMap[resourceType], trimmedField)
				}
			}
		}
	}

	return fieldsMap, errs
}

// ParseInclude parses the "include" query parameter and returns an Include slice.
func ParseInclude(values url.Values) (Include, jsonapierror.ErrorObjects) {
	includeParam := values.Get("include")
	var errs jsonapierror.ErrorObjects

	if includeParam == "" {
		return Include{}, errs
	}

	relations := strings.Split(includeParam, ",")
	for i, rel := range relations {
		trimmedRel := strings.TrimSpace(rel)
		if trimmedRel == "" {
			errs = append(errs, jsonapierror.NewValidationError(
				jsonapierror.ErrorSource{
					Parameter: "include",
				},
				"include relation cannot be empty",
			))
			continue
		}
		relations[i] = trimmedRel
	}
	return relations, errs
}

// ParseBody decodes an io.Reader into a struct of type T and returns it.
func ParseBody[T BodyType](body io.Reader) (T, jsonapierror.ErrorObjects) {
	var errs jsonapierror.ErrorObjects
	var req struct {
		Data T `json:"data"`
	}

	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		errs = append(errs, jsonapierror.NewValidationError(
			jsonapierror.ErrorSource{
				Pointer: "/data",
			},
			"invalid JSON body",
		))
	}

	// Ensure there's no extra data after the JSON object
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		errs = append(errs, jsonapierror.NewValidationError(
			jsonapierror.ErrorSource{
				Pointer: "/data",
			},
			"unexpected data after JSON object",
		))
	}

	return req.Data, errs
}

// ParseQueryParams aggregates all individual parsing functions and returns a QueryParams struct.
func ParseQueryParams(values url.Values) (Query, jsonapierror.ErrorObjects) {
	var errs jsonapierror.ErrorObjects

	filter, filterErrs := ParseFilter(values)
	errs = append(errs, filterErrs...)

	sort, sortErrs := ParseSort(values)
	errs = append(errs, sortErrs...)

	pagination, paginationErrs := ParsePagination(values)
	errs = append(errs, paginationErrs...)

	fields, fieldsErrs := ParseFields(values)
	errs = append(errs, fieldsErrs...)

	include, includeErrs := ParseInclude(values)
	errs = append(errs, includeErrs...)

	return Query{
		Filter:  filter,
		Sort:    sort,
		Page:    &pagination,
		Fields:  fields,
		Include: include,
	}, errs
}
