package error

// Error codes
const (
	CodeUnauthorized           ErrCode = "AUTH_UNAUTHORIZED"
	CodeInvalidCredentials     ErrCode = "AUTH_INVALID_CREDENTIALS"
	CodeTokenExpired           ErrCode = "AUTH_TOKEN_EXPIRED"
	CodeTokenInvalid           ErrCode = "AUTH_TOKEN_INVALID"
	CodeInsufficientPermission ErrCode = "AUTH_INSUFFICIENT_PERMISSION"
	CodeSessionExpired         ErrCode = "AUTH_SESSION_EXPIRED"

	CodeValidationFailed ErrCode = "VALIDATION_FAILED"
	CodeInvalidInput     ErrCode = "VALIDATION_INVALID_INPUT"
	CodeInvalidFormat    ErrCode = "VALIDATION_INVALID_FORMAT"
	CodeMissingField     ErrCode = "VALIDATION_MISSING_FIELD"
	CodeInvalidValue     ErrCode = "VALIDATION_INVALID_VALUE"
	CodeTooLong          ErrCode = "VALIDATION_TOO_LONG"
	CodeTooShort         ErrCode = "VALIDATION_TOO_SHORT"
	CodeOutOfRange       ErrCode = "VALIDATION_OUT_OF_RANGE"

	CodeNotFound      ErrCode = "RESOURCE_NOT_FOUND"
	CodeAlreadyExists ErrCode = "RESOURCE_ALREADY_EXISTS"
	CodeConflict      ErrCode = "RESOURCE_CONFLICT"
	CodeGone          ErrCode = "RESOURCE_GONE"
	CodeLocked        ErrCode = "RESOURCE_LOCKED"

	CodeBusinessRule       ErrCode = "BUSINESS_RULE_VIOLATION"
	CodeQuotaExceeded      ErrCode = "BUSINESS_QUOTA_EXCEEDED"
	CodeRateLimitExceeded  ErrCode = "BUSINESS_RATE_LIMIT_EXCEEDED"
	CodeInvalidState       ErrCode = "BUSINESS_INVALID_STATE"
	CodeDependencyConflict ErrCode = "BUSINESS_DEPENDENCY_CONFLICT"

	CodeDatabaseError      ErrCode = "INFRA_DATABASE_ERROR"
	CodeNetworkError       ErrCode = "INFRA_NETWORK_ERROR"
	CodeServiceUnavailable ErrCode = "INFRA_SERVICE_UNAVAILABLE"
	CodeTimeout            ErrCode = "INFRA_TIMEOUT"
	CodeInternalError      ErrCode = "INFRA_INTERNAL_ERROR"

	CodeUnknownError ErrCode = "UNKNOWN_ERROR"
)

// Language represents a language code
const (
	EN Language = "en"
	ES Language = "es"
	FR Language = "fr"
)
