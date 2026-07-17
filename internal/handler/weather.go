package handler

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/RomanGray77/weather-forecast/internal/config"
	"github.com/RomanGray77/weather-forecast/internal/storage"
)

//go:embed templates/weather.html
var templateFS embed.FS
var weatherTemplate = template.Must(template.ParseFS(templateFS, "templates/weather.html"))

type pageData struct {
	City    string
	Results []config.DayForecast
	Error   string
}

// WeatherPage renders the empty submission form.
func WeatherPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, pageData{})
}

// WeatherSubmit handles the form submission: it fetches the forecast for
// the submitted city, persists it, and renders the page with the results.
func WeatherSubmit(w http.ResponseWriter, r *http.Request) {
	city := r.FormValue("city")
	if city == "" {
		renderTemplate(w, pageData{Error: "Please enter a city."})
		return
	}

	savedForecasts, err := storage.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("The latest saved forecast day:", savedForecasts[len(savedForecasts)-1].Date)

	renderTemplate(w, pageData{City: city, Results: savedForecasts})
}

func renderTemplate(w http.ResponseWriter, data pageData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := weatherTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
