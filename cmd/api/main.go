package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/s0rg/genpassaas/pkg/config"
	"github.com/s0rg/genpassaas/pkg/gen"
)

const (
	appName = config.Name + "-api"
	envKey  = "KEY"
	envAddr = "ADDR"
)

func loadEnv() (key, addr string) {
	if key = os.Getenv(envKey); key == "" {
		log.Fatal("No api KEY is set")
	}

	if addr = os.Getenv(envAddr); addr == "" {
		addr = "localhost:8080"
	}

	return key, addr
}

func main() {
	key, addr := loadEnv()

	m := http.NewServeMux()

	m.HandleFunc("/v1/smart", withBearer(key, gen.Smart))
	m.HandleFunc("/v1/simple", withBearer(key, gen.Simple))

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := []string{
			appName,
			config.Version(),
			"api: GET /v1/simple?num={N}&len={L} or /v1/smart?num={N}&len={L}",
		}

		if err := RenderStrings(w, r, result); err != nil {
			log.Printf("client[%s] error: %v", r.RemoteAddr, err)
		}
	})

	s := &http.Server{
		Addr:              addr,
		Handler:           m,
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       time.Second,
		WriteTimeout:      time.Second,
	}

	log.Printf("%s %s", appName, config.Version())
	log.Printf("Key: `%s`", key)
	log.Printf("Serving on: %s", addr)

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
