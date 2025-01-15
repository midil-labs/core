package request

import (
	"net/url"
	"strconv"
	"strings"
	"encoding/json"
	"io"
	"errors"
)


// ParseFilter parses query parameters of the form filter[foo]=bar and returns a Filter.
func ParseFilter(values url.Values) Filter {
	filter := make(Filter)

	for key, vals := range values {
		if strings.HasPrefix(key, "filter[") && strings.HasSuffix(key, "]") {
			field := key[len("filter[") : len(key)-1] // Extract "foo" from "filter[foo]"
			filter[field] = append(filter[field], vals...)
		}
	}

	return filter
}

// ParseSort parses the "sort" query parameter and returns a Sort slice.
// Example: "?sort=-created,title" => Sort{"-created", "title"}
func ParseSort(values url.Values) Sort {
	sortParam := values.Get("sort")
	sortSlice := Sort{}

	if sortParam == "" {
		return sortSlice
	}

	fields := strings.Split(sortParam, ",")
	for _, field := range fields {
		field = strings.TrimSpace(field)
		if field != "" {
			sortSlice = append(sortSlice, field)
		}
	}

	return sortSlice
}

// ParsePagination parses "page[size]" and "page[number]" query parameters and returns a PaginationQuery.
// Defaults: Page{Size: 100, Number: 1}
// Example: "?page[size]=10&page[number]=2" => PaginationQuery{Size: 10, Number: 2}
func ParsePagination(values url.Values) Page {
	pagination := Page{
		Size:   MaximumPaginationSize,
		Number: DefaultPageNumber,
	}

	if pageSizeStr := values.Get("page[size]"); pageSizeStr != "" {
		if size, err := strconv.Atoi(pageSizeStr); err == nil && size > 0 {
			pagination.Size = size
		}
	}

	if pageNumberStr := values.Get("page[number]"); pageNumberStr != "" {
		if number, err := strconv.Atoi(pageNumberStr); err == nil && number > 0 {
			pagination.Number = number
		}
	}

	return pagination
}

// ParseFields parses the "fields[resourceType]=field1,field2" query parameters and returns a Fields map.
// Example: "?fields[articles]=title,body&fields[users]=email" => Fields{"articles": ["title", "body"], "users": ["email"]}
func ParseFields(values url.Values) Fields {
	fieldsMap := make(Fields, len(values)/2) // Estimate initial capacity

	for key, vals := range values {
		if strings.HasPrefix(key, "fields[") && strings.HasSuffix(key, "]") {
			resourceType := key[len("fields[") : len(key)-1] // Extract "articles" from "fields[articles]"
			if resourceType == "" {
				continue // Skip if resourceType is empty
			}

			for _, val := range vals {
				// Split once and trim each field
				fieldParts := strings.Split(val, ",")
				for _, field := range fieldParts {
					trimmedField := strings.TrimSpace(field)
					if trimmedField != "" {
						fieldsMap[resourceType] = append(fieldsMap[resourceType], trimmedField)
					}
				}
			}
		}
	}

	return fieldsMap
}



// ParseInclude parses the "include" query parameter and returns an Include slice.
// Example: "?include=comments,author.profile" => Include{"comments", "author.profile"}
func ParseInclude(values url.Values) Include {
	includeParam := values.Get("include")
	if includeParam == "" {
		return Include{}
	}

	relations := strings.Split(includeParam, ",")
	for i, rel := range relations {
		relations[i] = strings.TrimSpace(rel)
	}
	return relations
}

// ParseQueryParams aggregates all individual parsing functions and returns a QueryParams struct.
// The 'strictJSONAPI' flag determines how "fields" are parsed.
func ParseQueryParams(values url.Values) Query {
	pagination := ParsePagination(values)
	return Query{
		Filter:  ParseFilter(values),
		Sort:    ParseSort(values),
		Page:    &pagination,
		Fields:  ParseFields(values),
		Include: ParseInclude(values),
	}
}

// ParseBody decodes an io.Reader into a struct of type T and returns it.
func ParseBody[T RequestType](body io.Reader) (T, error) {
	var parsedBody T

	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&parsedBody); err != nil {
		return parsedBody, err
	}

	// Ensure there's no extra data after the JSON object
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return parsedBody, errors.New("unexpected data after JSON object")
	}
	return parsedBody, nil
}
