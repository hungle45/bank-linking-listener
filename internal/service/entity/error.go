package entity

type Error interface {
	Message() string
	ErrorType() ErrorType
	Error()
}

type errorImp struct {
	errorType ErrorType
	message   string
}

func (r *errorImp) Error() {
	// Noncompliant
}

func (r *errorImp) Message() string {
	return r.message
}

func (r *errorImp) ErrorType() ErrorType {
	return r.errorType
}

func NewError(ErrorType ErrorType, message string) Error {
	if message == "" {
		message = string(ErrorType)
	}

	return &errorImp{
		errorType: ErrorType,
		message:   message,
	}
}

type ErrorType string

const (
	// 400 Bad Request
	ErrorInvalidArgument   ErrorType = "invalid argument"
	ErrorFaildPrecondition ErrorType = "failed precondition"
	ErrorOutOfRange        ErrorType = "out of range"
	// 401 Unauthorized
	ErrorUnauthenticated ErrorType = "unauthenticated"
	// 403 Forbidden
	ErrorPermissionDenied ErrorType = "permission denied"
	// 404 Not Found
	ErrorNotFound ErrorType = "not found"
	// 409 Conflict
	ErrorAlreadyExists ErrorType = "already exists"
	ErrorAborted       ErrorType = "aborted"
	// 429 Too Many Request
	ErrorResourceExhausted ErrorType = "resource exhausted"
	// 499 Client Closed Request
	ErrorCancelled ErrorType = "cancelled"

	// 500 Internal Server Error
	ErrorUnknown  ErrorType = "unknown Erroror"
	ErrorInternal ErrorType = "internal server Erroror"
	ErrorDataLoss ErrorType = "data loss"
	// 501 Not Implemented
	ErrorUnimplemented ErrorType = "unimplemented"
	// 503 Service Unavailable
	ErrorUnavailable ErrorType = "unavailable"
	// 504 Gateway Timeout
	ErrorDeadlineExceeded ErrorType = "deadline exceeded"
)
