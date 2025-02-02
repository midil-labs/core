package jsonapierror

import (
	"fmt"
	"net/http"
)

// Authentication/Authorization Errors

// NewUnauthorizedError creates a new ErrorObject representing an unauthorized error.
// It sets the HTTP status to 401 Unauthorized and uses the CodeUnauthorized code.
// The error message includes a detailed description of the unauthorized error.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - detail: A string providing additional details about the error.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewUnauthorizedError(detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusUnauthorized,
		CodeUnauthorized,
		"Unauthorized",
		detail,
		opts...,
	)
}

// NewForbiddenError creates a new ErrorObject representing a forbidden error.
// It sets the HTTP status to 403 Forbidden and uses the CodeForbidden code.
// The error message includes a detailed description of the forbidden error.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - detail: A string providing additional details about the error.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewInvalidCredentialsError(detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusUnauthorized,
		CodeInvalidCredentials,
		"Invalid Credentials",
		detail,
		opts...,
	)
}

// NewTokenExpiredError creates a new ErrorObject representing a token expired error.
// It sets the HTTP status to 401 Unauthorized and uses the CodeTokenExpired code.
// The error message includes a detailed description of the token expired error.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - detail: A string providing additional details about the error.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewTokenExpiredError(detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusUnauthorized,
		CodeTokenExpired,
		"Token Expired",
		detail,
		opts...,
	)
}

// NewInsufficientPermissionError creates a new ErrorObject representing an insufficient permission error.
// It sets the HTTP status to 403 Forbidden and uses the CodeInsufficientPermission code.
// The error message includes a detailed description of the insufficient permission error.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - detail: A string providing additional details about the error.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewInsufficientPermissionError(detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusForbidden,
		CodeInsufficientPermission,
		"Insufficient Permission",
		detail,
		opts...,
	)
}

// Validation Errors
// 422 Unprocessable Entity: The format is correct, but the data is semantically invalid.
// 400 Bad Request: The query parameters are invalid, and the server cannot process the request because it is malformed or includes invalid parameters.

func NewValidationError(source ErrorSource, detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusUnprocessableEntity,
		CodeValidationFailed,
		"Validation Failed",
		detail,
		append(opts, WithSource(source.Pointer, source.Parameter, source.Header))...,
	)
}

// NewInvalidFieldError creates a new ErrorObject representing an invalid field error.
// It sets the HTTP status to 400 Bad Request and uses the CodeValidationFailed code.
// The error message includes the field pointer and a detailed message.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - pointer: A string representing the field that is invalid.
//   - detail: A string providing details about why the field is invalid.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewInvalidFieldError(pointer, detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusBadRequest,
		CodeValidationFailed,
		"Invalid Field",
		fmt.Sprintf("The field '%s' is invalid: %s", pointer, detail),
		append(opts, WithSource("/data/attributes/"+pointer, "", ""))...,
	)
}

// NewUnprocessibleFieldError creates a new ErrorObject representing an unprocessible field error.
// It sets the HTTP status to 422 Unprocessable Entity and uses the CodeValidationFailed code.
// The error message includes the field pointer and a detailed message.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - pointer: A string representing the field that is unprocessible.
//   - detail: A string providing details about why the field is unprocessible.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewUnprocessibleFieldError(pointer, detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusUnprocessableEntity,
		CodeValidationFailed,
		"Unprocessable Field",
		fmt.Sprintf("The field '%s' is unprocessable: %s", pointer, detail),
		append(opts, WithSource("/data/attributes/"+pointer, "", ""))...,
	)
}

// NewInvalidQueryError creates a new ErrorObject representing an invalid query parameter error.
// It sets the HTTP status to 400 Bad Request and uses the CodeValidationFailed code.
// The error message includes the query parameter name and a detailed message.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - parameter: A string representing the query parameter that is invalid.
//   - detail: A string providing details about why the query parameter is invalid.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewInvalidQueryError(parameter, detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusBadRequest,
		CodeValidationFailed,
		"Invalid Query Parameter",
		fmt.Sprintf("The query parameter '%s' is invalid: %s", parameter, detail),
		append(opts, WithSource("", parameter, ""))...,
	)
}

// NewUnprocessableQueryError creates a new ErrorObject representing an unprocessible query parameter error.
// It sets the HTTP status to 422 Unprocessable Entity and uses the CodeValidationFailed code.
// The error message includes the query parameter name and a detailed message.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - parameter: A string representing the query parameter that is unprocessible.
//   - detail: A string providing details about why the query parameter is unprocessible.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewUnprocessableQueryError(parameter, detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusUnprocessableEntity,
		CodeValidationFailed,
		"Unprocessable Query Parameter",
		fmt.Sprintf("The query parameter '%s' is unprocessable: %s", parameter, detail),
		append(opts, WithSource("", parameter, ""))...,
	)
}

// NewInvalidHeaderError creates a new ErrorObject representing an invalid header error.
// It sets the HTTP status to 422 Unprocessable Entity and uses the CodeValidationFailed code.
// The error message includes the header name and a detailed message.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - header: A string representing the header that is invalid.
//   - detail: A string providing details about why the header is invalid.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewInvalidHeaderError(header, detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusBadRequest,
		CodeValidationFailed,
		"Invalid Header",
		fmt.Sprintf("The header '%s' is invalid: %s", header, detail),
		append(opts, WithSource("", "", header))...,
	)
}

// NewUnprocessableHeaderError creates a new ErrorObject representing an unprocessible header error.
// It sets the HTTP status to 422 Unprocessable Entity and uses the CodeValidationFailed code.
// The error message includes the header name and a detailed message.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - header: A string representing the header that is unprocessible.
//   - detail: A string providing details about why the header is unprocessible.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewUnprocessableHeaderError(header, detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusUnprocessableEntity,
		CodeValidationFailed,
		"Unprocessable Header",
		fmt.Sprintf("The header '%s' is unprocessable: %s", header, detail),
		append(opts, WithSource("", "", header))...,
	)
}

// NewMissingFieldError creates a new ErrorObject representing a missing field error.
// It sets the HTTP status to 400 Bad Request and uses the CodeMissingField code.
// The error message includes the field name and a detailed message.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - field: A string representing the missing field.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewMissingFieldError(field string, opts ...Option) *ErrorObject {
	return New(
		http.StatusBadRequest,
		CodeMissingField,
		"Missing Required Field",
		fmt.Sprintf("The field '%s' is required", field),
		append(opts, WithSource("/data/attributes/"+field, "", ""))...,
	)
}

// NewInvalidFormatError creates a new ErrorObject representing an invalid format error.
// It sets the HTTP status to 400 Bad Request and uses the CodeInvalidFormat code.
// The error message includes the field name, the expected format, and a detailed message.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - field: A string representing the field with the invalid format.
//   - expectedFormat: A string representing the expected format for the field.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewInvalidFormatError(field, expectedFormat string, opts ...Option) *ErrorObject {
	return New(
		http.StatusBadRequest,
		CodeInvalidFormat,
		"Invalid Format",
		fmt.Sprintf("The field '%s' must be in %s format", field, expectedFormat),
		append(opts, WithSource("/data/attributes/"+field, "", ""))...,
	)
}

// NewTooLongError creates a new ErrorObject representing a field too long error.
// It sets the HTTP status to 400 Bad Request and uses the CodeTooLong code.
// The error message includes the field name and a detailed message.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - detail: A string providing additional details about the error.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewTooLongError(detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusBadRequest,
		CodeTooLong,
		"Field Too Long",
		detail,
		opts...,
	)
}

// Resource Errors

// NewNotFoundError creates a new ErrorObject indicating that a resource was not found.
// It sets the HTTP status to 404 Not Found and uses the CodeNotFound code.
// The error message includes a detailed description of the resource that was not found.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - detail: A string providing additional details about the error.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewNotFoundError(detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusNotFound,
		CodeNotFound,
		"Resource Not Found",
		detail,
		opts...,
	)
}

// NewResourceAlreadyExistsError creates a new ErrorObject indicating that a resource already exists.
// It sets the HTTP status to 409 Conflict and uses the CodeAlreadyExists code.
// The error message includes a detailed description of the resource that already exists.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - detail: A string providing additional details about the error.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewResourceAlreadyExistsError(detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusConflict,
		CodeAlreadyExists,
		"Resource Already Exists",
		detail,
		opts...,
	)
}

// NewResourceLockedError creates a new ErrorObject indicating that a resource is locked.
// It sets the HTTP status to 423 Locked and uses the CodeLocked code.
// The error message includes a detailed description of the locked resource.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - detail: A string providing additional details about the error.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewResourceLockedError(detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusLocked,
		CodeLocked,
		"Resource Locked",
		detail,
		opts...,
	)
}

// Business Logic Errors

// NewQuotaExceededError creates a new ErrorObject indicating that a quota limit has been exceeded.
// It sets the HTTP status to 403 Forbidden and uses the CodeQuotaExceeded code.
// The error message includes the quota name and the limit that was exceeded.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - quota: A string representing the quota that was exceeded.
//   - limit: An integer representing the limit that was exceeded.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewQuotaExceededError(quota string, limit int, opts ...Option) *ErrorObject {
	return New(
		http.StatusForbidden,
		CodeQuotaExceeded,
		"Quota Exceeded",
		fmt.Sprintf("The %s quota limit of %d has been exceeded", quota, limit),
		opts...,
	)
}

// NewRateLimitExceededError creates a new ErrorObject indicating that a rate limit has been exceeded.
// It sets the HTTP status to 429 Too Many Requests and uses the CodeRateLimitExceeded code.
// The error message includes the time window when the rate limit will reset.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - window: A string representing the time window when the rate limit will reset.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewRateLimitExceededError(window string, opts ...Option) *ErrorObject {
	return New(
		http.StatusTooManyRequests,
		CodeRateLimitExceeded,
		"Rate Limit Exceeded",
		fmt.Sprintf("Rate limit exceeded. Please try again in %s", window),
		opts...,
	)
}

// NewInvalidStateError creates a new ErrorObject indicating that a resource is in an invalid state.
// It sets the HTTP status to 409 Conflict and uses the CodeInvalidState code.
// The error message includes the resource name, the current state, and the required state.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - resource: A string representing the resource that is in an invalid state.
//   - currentState: A string representing the current state of the resource.
//   - requiredState: A string representing the required state of the resource.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewInvalidStateError(resource, currentState, requiredState string, opts ...Option) *ErrorObject {
	return New(
		http.StatusConflict,
		CodeInvalidState,
		"Invalid State",
		fmt.Sprintf("%s is in %s state. Required state: %s", resource, currentState, requiredState),
		opts...,
	)
}

// Infrastructure Errors

// NewDatabaseError creates a new ErrorObject indicating a database error.
// It sets the HTTP status to 500 Internal Server Error and uses the CodeDatabaseError code.
// The error message includes a detailed description of the database error.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - detail: A string providing additional details about the database error.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewDatabaseError(detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusInternalServerError,
		CodeDatabaseError,
		"Database Error",
		detail,
		opts...,
	)
}

// NewNetworkError creates a new ErrorObject indicating a network error.
// It sets the HTTP status to 502 Bad Gateway and uses the CodeNetworkError code.
// The error message includes a detailed description of the network error.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - service: A string representing the service that caused the network error.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewNetworkError(service string, opts ...Option) *ErrorObject {
	return New(
		http.StatusBadGateway,
		CodeNetworkError,
		"Network Error",
		fmt.Sprintf("Error communicating with %s", service),
		opts...,
	)
}

// NewTimeoutError creates a new ErrorObject indicating a timeout error.
// It sets the HTTP status to 504 Gateway Timeout and uses the CodeTimeout code.
// The error message includes a detailed description of the timeout error.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - service: A string representing the service that caused the timeout error.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewTimeoutError(service string, opts ...Option) *ErrorObject {
	return New(
		http.StatusGatewayTimeout,
		CodeTimeout,
		"Timeout",
		fmt.Sprintf("Request to %s timed out", service),
		opts...,
	)
}

// NewServiceUnavailableError creates a new ErrorObject indicating that a service is unavailable.
// It sets the HTTP status to 503 Service Unavailable and uses the CodeServiceUnavailable code.
// The error message includes a detailed description of the unavailable service.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - service: A string representing the service that is unavailable.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewServiceUnavailableError(service string, opts ...Option) *ErrorObject {
	return New(
		http.StatusServiceUnavailable,
		CodeServiceUnavailable,
		"Service Unavailable",
		fmt.Sprintf("The %s service is currently unavailable", service),
		opts...,
	)
}

// NewInternalError creates a new ErrorObject indicating an internal server error.
// It sets the HTTP status to 500 Internal Server Error and uses the CodeInternalError code.
// The error message includes a detailed description of the internal server error.
// Additional options can be provided via the opts parameter.
//
// Parameters:
//   - detail: A string providing additional details about the internal server error.
//   - opts: Additional options to customize the ErrorObject.
//
// Returns:
//   - *ErrorObject: A pointer to the newly created ErrorObject.
func NewInternalError(detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusInternalServerError,
		CodeInternalError,
		"Internal Server Error",
		detail,
		opts...,
	)
}
