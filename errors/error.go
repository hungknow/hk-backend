package errors

import (
	"fmt"
	"io"
)

type AppErrorCode int

const (
	AppErrorUnknown       AppErrorCode = iota
	AppErrorInvalidParams AppErrorCode = iota
	AppInternalError      AppErrorCode = iota
	AppConfigError        AppErrorCode = iota
	AppDatabaseError      AppErrorCode = iota
	AppLoadDataError      AppErrorCode = iota
)

type AppError struct {
	Message   string       `json:"message"`
	ErrorCode AppErrorCode `json:"code"`
	stack     *stack
}

func NewAppErrorf(code AppErrorCode, format string, a ...interface{}) *AppError {
	return &AppError{
		Message:   fmt.Sprintf(format, a...),
		ErrorCode: code,
		stack:     callers(),
	}
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.ErrorCode, e.Message)
}

func (e *AppError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, e.Message)
			e.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.Message)
	case 'q':
		fmt.Fprintf(s, "%q", e.Message)
	}
}
