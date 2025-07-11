package error

import (
	"fmt"
)

// Really anything that doesn't turn to Internal Server Error.
const (
	InvalidCredentialsError AppErrorKind = iota + 1
	EmailTakenError
	MfaFailedError
	BadRequestError
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
