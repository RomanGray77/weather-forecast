package config

// DayForecast is the subset of wttr.in's per-day data this app cares about.
type DayForecast struct {
	Date     string `json:"date"`
	AvgTempC string `json:"avgtempC"`
	Sunrise  string `json:"sunrise"`
	Sunset   string `json:"sunset"`
}
