package request

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const QueryParamsKey string = "queryParams"


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


func ParseSort(values url.Values) Sort {
	sortParam := values.Get("sort")
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


func ParseFields(values url.Values) Fields {
	includeParam:= values.Get("fields")

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

func ParseQueryParams(values url.Values) QueryParams {
	return QueryParams{
		Filter:     ParseFilter(values),
		Sort:       ParseSort(values),
		Page: 		ParsePagination(values),
		Fields:     ParseFields(values),
		Include:    ParseInclude(values),
	}
}


func GetQueryParams(r *http.Request) (QueryParams, bool) {
    queryParams, ok := r.Context().Value(QueryParamsKey).(QueryParams)
    return queryParams, ok
}


func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
