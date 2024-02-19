package errorx

import (
	"errors"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
)

type ErrorWrapper struct {
	Err  string
	Code int32
}

func (e *ErrorWrapper) Error() string {
	return e.Err
}

func New(code int32, err string) error {
	return &ErrorWrapper{
		Err:  err,
		Code: code,
	}
}

func GetHTTPCode(err error) int {
	var (
		errWrapper *ErrorWrapper
		statusCode = http.StatusInternalServerError
	)
	ok := errors.As(err, &errWrapper)
	if ok {
		statusCode = runtime.HTTPStatusFromCode(codes.Code(errWrapper.Code))
	}

	return statusCode
}

var (
	ErrorInternal = New(int32(codes.Internal), "Internal server error")
)
