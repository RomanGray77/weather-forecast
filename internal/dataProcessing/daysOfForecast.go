package dataprocessing

import "time"

// StartOfWeekDate returns the first date of the current week
func StartOfWeekDate() time.Time {
	// Go's Weekday: Sunday=0, Monday=1, ..., Saturday=6
	offset := (int(time.Now().Weekday()) - int(time.Monday) + 7) % 7
	return time.Now().AddDate(0, 0, -offset)
}

func TomorrowDate() time.Time {
	return time.Now().AddDate(0, 0, 1)
}
