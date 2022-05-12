package cfg

import (
	"github.com/abnt713/cappuccino/pkg/cappuccino"
)

// Config has all settings for the bar.
type Config struct {
	LogFilePath   string                      `json:"log_file_path"`
	TypiconsPath  string                      `json:"typicons_path"`
	Events        []cappuccino.CountdownEvent `json:"events"`
	BatteryLevels cappuccino.BatteryIntervals `json:"battery_levels"`
	Github        struct {
		Token      string `json:"token"`
		IgnoreRead bool   `json:"ignore_read"`
	} `json:"github"`
}
