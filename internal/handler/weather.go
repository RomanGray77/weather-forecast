// Package handler implements the HTTP handlers for the weather page.
package handler

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/RomanGray77/weather-forecast/internal/storage"
	"github.com/RomanGray77/weather-forecast/internal/weather"
)

//go:embed templates/*.html
var templateFS embed.FS

var weatherTemplate = template.Must(template.ParseFS(templateFS, "templates/weather.html"))

type pageData struct {
	City    string
	Results []weather.DayForecast
	Error   string
}

// WeatherPage renders the empty submission form.
func WeatherPage(w http.ResponseWriter, r *http.Request) {
	render(w, pageData{})
}

// WeatherSubmit handles the form submission: it fetches the forecast for
// the submitted city, persists it, and renders the page with the results.
func WeatherSubmit(w http.ResponseWriter, r *http.Request) {
	city := r.FormValue("city")
	if city == "" {
		render(w, pageData{Error: "Please enter a city."})
		return
	}

	days, err := weather.Fetch(city)
	if err != nil {
		render(w, pageData{City: city, Error: "Could not fetch forecast: " + err.Error()})
		return
	}

	if err := storage.Save(days); err != nil {
		render(w, pageData{City: city, Error: "Could not save forecast: " + err.Error()})
		return
	}

	render(w, pageData{City: city, Results: days})
}

func render(w http.ResponseWriter, data pageData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := weatherTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
