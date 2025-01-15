package http

import (
	"net/http"
	"github.com/midil-labs/core/shared/jsonapi/request"
)


func NewHTTPRequest[T request.RequestType](r *http.Request) request.JSONAPIRequest[T] {

	req := request.NewJSONAPIRequest[T](r.URL.Path, r.Method,
		request.WithQuery[T](r.URL.Query()),
		request.WithHeaders[T](r.Header),
		// WithBody[T](r),
	)

	return req
}
