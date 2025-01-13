package request

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