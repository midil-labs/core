package errors

import ("fmt"
		"net/http"
	    "github.com/midil-labs/core/shared/dtos/response"
		"strings")

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%d - %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}

func NewAppError(code int, message string, err error, ) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}


func BadRequest(detail string, err error) *response.ErrorResponse {
	errorObject := response.ErrorObject{
		Code:   fmt.Sprintf("%d", http.StatusBadRequest),
		Title:  http.StatusText(http.StatusBadRequest),
		Detail: detail,
	}
	errors := []response.ErrorObject{errorObject}
	
	return response.NewErrorResponse(errors)
}

func NotFound(message string, err error) *AppError {
	return NewAppError(http.StatusNotFound, message, err)
}

func InternalServerError(message string, err error) *AppError {
	return NewAppError(http.StatusInternalServerError, message, err)
}

func Unauthorized(message string, err error) *AppError {
	return NewAppError(http.StatusUnauthorized, message, err)
}

func Forbidden(message string, err error) *AppError {
	return NewAppError(http.StatusForbidden, message, err)
}

func Conflict(message string, err error) *AppError {
	return NewAppError(http.StatusConflict, message, err)
}

func UnprocessableEntity(message string, err error) *AppError {
	return NewAppError(http.StatusUnprocessableEntity, message, err)
}
