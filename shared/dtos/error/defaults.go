package error

import (
	"fmt"
	"net/http"
)

const (
	CodeUnauthorized           = "AUTH_UNAUTHORIZED"
	CodeInvalidCredentials     = "AUTH_INVALID_CREDENTIALS"
	CodeTokenExpired           = "AUTH_TOKEN_EXPIRED"
	CodeTokenInvalid           = "AUTH_TOKEN_INVALID"
	CodeInsufficientPermission = "AUTH_INSUFFICIENT_PERMISSION"
	CodeSessionExpired         = "AUTH_SESSION_EXPIRED"

	CodeValidationFailed      = "VALIDATION_FAILED"
	CodeInvalidInput          = "VALIDATION_INVALID_INPUT"
	CodeInvalidFormat         = "VALIDATION_INVALID_FORMAT"
	CodeMissingField          = "VALIDATION_MISSING_FIELD"
	CodeInvalidValue          = "VALIDATION_INVALID_VALUE"
	CodeTooLong               = "VALIDATION_TOO_LONG"
	CodeTooShort              = "VALIDATION_TOO_SHORT"
	CodeOutOfRange            = "VALIDATION_OUT_OF_RANGE"

	CodeNotFound             = "RESOURCE_NOT_FOUND"
	CodeAlreadyExists        = "RESOURCE_ALREADY_EXISTS"
	CodeConflict             = "RESOURCE_CONFLICT"
	CodeGone                 = "RESOURCE_GONE"
	CodeLocked               = "RESOURCE_LOCKED"

	CodeBusinessRule         = "BUSINESS_RULE_VIOLATION"
	CodeQuotaExceeded        = "BUSINESS_QUOTA_EXCEEDED"
	CodeRateLimitExceeded    = "BUSINESS_RATE_LIMIT_EXCEEDED"
	CodeInvalidState         = "BUSINESS_INVALID_STATE"
	CodeDependencyConflict   = "BUSINESS_DEPENDENCY_CONFLICT"

	CodeDatabaseError       = "INFRA_DATABASE_ERROR"
	CodeNetworkError        = "INFRA_NETWORK_ERROR"
	CodeServiceUnavailable  = "INFRA_SERVICE_UNAVAILABLE"
	CodeTimeout             = "INFRA_TIMEOUT"
	CodeInternalError       = "INFRA_INTERNAL_ERROR"
)

// Authentication/Authorization Errors
func NewUnauthorizedError(detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusUnauthorized,
		CodeUnauthorized,
		"Unauthorized",
		detail,
		opts...,
	)
}

func NewInvalidCredentialsError(detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusUnauthorized,
		CodeInvalidCredentials,
		"Invalid Credentials",
		detail,
		opts...,
	)
}

func NewTokenExpiredError(detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusUnauthorized,
		CodeTokenExpired,
		"Token Expired",
		detail,
		opts...,
	)
}

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
func NewValidationError(field, detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusUnprocessableEntity,
		CodeValidationFailed,
		"Validation Failed",
		detail,
		append(opts, WithSource("/data/attributes/"+field, "", ""))...,
	)
}

func NewMissingFieldError(field string, opts ...Option) *ErrorObject {
	return New(
		http.StatusBadRequest,
		CodeMissingField,
		"Missing Required Field",
		fmt.Sprintf("The field '%s' is required", field),
		append(opts, WithSource("/data/attributes/"+field, "", ""))...,
	)
}

func NewInvalidFormatError(field, expectedFormat string, opts ...Option) *ErrorObject {
	return New(
		http.StatusBadRequest,
		CodeInvalidFormat,
		"Invalid Format",
		fmt.Sprintf("The field '%s' must be in %s format", field, expectedFormat),
		append(opts, WithSource("/data/attributes/"+field, "", ""))...,
	)
}

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
func NewNotFoundError(detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusNotFound,
		CodeNotFound,
		"Resource Not Found",
		detail,
		opts...,
	)
}

func NewResourceAlreadyExistsError(detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusConflict,
		CodeAlreadyExists,
		"Resource Already Exists",
		detail,
		opts...,
	)
}

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
func NewQuotaExceededError(quota string, limit int, opts ...Option) *ErrorObject {
	return New(
		http.StatusForbidden,
		CodeQuotaExceeded,
		"Quota Exceeded",
		fmt.Sprintf("The %s quota limit of %d has been exceeded", quota, limit),
		opts...,
	)
}

func NewRateLimitExceededError(window string, opts ...Option) *ErrorObject {
	return New(
		http.StatusTooManyRequests,
		CodeRateLimitExceeded,
		"Rate Limit Exceeded",
		fmt.Sprintf("Rate limit exceeded. Please try again in %s", window),
		opts...,
	)
}

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
func NewDatabaseError(detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusInternalServerError,
		CodeDatabaseError,
		"Database Error",
		detail,
		opts...,
	)
}

func NewNetworkError(service string, opts ...Option) *ErrorObject {
	return New(
		http.StatusBadGateway,
		CodeNetworkError,
		"Network Error",
		fmt.Sprintf("Error communicating with %s", service),
		opts...,
	)
}

func NewTimeoutError(service string, opts ...Option) *ErrorObject {
	return New(
		http.StatusGatewayTimeout,
		CodeTimeout,
		"Timeout",
		fmt.Sprintf("Request to %s timed out", service),
		opts...,
	)
}

func NewServiceUnavailableError(service string, opts ...Option) *ErrorObject {
	return New(
		http.StatusServiceUnavailable,
		CodeServiceUnavailable,
		"Service Unavailable",
		fmt.Sprintf("The %s service is currently unavailable", service),
		opts...,
	)
}

// Helper for generic internal errors
func NewInternalError(detail string, opts ...Option) *ErrorObject {
	return New(
		http.StatusInternalServerError,
		CodeInternalError,
		"Internal Server Error",
		detail,
		opts...,
	)
}