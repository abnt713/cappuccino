package cfg

import "time"

// Config has all settings for the bar.
type Config struct {
	TypiconsPath string
	Events       []Event
}

// Event is an event worth remembering.
type Event struct {
	Name string
	Date string
	Rate time.Duration
}
