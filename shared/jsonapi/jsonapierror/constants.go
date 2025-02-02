package jsonapierror

import "net/http"


//Authentication & Authorization Errors
const (
	CodeUnauthorized           ErrCode = "AUTH_UNAUTHORIZED"
	CodeInvalidCredentials     ErrCode = "AUTH_INVALID_CREDENTIALS"
	CodeTokenExpired           ErrCode = "AUTH_TOKEN_EXPIRED"
	CodeTokenInvalid           ErrCode = "AUTH_TOKEN_INVALID"
	CodeInsufficientPermission ErrCode = "AUTH_INSUFFICIENT_PERMISSION"
	CodeSessionExpired         ErrCode = "AUTH_SESSION_EXPIRED"
)

//Validation Errors
const(
	CodeValidationFailed ErrCode = "VALIDATION_FAILED"
	CodeInvalidInput     ErrCode = "VALIDATION_INVALID_INPUT"
	CodeInvalidFormat    ErrCode = "VALIDATION_INVALID_FORMAT"
	CodeMissingField     ErrCode = "VALIDATION_MISSING_FIELD"
	CodeInvalidValue     ErrCode = "VALIDATION_INVALID_VALUE"
	CodeTooLong          ErrCode = "VALIDATION_TOO_LONG"
	CodeTooShort         ErrCode = "VALIDATION_TOO_SHORT"
	CodeOutOfRange       ErrCode = "VALIDATION_OUT_OF_RANGE"
	CodeMalformedRequest ErrCode = "VALIDATION_MALFORMED_REQUEST"
)

// Resource Errors
const(
	CodeNotFound      ErrCode = "RESOURCE_NOT_FOUND"
	CodeAlreadyExists ErrCode = "RESOURCE_ALREADY_EXISTS"
	CodeConflict      ErrCode = "RESOURCE_CONFLICT"
	CodeGone          ErrCode = "RESOURCE_GONE"
	CodeLocked        ErrCode = "RESOURCE_LOCKED"
)

// Business Logic Errors
const(
	CodeBusinessRuleViolation ErrCode = "BUSINESS_RULE_VIOLATION"
	CodeQuotaExceeded         ErrCode = "BUSINESS_QUOTA_EXCEEDED"
	CodeRateLimitExceeded     ErrCode = "BUSINESS_RATE_LIMIT_EXCEEDED"
	CodeInvalidState          ErrCode = "BUSINESS_INVALID_STATE"
	CodeDependencyConflict    ErrCode = "BUSINESS_DEPENDENCY_CONFLICT"
)

//Server Errors
const(
	CodeDatabaseError      ErrCode = "SERVER_DATABASE_ERROR"
	CodeNetworkError       ErrCode = "SERVER_NETWORK_ERROR"
	CodeServiceUnavailable ErrCode = "SERVER_SERVICE_UNAVAILABLE"
	CodeTimeout            ErrCode = "SERVER_TIMEOUT"
	CodeInternalError      ErrCode = "SERVER_ERROR"
)

// Unknown Errors
const(
	CodeUnknownError ErrCode = "UNKNOWN_ERROR"
)

// Supported languages
const (
	EN Language = "en"
	ES Language = "es"
	FR Language = "fr"
)

// ErrorMessage holds language-specific error messages.
type ErrorMessage struct {
	EN string
	ES string
	FR string
}

// ErrorMessages maps error codes to their language-specific messages.
var ErrorMessages = map[ErrCode]ErrorMessage{
	// --- Authentication & Authorization Errors ---
	CodeUnauthorized: {
		EN: "Unauthorized access",
		ES: "Acceso no autorizado",
		FR: "Accès non autorisé",
	},
	CodeInvalidCredentials: {
		EN: "Invalid credentials",
		ES: "Credenciales inválidas",
		FR: "Identifiants invalides",
	},
	CodeTokenExpired: {
		EN: "Token has expired",
		ES: "El token ha expirado",
		FR: "Le jeton a expiré",
	},
	CodeInsufficientPermission: {
		EN: "Insufficient permissions",
		ES: "Permisos insuficientes",
		FR: "Permissions insuffisantes",
	},
	CodeSessionExpired: {
		EN: "Session has expired",
		ES: "La sesión ha expirado",
		FR: "La session a expiré",
	},

	// --- Validation Errors ---
	CodeValidationFailed: {
		EN: "Validation failed",
		ES: "Validación fallida",
		FR: "Échec de la validation",
	},
	CodeInvalidInput: {
		EN: "Invalid input",
		ES: "Entrada inválida",
		FR: "Entrée invalide",
	},
	CodeMalformedRequest: {
		EN: "Malformed request",
		ES: "Solicitud malformada",
		FR: "Requête malformée",
	},
	CodeInvalidFormat: {
		EN: "Invalid format",
		ES: "Formato inválido",
		FR: "Format invalide",
	},

	CodeMissingField: {
		EN: "Missing field",
		ES: "Campo faltante",
		FR: "Champ manquant",
	},
	CodeInvalidValue: {
		EN: "Invalid value",
		ES: "Valor inválido",
		FR: "Valeur invalide",
	},
	CodeTooLong: {
		EN: "Value is too long",
		ES: "El valor es demasiado largo",
		FR: "La valeur est trop longue",
	},
	CodeTooShort: {
		EN: "Value is too short",
		ES: "El valor es demasiado corto",
		FR: "La valeur est trop courte",
	},
	// --- Resource Errors ---
	CodeNotFound: {
		EN: "Resource not found",
		ES: "Recurso no encontrado",
		FR: "Ressource non trouvée",
	},
	CodeAlreadyExists: {
		EN: "Resource already exists",
		ES: "El recurso ya existe",
		FR: "La ressource existe déjà",
	},
	CodeConflict: {
		EN: "Resource conflict",
		ES: "Conflicto de recursos",
		FR: "Conflit de ressources",
	},
	CodeGone: {
		EN: "Resource is gone",
		ES: "El recurso ha desaparecido",
		FR: "La ressource a disparu",
	},
	CodeLocked: {
		EN: "Resource is locked",
		ES: "El recurso está bloqueado",
		FR: "La ressource est verrouillée",
	},
	// --- Business Logic Errors ---
	CodeQuotaExceeded: {
		EN: "Quota exceeded",
		ES: "Cuota excedida",
		FR: "Quota dépassé",
	},
	CodeRateLimitExceeded: {
		EN: "Rate limit exceeded",
		ES: "Límite de velocidad excedido",
		FR: "Limite de taux dépassé",
	},
	CodeInvalidState: {
		EN: "Invalid state",
		ES: "Estado inválido",
		FR: "État invalide",
	},
	CodeDependencyConflict: {
		EN: "Dependency conflict",
		ES: "Conflicto de dependencia",
		FR: "Conflit de dépendance",
	},
	CodeBusinessRuleViolation: {
		EN: "Business rule violation",
		ES: "Violación de reglas de negocio",
		FR: "Violation de règle métier",
	},
	// --- Server Errors ---
	CodeInternalError: {
		EN: "Internal server error",
		ES: "Error interno del servidor",
		FR: "Erreur interne du serveur",
	},
	CodeDatabaseError: {
		EN: "Database error",
		ES: "Error de base de datos",
		FR: "Erreur de base de données",
	},
	CodeNetworkError: {
		EN: "Network error",
		ES: "Error de red",
		FR: "Erreur réseau",
	},
	CodeServiceUnavailable: {
		EN: "Service unavailable",
		ES: "Servicio no disponible",
		FR: "Service indisponible",
	},
	CodeTimeout: {
		EN: "Request timeout",
		ES: "Tiempo de espera agotado",
		FR: "Délai d'attente dépassé",
	},
	// --- Unknown Errors ---
	CodeUnknownError: {
		EN: "Unknown error",
		ES: "Error desconocido",
		FR: "Erreur inconnue",
	},
}

// ErrorCodeToHTTPStatus maps error codes to HTTP status codes.
var ErrorCodeToHTTPStatus = map[ErrCode]int{
	// --- Authentication & Authorization Errors ---
	CodeUnauthorized:           http.StatusUnauthorized,
	CodeInvalidCredentials:     http.StatusUnauthorized,
	CodeTokenExpired:           http.StatusUnauthorized,
	CodeTokenInvalid:           http.StatusUnauthorized,
	CodeInsufficientPermission: http.StatusForbidden,
	CodeSessionExpired:         http.StatusForbidden,

	// --- Validation Errors ---
	CodeValidationFailed: http.StatusUnprocessableEntity,
	CodeInvalidInput:     http.StatusUnprocessableEntity,
	CodeInvalidFormat:    http.StatusUnprocessableEntity,
	CodeMissingField:     http.StatusUnprocessableEntity,
	CodeInvalidValue:     http.StatusUnprocessableEntity,
	CodeTooLong:          http.StatusUnprocessableEntity,
	CodeTooShort:         http.StatusUnprocessableEntity,
	CodeOutOfRange:       http.StatusUnprocessableEntity,
	CodeMalformedRequest: http.StatusBadRequest,

	// --- Resource Errors ---
	CodeNotFound:      http.StatusNotFound,
	CodeAlreadyExists: http.StatusConflict,
	CodeConflict:      http.StatusConflict,
	CodeGone:          http.StatusGone,
	CodeLocked:        http.StatusLocked,

	// --- Business Logic Errors ---
	CodeBusinessRuleViolation: http.StatusUnprocessableEntity,
	CodeQuotaExceeded:         http.StatusTooManyRequests,
	CodeRateLimitExceeded:     http.StatusTooManyRequests,
	CodeInvalidState:          http.StatusUnprocessableEntity,
	CodeDependencyConflict:    http.StatusConflict,

	// --- Server Errors ---
	CodeDatabaseError:      http.StatusInternalServerError,
	CodeNetworkError:       http.StatusInternalServerError,
	CodeServiceUnavailable: http.StatusServiceUnavailable,
	CodeTimeout:            http.StatusGatewayTimeout,
	CodeInternalError:      http.StatusInternalServerError,

	// --- Unknown Errors ---
	CodeUnknownError: http.StatusInternalServerError,
}
