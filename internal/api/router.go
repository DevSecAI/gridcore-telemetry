package api

import (
	"net/http"

	"github.com/DevSecAI/gridcore-telemetry/internal/config"
	"github.com/DevSecAI/gridcore-telemetry/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter(_ config.Config) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(`{"ok":true}`))
	})
	r.Post("/meters/reading", MeterReading)
	return r
}
