package main

import (
	"log"
	"net/http"

	"github.com/DevSecAI/gridcore-telemetry/internal/api"
	"github.com/DevSecAI/gridcore-telemetry/internal/config"
)

func main() {
	cfg := config.Load()
	r := api.NewRouter(cfg)
	log.Printf("gridcore listening on %s", cfg.Listen)
	if err := http.ListenAndServe(cfg.Listen, r); err != nil {
		log.Fatal(err)
	}
}
