package util

import "google.golang.org/grpc/codes"

type ApiError struct {
	HTTPErrorCode int
	GRPCErrorCode codes.Code
	ErrorMessage  string
	Err           error
}

func (ae ApiError) Error() string {
	return ae.ErrorMessage
}

func (ae ApiError) Unwrap() error {
	return ae.Err
}

func ErrAlreadyExists(err error, message string) ApiError {
	return ApiError{
		GRPCErrorCode: codes.AlreadyExists,
		HTTPErrorCode: 403,
		ErrorMessage:  message,
		Err:           err,
	}
}
