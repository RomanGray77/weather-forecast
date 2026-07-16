// Package router wires up the HTTP routes for the application.
package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/RomanGray77/weather-forecast/internal/handler"
)

// New returns the application's HTTP handler.
func New() http.Handler {
	r := chi.NewRouter()

	r.Get("/weather", handler.WeatherPage)
	r.Post("/weather", handler.WeatherSubmit)

	return r
}
