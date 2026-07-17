package main

import (
	"log"
	"net/http"

	"github.com/RomanGray77/weather-forecast/internal/router"
)

const portNumber = ":8081"

func main() {
	r := router.New()

	log.Println("Starting application on port ", portNumber)
	if err := http.ListenAndServe(portNumber, r); err != nil {
		log.Fatal(err)
	}
}
