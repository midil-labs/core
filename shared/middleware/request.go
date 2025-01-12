package middleware

import (
	"context"
	"net/http"

	"github.com/midil-labs/core/shared/dtos/request"
	"github.com/midil-labs/core/shared/dtos/error"
    jhttp "github.com/midil-labs/core/shared/http"

)

type contextKey string

// JSONAPIContextKey is the key used to store JSONAPIRequest in the request context.
const JSONAPIRequestContextKey contextKey = "JSONAPIREQUESTCONTEXT"

func WithJSONAPIContext[T request.RequestType](validateFunc func(request.JSONAPIRequest[T]) []error.ErrorObject) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jsonAPIRequest := request.NewHTTPRequest[T](r)
			if validateFunc != nil {
				if err := validateFunc(jsonAPIRequest); err != nil {
                    response := *error.NewJSONAPIError(nil, err...)
                    jhttp.UnprocessableEntity(response)
					return
				}
			}
			ctx := context.WithValue(r.Context(), JSONAPIRequestContextKey, jsonAPIRequest)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// JSONAPIContext retrieves JSONAPIRequest from the request context.
// It returns the JSONAPIRequest and a boolean indicating if it was found.
func JSONAPIContext[T request.RequestType](r *http.Request) (request.JSONAPIRequest[T], bool) {
	spec, ok := r.Context().Value(JSONAPIRequestContextKey).(request.JSONAPIRequest[T])
	return spec, ok
}
