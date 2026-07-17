// Package weather fetches forecast data from wttr.in.
package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/RomanGray77/weather-forecast/internal/config"
)

// The data structure returned by wttr.in's JSON API.
type wttrResponse struct {
	Weather []struct {
		Date      string `json:"date"`
		AvgTempC  string `json:"avgtempC"`
		Astronomy []struct {
			Sunrise string `json:"sunrise"`
			Sunset  string `json:"sunset"`
		} `json:"astronomy"`
	} `json:"weather"`
}

// Fetch requests the forecast for city from wttr.in and returns one
// DayForecast per day in the response.
func Fetch(city string) ([]config.DayForecast, error) {
	reqURL := fmt.Sprintf("https://wttr.in/%s?format=j2", url.PathEscape(city))

	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, fmt.Errorf("requesting wttr.in: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wttr.in returned status %d", resp.StatusCode)
	}

	var parsed wttrResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, fmt.Errorf("decoding wttr.in response: %w", err)
	}

	days := make([]config.DayForecast, 0, len(parsed.Weather))
	for _, w := range parsed.Weather {
		day := config.DayForecast{Date: w.Date, AvgTempC: w.AvgTempC}
		if len(w.Astronomy) > 0 {
			day.Sunrise = w.Astronomy[0].Sunrise
			day.Sunset = w.Astronomy[0].Sunset
		}
		days = append(days, day)
	}

	return days, nil
}
