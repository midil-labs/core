package middleware

import (
	"context"
	"net/http"

	"github.com/midil-labs/core/shared/jsonapi/request"
	"github.com/midil-labs/core/shared/jsonapi/error"
    jhttp "github.com/midil-labs/core/shared/http"

)

type contextKey string

// JSONAPIContextKey is the key used to store JSONAPIRequest in the request context.
const JSONAPIRequestContextKey contextKey = "JSONAPIREQUESTCONTEXT"

func WithJSONAPIContext[T request.BodyType](callback func(request.JSONAPIRequest[T]) error.ErrorObjects) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jsonAPIRequest := jhttp.NewHTTPRequest[T](r)
			if callback != nil {
				if err := callback(jsonAPIRequest); err != nil {
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
func JSONAPIContext[T request.BodyType](r *http.Request) (request.JSONAPIRequest[T], bool) {
	spec, ok := r.Context().Value(JSONAPIRequestContextKey).(request.JSONAPIRequest[T])
	return spec, ok
}
