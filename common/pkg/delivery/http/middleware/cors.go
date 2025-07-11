package middleware

import (
	"net/http"
	"strconv"
	"strings"
)

type Cors struct {
	Origin      string
	Methods     []string
	Headers     []string
	Credentials bool
}

func DefaultCors() Cors {
	return Cors{
		Origin:      "*",
		Methods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		Headers:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		Credentials: false,
	}
}

func (cors Cors) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", cors.Origin)
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(cors.Methods, ", "))
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(cors.Headers, ", "))
			w.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(cors.Credentials))

			// Handle preflight OPTIONS requests
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
