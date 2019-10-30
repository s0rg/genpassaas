package handlers

import (
	"net/http"
	"strings"
)

func GET(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			code := http.StatusMethodNotAllowed
			http.Error(w, http.StatusText(code), code)
			return
		}

		next(w, r)
	}
}

func BearerAuth(key string, next http.HandlerFunc) http.HandlerFunc {
	const bearerName = "bearer"

	return func(w http.ResponseWriter, r *http.Request) {
		var (
			auth  string
			parts []string
		)

		if auth = r.Header.Get("Authorization"); auth == "" {
			http.Error(w, "Auth header missing", http.StatusUnauthorized)
			return
		}

		if parts = strings.Fields(auth); len(parts) != 2 {
			http.Error(w, "Auth header malformed", http.StatusUnauthorized)
			return
		}

		if !strings.EqualFold(parts[0], bearerName) {
			http.Error(w, "Auth header invalid type - not a bearer!", http.StatusUnauthorized)
			return
		}

		if parts[1] != key {
			http.Error(w, "Invalid bearer token", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}
