package errors

import (
	"net/http"
	"encoding/json"
	"github.com/midil-labs/core/shared/dtos/response"
)

// ResponseWriter is a type alias for a function that takes an http.ResponseWriter
// and returns an error. It is used to define a function signature for handling
// HTTP responses.
type ResponseWriter = func(w http.ResponseWriter) error

// The provided code must be a valid HTTP 2xx status code.
func Success[T response.DataType](w http.ResponseWriter, statusCode int, response response.JSONAPIResponse[T]) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
	return nil
}

// The provided status code must be a valid HTTP 4xx-5xx status code.
func Error(w http.ResponseWriter, statusCode int, response response.ErrorResponse) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
	return nil
}

func NotFound(response response.ErrorResponse) ResponseWriter {
	return func(w http.ResponseWriter) error {
		return Error(w, http.StatusNotFound, response)
	}
}

func InternalServerError(response response.ErrorResponse) ResponseWriter {
	return func(w http.ResponseWriter) error {
		return Error(w, http.StatusInternalServerError, response)
	}
}

func Unauthorized(response response.ErrorResponse) ResponseWriter {
	return func(w http.ResponseWriter) error {
		return Error(w, http.StatusUnauthorized, response)
	}
}

func Forbidden(response response.ErrorResponse) ResponseWriter {
	return func(w http.ResponseWriter) error {
		return Error(w, http.StatusForbidden, response)
	}
}

func Conflict(response response.ErrorResponse) ResponseWriter {
	return func(w http.ResponseWriter) error {
		return Error(w, http.StatusConflict, response)
	}
}

func UnprocessableEntity(response response.ErrorResponse) ResponseWriter {
	return func(w http.ResponseWriter) error {
		return Error(w, http.StatusUnprocessableEntity, response)
	}
}

func BadRequest(response response.ErrorResponse) ResponseWriter {
	return func(w http.ResponseWriter) error {
		return Error(w, http.StatusBadRequest, response)
	}
}

func OK[T response.DataType](response response.JSONAPIResponse[T]) ResponseWriter {
	return func(w http.ResponseWriter) error {
		return Success(w, http.StatusOK, response)
	}
}

func Created[T response.DataType](response response.JSONAPIResponse[T]) ResponseWriter {
	return func(w http.ResponseWriter) error {
		return Success(w, http.StatusCreated, response)
	}
}

func Accepted[T response.DataType](response response.JSONAPIResponse[T]) ResponseWriter {
	return func(w http.ResponseWriter) error {
		return Success(w, http.StatusAccepted, response)
	}
}
