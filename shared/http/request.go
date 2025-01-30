package http

import (
	"net/http"
	"github.com/midil-labs/core/shared/jsonapi/request"
)


func NewHTTPRequest[T request.BodyType](r *http.Request) request.JSONAPIRequest[T] {

	req := request.NewJSONAPIRequest[T](
		r.URL.Path, r.Method,
		request.WithQuery[T](r.URL.Query()),
		request.WithHeaders[T](r.Header),
		request.WithBody[T](r.Body),
	)

	return req
}
