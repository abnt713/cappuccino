package icons

import (
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
func (t Typicons) Battery(level cappuccino.BatteryLevel) *pango.Node {
	return t.icon("battery-" + string(level))
}

// BatteryCharging creates a battery charging icon.
func (t Typicons) BatteryCharging() *pango.Node {
	return t.smallIcon("plug")
}

// Calendar is a calendar icon.
func (t Typicons) Calendar() *pango.Node {
	return t.smallIcon("calendar")
}

// Clock is a clock icon.
func (t Typicons) Clock() *pango.Node {
	return t.icon("time")
}

// Stopwatch is a stopwatch icon.
func (t Typicons) Stopwatch() *pango.Node {
	return t.icon("stopwatch")
}

// Lock is a padlock icon.
func (t Typicons) Lock(opened bool) *pango.Node {
	if opened {
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
