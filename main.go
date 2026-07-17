package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RomanGray77/weather-forecast/internal/config"
	dataprocessing "github.com/RomanGray77/weather-forecast/internal/dataProcessing"
	"github.com/RomanGray77/weather-forecast/internal/router"
	"github.com/RomanGray77/weather-forecast/internal/storage"
	"github.com/RomanGray77/weather-forecast/internal/weather"
)

const portNumber = ":8081"

func main() {
	// Load the previously saved forecast from disk.``
	savedForecasts, err := storage.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("The latest saved forecast day:", savedForecasts[len(savedForecasts)-1].Date)

	// Fetch the latest forecast from wttr.in for Wetzikon.
	newForecasts, err := weather.Fetch("Wetzikon")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("The latest forecast day from wttr.in:", newForecasts[len(newForecasts)-1].Date)

	startDate := dataprocessing.StartOfWeekDate()
	endDate := dataprocessing.TomorrowDate()
	// Compare the latest saved forecast with the fetched forecast and add any new days.

	var requiredForecasts []config.DayForecast
	// Collect from savedForecasts the required date range till yesterday
	func() {
		for _, day := range savedForecasts {
			if day.Date >= startDate.Format("2006-01-02") && day.Date < time.Now().Format("2006-01-02") {
				requiredForecasts = append(requiredForecasts, day)
			}
		}
	}()
	// Collect from newForecasts the required date range till tomorrow
	func() {
		for _, day := range newForecasts {
			if day.Date >= time.Now().Format("2006-01-02") && day.Date <= endDate.Format("2006-01-02") {
				requiredForecasts = append(requiredForecasts, day)
			}
		}
	}()

	fmt.Println("\nThe forecast for the required date range:")
	for _, day := range requiredForecasts {
		fmt.Printf("Date %s: AvgTempC=%s, Sunrise=%s, Sunset=%s\n", day.Date, day.AvgTempC, day.Sunrise, day.Sunset)
	}

	if err := storage.Save(requiredForecasts); err != nil {
		log.Fatal(err)
	}

	r := router.New()

	log.Println("Starting application on port ", portNumber)

	if err := http.ListenAndServe(portNumber, r); err != nil {
		log.Fatal(err)
	}

}
