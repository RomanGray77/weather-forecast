// Package storage persists forecast data to disk.
package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/RomanGray77/weather-forecast/internal/weather"
)

const filename = "weather_forecast.json"

// Save writes days to weather_forecast.json as indented JSON, overwriting
// any previous contents.
func Save(days []weather.DayForecast) error {
	data, err := json.MarshalIndent(days, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling forecast: %w", err)
	}

	if err := os.WriteFile(filename, data, 0o644); err != nil {
		return fmt.Errorf("writing %s: %w", filename, err)
	}

	return nil
}
