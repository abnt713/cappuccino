package cfg

import (
	"github.com/abnt713/cappuccino/pkg/cappuccino"
)

// Config has all settings for the bar.
type Config struct {
	LogFilePath  string
	TypiconsPath string
	Events       []cappuccino.CountdownEvent
}
