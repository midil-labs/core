package request

import (
	"net/http"
)

func NewJSONAPIRequest[T RequestType](path string, method string, opts ...RequestOption[T]) JSONAPIRequest[T] {
	req := JSONAPIRequest[T]{
		Path:   path,
		Method: method,
	}
	for _, opt := range opts {
		opt(&req)
	}
	return req
}

func NewHTTPRequest[T RequestType](r *http.Request) JSONAPIRequest[T] {

	req := NewJSONAPIRequest[T](r.URL.Path, r.Method,
		WithQuery[T](r),
		WithHeaders[T](r),
		// WithBody[T](r),
	)

	return req
}
