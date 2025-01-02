package middleware

import (
	"context"
	"net/http"
	"github.com/midil-labs/core/shared/dtos/request"
    "fmt"
    "strings"
)
type contextKey string
const QueryParamsKey contextKey = contextKey(request.QueryParamsKey)

type Option func(*request.QueryParams)

func WithQueryParamsMiddleware(opts ...request.QueryParams) func(http.Handler) http.Handler {
    config := DefaultConfig()
    if len(opts) > 0 {
        config = opts[0]
    }
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            queryValues := r.URL.Query()
            queryParams := request.ParseQueryParams(queryValues)
            if err := queryParams.Validate(config); err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }
            ctx := context.WithValue(r.Context(), request.QueryParamsKey, queryParams)
            r = r.WithContext(ctx)
            next.ServeHTTP(w, r)
        })
    }
}


func hasValidationRules(config QueryParamsConfig) bool {
    return len(config.AllowedSortFields) > 0 ||
           len(config.AllowedFilterFields) > 0 ||
           len(config.AllowedIncludes) > 0 ||
           len(config.AllowedFields) > 0 ||
           config.MaxPageSize != DefaultConfig().MaxPageSize
}


func WithQueryParams(opts ...QueryParamsConfig) func(http.Handler) http.Handler {
    // If no config provided, use defaults
    config := DefaultConfig()
    if len(opts) > 0 {
        config = opts[0]
    }

    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            extReq := &ExtendedRequest{
                Request: r,
                QueryParams: request.ParseQueryParams(r.URL.Query()),
            }

            // Only validate if config has validation rules set
            if hasValidationRules(config) {
                validateQueryParams(extReq, config)
            }

            next.ServeHTTP(w, extReq)
        })
    }
}


type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// QueryParamsConfig holds configuration for query parameter processing
type QueryParamsConfig struct {
    MaxPageSize        int
    DefaultPageSize    int
    AllowedSortFields  map[string]bool
    AllowedFilterFields map[string]bool
    AllowedIncludes    map[string]bool
    AllowedFields      map[string][]string // map[resourceType][]allowedFields
}

// DefaultConfig returns a default configuration

const (
    DefaultPageSize = 10
 )

func DefaultConfig() request.QueryParams {
    return request.QueryParams{
        Page: request.PaginationQuery{
            PageSize: DefaultPageSize,
        },
    }
}


type ExtendedRequest struct {
    *http.Request
    QueryParams request.QueryParams
    Errors      []ValidationError
}

func (r *ExtendedRequest) HasErrors() bool {
    return len(r.Errors) > 0
}

func AsExtended(r *http.Request) *ExtendedRequest {
    if ext, ok := r.(*ExtendedRequest); ok {
        return ext
    }
    return nil
}

func validateQueryParams(r *ExtendedRequest, config QueryParamsConfig) {
    if r.QueryParams.Page.PageSize > config.MaxPageSize {
        r.Errors = append(r.Errors, ValidationError{
            Field:   "page[size]",
            Message: fmt.Sprintf("page size exceeds maximum allowed value of %d", config.MaxPageSize),
        })
        r.QueryParams.Page.PageSize = config.MaxPageSize
    }

    if len(config.AllowedSortFields) > 0 {
        for _, field := range r.QueryParams.Sort.Fields {
            cleanField := strings.TrimPrefix(field, "-")
            if !config.AllowedSortFields[cleanField] {
                r.Errors = append(r.Errors, ValidationError{
                    Field:   "sort",
                    Message: fmt.Sprintf("invalid sort field: %s", cleanField),
                })
            }
        }
    }

    if len(config.AllowedFilterFields) > 0 {
        for field := range r.QueryParams.Filter.Fields {
            if !config.AllowedFilterFields[field] {
                r.Errors = append(r.Errors, ValidationError{
                    Field:   "filter",
                    Message: fmt.Sprintf("invalid filter field: %s", field),
                })
            }
        }
    }

    if len(config.AllowedIncludes) > 0 {
        for _, include := range r.QueryParams.Include {
            if !config.AllowedIncludes[include] {
                r.Errors = append(r.Errors, ValidationError{
                    Field:   "include",
                    Message: fmt.Sprintf("invalid include: %s", include),
                })
            }
        }
    }

    if len(config.AllowedFields) > 0 {
        for resourceType, fields := range r.QueryParams.Fields {
            allowedFields, exists := config.AllowedFields[resourceType]
            if !exists {
                r.Errors = append(r.Errors, ValidationError{
                    Field:   "fields",
                    Message: fmt.Sprintf("invalid resource type: %s", resourceType),
                })
                continue
            }

            allowedFieldMap := make(map[string]bool)
            for _, f := range allowedFields {
                allowedFieldMap[f] = true
            }

            for _, field := range fields {
                if !allowedFieldMap[field] {
                    r.Errors = append(r.Errors, ValidationError{
                        Field:   fmt.Sprintf("fields[%s]", resourceType),
                        Message: fmt.Sprintf("invalid field: %s", field),
                    })
                }
            }
        }
    }
}
