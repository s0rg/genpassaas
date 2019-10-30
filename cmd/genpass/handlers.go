package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/s0rg/genpassaas/pkg/handlers"
)

const (
	maxCount  = 100
	maxLength = 64
)

type genFunc func(length int) string

func getInt(s string) (val int, ok bool) {
	if s == "" {
		return
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return
	}

	return i, true
}

func getIntOrDefault(s string, d int) int {
	i, ok := getInt(s)
	if !ok {
		return d
	}

	return i
}

func ensureRange(v, min, max int) int {
	switch {
	case v < min:
		return min
	case v > max:
		return max
	}

	return v
}

func genHandler(gen genFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			args   = r.URL.Query()
			count  = getIntOrDefault(args.Get("num"), 10)
			length = getIntOrDefault(args.Get("len"), 8)
		)

		count = ensureRange(count, 1, maxCount)
		length = ensureRange(length, 6, maxLength)

		result := make([]string, count)
		for i := 0; i < count; i++ {
			result[i] = gen(length)
		}

		if err := handlers.RenderStrings(w, r, result); err != nil {
			log.Printf("client[%s] error: %v", r.RemoteAddr, err)
		}
	}
}

func getWithBearer(key string, gen genFunc) http.HandlerFunc {
	return handlers.GET(handlers.BearerAuth(key, genHandler(gen)))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	result := []string{appName, GitHash}

	if err := handlers.RenderStrings(w, r, result); err != nil {
		log.Printf("client[%s] error: %v", r.RemoteAddr, err)
	}
}
