package utils

import (
	"demo/bank-linking-listener/internal/service/entity"
	"net/http"
)

const (
	ResponseStatusSuccess  = "SUCCESSFUL"
	ResponseStatusFail    = "FAILED"
	ResponseStatusProcess = "PROCESSING"
	ResponseStatusPending = "PENDING"
)

func ResponseWithData(status string, data map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status": status,
		"data":   data,
	}
}

func ResponseWithMessage(status string, message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message,
	}
}

func GetStatusCode(rerr entity.Error) int {
	if rerr == nil {
		return http.StatusOK
	}

	// 400 Bad Request
	if rerr.ErrorType() == entity.ErrorInvalidArgument ||
		rerr.ErrorType() == entity.ErrorFaildPrecondition ||
		rerr.ErrorType() == entity.ErrorOutOfRange {
		return http.StatusBadRequest
	}

	// 401 Unauthorized
	if rerr.ErrorType() == entity.ErrorUnauthenticated {
		return http.StatusUnauthorized
	}

	// 403 Forbidden
	if rerr.ErrorType() == entity.ErrorPermissionDenied {
		return http.StatusForbidden
	}

	// 404 Not Found
	if rerr.ErrorType() == entity.ErrorNotFound {
		return http.StatusNotFound
	}

	// 409 Conflict
	if rerr.ErrorType() == entity.ErrorAlreadyExists ||
		rerr.ErrorType() == entity.ErrorAborted {
		return http.StatusConflict
	}

	// 429 Too Many Request
	if rerr.ErrorType() == entity.ErrorResourceExhausted {
		return http.StatusTooManyRequests
	}

	// 499 Client Closed Request
	if rerr.ErrorType() == entity.ErrorCancelled {
		return 499
	}

	// 500 Internal Server Error
	if rerr.ErrorType() == entity.ErrorUnknown ||
		rerr.ErrorType() == entity.ErrorInternal ||
		rerr.ErrorType() == entity.ErrorDataLoss {
		return http.StatusInternalServerError
	}

	// 501 Not Implemented
	if rerr.ErrorType() == entity.ErrorUnimplemented {
		return http.StatusNotImplemented
	}

	// 503 Service Unavailable
	if rerr.ErrorType() == entity.ErrorUnavailable {
		return http.StatusServiceUnavailable
	}

	// 504 Gateway Timeout
	if rerr.ErrorType() == entity.ErrorDeadlineExceeded {
		return http.StatusGatewayTimeout
	}

	// Default to Internal Server Error
	return http.StatusInternalServerError
}
