package error

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Really anything that doesn't turn to Internal Server Error.
const (
	InvalidCredentialsError AppErrorKind = iota
	EmailTakenError
	TfaFailedError
)

type AppErrorKind int

type AppError interface {
	Message() string
	Kind() AppErrorKind
}

func WriteErrorResponse(w http.ResponseWriter, e AppError) {
	message, kind := e.Message(), e.Kind()

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(func() int {
		switch kind {
		case InvalidCredentialsError:
			return http.StatusUnauthorized
		default:
			return http.StatusBadRequest
		}
	}())

	jsonResponse, err := json.Marshal(map[string]any{
		"message": message,
		"kind":    kind,
		"success": false,
		"data":    nil,
	})

	if err != nil {
		http.Error(w, "Failed to marshal error response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
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
