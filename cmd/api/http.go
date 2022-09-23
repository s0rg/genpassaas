package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/s0rg/genpassaas/pkg/config"
	"github.com/s0rg/genpassaas/pkg/gen"
)

const (
	headerAuth = "Authorization"
	bearerName = "Bearer"
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

func Bearer(key string, next http.HandlerFunc) http.HandlerFunc {
	const fieldsCount = 2

	return func(w http.ResponseWriter, r *http.Request) {
		var (
			auth  string
			parts []string
		)

		if auth = r.Header.Get(headerAuth); auth == "" {
			http.Error(w, "Auth header missing", http.StatusUnauthorized)

			return
		}

		if parts = strings.Fields(auth); len(parts) != fieldsCount {
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

func withBearer(key string, fn gen.Fn) http.HandlerFunc {
	return GET(Bearer(key, genHandler(fn)))
}

func genHandler(fn gen.Fn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		args := r.URL.Query()

		count := config.ClampCount(getIntParam(
			args.Get("num"),
			config.DefaultCount,
		))

		length := config.ClampLength(getIntParam(
			args.Get("len"),
			config.DefaultLength,
		))

		result := make([]string, count)

		for i := 0; i < count; i++ {
			result[i] = fn(length)
		}

		if err := RenderStrings(w, r, result); err != nil {
			log.Printf("client[%s] error: %v", r.RemoteAddr, err)
		}
	}
}

func getInt(s string) (val int, ok bool) {
	if s == "" {
		return
	}

	if i, err := strconv.Atoi(s); err == nil {
		return i, true
	}

	return 0, false
}

func getIntParam(s string, d int) (rv int) {
	if i, ok := getInt(s); ok {
		return i
	}

	return d
}
