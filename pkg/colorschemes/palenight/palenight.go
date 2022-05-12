package palenight

import (
	"image/color"
	"time"

	"github.com/abnt713/cappuccino/pkg/cappuccino"
)

// New creates a new palenight instance.
func New() Palenight {
	return Palenight{
		colors: getColors(),
	}
}

// Palenight is the palenight colorscheme.
type Palenight struct {
	evtRemainingWarning time.Duration

	colors palenightColors
}

// VPN informs all vpn related colors.
func (p Palenight) VPN(on bool) color.Color {
	if on {
		return p.colors.cyan
	}

	return p.colors.blue
}

// Clock is the color of a clock.
func (p Palenight) Clock(now time.Time) color.Color {
	currHour := now.Hour()
	if currHour < 12 {
		return p.colors.cyan
	}

	if currHour < 18 {
		return p.colors.darkYellow
	}

	return p.colors.purple
}

// Calendar is the color of a calendar.
func (p Palenight) Calendar(now time.Time) color.Color {
	weekday := now.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		return p.colors.matGreenLight
	}

	return p.colors.green
}

// Battery contains all battery colors.
func (p Palenight) Battery(level cappuccino.BatteryLevel, isCharging bool) color.Color {
	if isCharging {
		return p.colors.cyan
	}

	switch level {
	case cappuccino.BatteryLevelLow:
		return p.colors.red
	case cappuccino.BatteryLevelFull:
		return p.colors.cyan
	default:
		return p.colors.yellow
	}
}

// Sound contains the sound colors.
func (p Palenight) Sound(isMuted bool, _ float32) color.Color {
	if isMuted {
		return p.colors.commentGrey
	}

	return p.colors.matWhiteLight
}

// Stopwatch contains all stopwatch related colors.
func (p Palenight) Stopwatch(
	evt cappuccino.CountdownEvent,
	remaining time.Duration,
) color.Color {
	happened := remaining < 0
	isClose := evt.IsClose(remaining)
	if happened {
		return p.colors.darkRed
	}

	if isClose {
		return p.colors.red
	}

	return p.colors.lightRed
}

// Github is the color of the github viewer.
func (p Palenight) Github() color.Color {
	return p.colors.cyan
}
