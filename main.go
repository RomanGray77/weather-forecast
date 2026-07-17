package main

import (
	"fmt"
	"log"

	"github.com/RomanGray77/weather-forecast/internal/storage"
	"github.com/RomanGray77/weather-forecast/internal/weather"
)

const portNumber = ":8081"

func main() {
	savedDays, err := storage.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("The latest saved forecast day:", savedDays[len(savedDays)-1].Date)

	fetchedDays, err := weather.Fetch("Wetzikon")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("The latest forecast day from wttr.in:", fetchedDays[len(fetchedDays)-1].Date)

	fmt.Println("\nThe date that will be added to the forecast:")
	for _, day := range fetchedDays {
		if savedDays[len(savedDays)-1].Date < day.Date {
			savedDays = append(savedDays, day)
			fmt.Printf("Added date %s: AvgTempC=%s, Sunrise=%s, Sunset=%s\n", day.Date, day.AvgTempC, day.Sunrise, day.Sunset)
		}
	}

	if err := storage.Save(savedDays); err != nil {
		log.Fatal(err)
	}

	/*
	   r := router.New()

	   log.Println("Starting application on port ", portNumber)

	   	if err := http.ListenAndServe(portNumber, r); err != nil {
	   		log.Fatal(err)
	   	}
	*/
}
