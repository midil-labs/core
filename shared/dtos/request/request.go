package response

import (
	"net/url"
	"strconv"
	"strings"
)


type Filter struct {
	Fields map[string][]string
}

type Sort struct {
	Fields []string
}

type PaginationQuery struct {
	PageSize   int
	PageNumber int
}

type Fields map[string][]string

type Include []string

type QueryParams struct {
	Filter     Filter
	Sort       Sort
	Page       PaginationQuery
	Fields     Fields
	Include    Include
}


func ParseFilter(values url.Values) Filter {
	filter := Filter{
		Fields: make(map[string][]string),
	}

	for key, vals := range values {
		if strings.HasPrefix(key, "filter[") && strings.HasSuffix(key, "]") {
			field := key[len("filter[") : len(key)-1]
			filter.Fields[field] = append(filter.Fields[field], vals...)
		}
	}

	return filter
}


func ParseSort(sortParam string) Sort {
	sort := Sort{}
	if sortParam == "" {
		return sort
	}

	fields := strings.Split(sortParam, ",")
	for _, field := range fields {
		field = strings.TrimSpace(field)
		if field != "" {
			sort.Fields = append(sort.Fields, field)
		}
	}
	return sort
}


func ParsePagination(values url.Values) PaginationQuery {
	pagination := PaginationQuery{
		PageSize:   100, // default page size
		PageNumber: 1,  // default page number
	}

	if pageSizeStr := values.Get("page[size]"); pageSizeStr != "" {
		if size, err := strconv.Atoi(pageSizeStr); err == nil && size > 0 {
			pagination.PageSize = size
		}
	}

	if pageNumberStr := values.Get("page[number]"); pageNumberStr != "" {
		if number, err := strconv.Atoi(pageNumberStr); err == nil && number > 0 {
			pagination.PageNumber = number
		}
	}

	return pagination
}


func ParseFields(includeParam string) Fields {
	fields := make(Fields)
	if includeParam == "" {
		return fields
	}

	resourceFields := strings.Split(includeParam, ",")
	for _, rf := range resourceFields {
		rf = strings.TrimSpace(rf)
		if rf == "" {
			continue
		}
		parts := strings.SplitN(rf, ".", 2)
		if len(parts) == 2 {
			resourceType := parts[0]
			field := parts[1]
			fields[resourceType] = append(fields[resourceType], field)
		} else {
			// If no resource type is specified, apply to all or handle accordingly
			// This depends on your API's design decisions
		}
	}

	return fields
}

func ParseInclude(includeParam string) Include {
	if includeParam == "" {
		return Include{}
	}

	relations := strings.Split(includeParam, ",")
	for i, rel := range relations {
		relations[i] = strings.TrimSpace(rel)
	}
	return relations
}

func ParseQueryParams(values url.Values) QueryParams {
	return QueryParams{
		Filter:     ParseFilter(values),
		Sort:       ParseSort(values.Get("sort")),
		Page: 		ParsePagination(values),
		Fields:     ParseFields(values.Get("fields")),
		Include:    ParseInclude(values.Get("include")),
	}
}
