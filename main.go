package main

import (
	"log"
	"net/http"

	"github.com/RomanGray77/weather-forecast/internal/router"
)

func main() {
	r := router.New()

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
