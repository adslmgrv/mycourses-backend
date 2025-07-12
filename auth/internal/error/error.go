package error

import (
	"fmt"
)

const (
	// 0 is for internal server error and 1 is for gin's default errors.
	InvalidCredentialsError AppErrorKind = iota + 2
	EmailTakenError
	MfaFailedError
)

type AppErrorKind int

type AppError interface {
	Message() string
	Kind() AppErrorKind
}

type errorWrap struct {
	message string
	kind    AppErrorKind
}

func Errorf(kind AppErrorKind, format string, args ...any) errorWrap {
	return errorWrap{
		message: fmt.Errorf(format, args...).Error(),
		kind:    kind,
	}
}

func (e errorWrap) Error() string {
	return e.message
}

func (e errorWrap) Message() string {
	return e.message
}

func (e errorWrap) Kind() AppErrorKind {
	return e.kind
}
