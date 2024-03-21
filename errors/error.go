package errors

import "fmt"

type AppErrorCode int

const (
	AppErrorUnknown       AppErrorCode = iota
	AppErrorInvalidParams AppErrorCode = iota
	AppInternalError      AppErrorCode = iota
	AppConfigError        AppErrorCode = iota
	AppDatabaseError      AppErrorCode = iota
)

type AppError struct {
	Message   string       `json:"message"`
	ErrorCode AppErrorCode `json:"code"`
}

func NewAppErrorf(code AppErrorCode, format string, a ...interface{}) *AppError {
	return &AppError{
		Message:   fmt.Sprintf(format, a...),
		ErrorCode: code,
	}
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.ErrorCode, e.Message)
}
