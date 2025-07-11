package error

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	InternalError ErrorKind = iota
	InvalidCredentialsError
	EmailTakenError
)

type ErrorKind int

type Error interface {
	Message() string
	Kind() ErrorKind
}

func WriteErrorResponse(w http.ResponseWriter, e Error) {
	message, kind := e.Message(), e.Kind()

	if kind == InternalError {
		message = "An internal error occurred. Please try again later."
		log.Printf("Internal error occurred: %s", message)
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(func() int {
		switch kind {
		case InternalError:
			return http.StatusInternalServerError
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
	kind    ErrorKind
}

func Errorf(kind ErrorKind, format string, args ...any) errorWrap {
	return errorWrap{
		message: fmt.Errorf(format, args...).Error(),
		kind:    kind,
	}
}

func (e errorWrap) Message() string {
	return e.message
}

func (e errorWrap) Kind() ErrorKind {
	return e.kind
}
