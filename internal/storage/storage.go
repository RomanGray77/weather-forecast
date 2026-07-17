// Package storage persists forecast data to disk.
package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/RomanGray77/weather-forecast/internal/config"
)

const filename = "weather_Wetzikon.json"

// SaveForecastToJSON writes days to weather_Wetzikon.json as indented JSON, overwriting any previous contents.
func Save(days []config.DayForecast) error {
	data, err := json.MarshalIndent(days, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling forecast: %w", err)
	}

	if err := os.WriteFile(filename, data, 0o644); err != nil {
		return fmt.Errorf("writing %s: %w", filename, err)
	}

	return nil
}

// LoadForecastFromJSON reads days from weather_Wetzikon.json.
func Load() ([]config.DayForecast, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", filename, err)
	}

	var days []config.DayForecast
	if err := json.Unmarshal(data, &days); err != nil {
		return nil, fmt.Errorf("unmarshaling forecast: %w", err)
	}

	return days, nil
}
