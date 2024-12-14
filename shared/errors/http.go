package errors

import ("fmt"
		"net/http")

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

func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func BadRequest(message string, err error) *AppError {
	return NewAppError(http.StatusBadRequest, message, err)
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
