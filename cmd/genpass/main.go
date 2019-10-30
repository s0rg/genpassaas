package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/s0rg/genpassaas/pkg/gen"
	"github.com/s0rg/genpassaas/pkg/handlers"
)

const (
	appName = "GenPassaaS"
	envKey  = "KEY"
	envAddr = "ADDR"
)

var (
	GitHash   string
	BuildDate string
)

func main() {
	var (
		appKey, appAddr string
	)

	if appKey = os.Getenv(envKey); appKey == "" {
		log.Fatal("No api KEY is set")
	}

	if appAddr = os.Getenv(envAddr); appAddr == "" {
		appAddr = "localhost:8080"
	}

	m := http.NewServeMux()
	m.HandleFunc("/", handlers.GET(indexHandler))
	m.HandleFunc("/simple", getWithBearer(appKey, gen.Simple))
	m.HandleFunc("/smart", getWithBearer(appKey, gen.Smart))

	s := &http.Server{
		Addr:              appAddr,
		Handler:           m,
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       time.Second,
		WriteTimeout:      time.Second,
	}

	log.Printf("%s %s [build at %s]", appName, GitHash, BuildDate)
	log.Printf("Key: `%s`", appKey)
	log.Printf("Serving on: %s", appAddr)

	log.Fatal(s.ListenAndServe())
}
