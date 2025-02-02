package request

import (
	"github.com/midil-labs/core/shared/jsonapi/jsonapierror"
)

func NewJSONAPIRequest[T BodyType](path string, method string, opts ...RequestOption[T]) (JSONAPIRequest[T], jsonapierror.ErrorObjects) {

	var errs = jsonapierror.ErrorObjects{}

	req := JSONAPIRequest[T]{
		Path:   path,
		Method: method,
	}

	for _, opt := range opts {
		if err := opt(&req); err != nil {
			errs = append(errs, err...)
		}
	}

	return req, errs
}
