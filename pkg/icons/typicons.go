package icons

import (
	"time"

	"barista.run/pango"
	"barista.run/pango/icons/typicons"
	"github.com/abnt713/cappuccino/pkg/cappuccino"
)

// NewTypicons creates a new typicons instance.
func NewTypicons(fontsPath string) Typicons {
	typicons.Load(fontsPath)
	return Typicons{fontSize: 18}
}

// Typicons is a specific collection of icons which uses typicons.
type Typicons struct {
	fontSize float64
}

func (t Typicons) icon(icon string) *pango.Node {
	return pango.Icon("typecn-" + icon).Size(t.fontSize)
}

func (t Typicons) smallIcon(icon string) *pango.Node {
	return pango.Icon("typecn-" + icon).Size(t.fontSize - 2)
}

// Battery creates a battery icon.
func (t Typicons) Battery(level cappuccino.BatteryLevel, isCharging bool) *pango.Node {
	if isCharging {
		return t.smallIcon("battery-charge")
	}
	return t.icon("battery-" + string(level))
}

// Calendar is a calendar icon.
func (t Typicons) Calendar(_ time.Time) *pango.Node {
	return t.smallIcon("calendar")
}

// Clock is a clock icon.
func (t Typicons) Clock(_ time.Time) *pango.Node {
	return t.icon("time")
}

// Stopwatch is a stopwatch icon.
func (t Typicons) Stopwatch() *pango.Node {
	return t.icon("stopwatch")
}

// VPN is the vpn icon.
func (t Typicons) VPN(on bool) *pango.Node {
	if !on {
		return t.icon("lock-open-outline")
	}

	return t.icon("lock-closed-outline")
}

// Sound is a sound icon.
func (t Typicons) Sound(muted bool, intensity float32) *pango.Node {
	if muted {
		return t.icon("volume-mute")
	}

	if intensity < 0.5 {
		return t.icon("volume-down")
	}

	return t.icon("volume-up")
}
