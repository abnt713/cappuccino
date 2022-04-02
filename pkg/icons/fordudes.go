package icons

import (
	"time"

	"barista.run/pango"
)

// NewTypicons4Dudes returns a new Typicons4Dudes instance.
func NewTypicons4Dudes(icons Typicons) Typicons4Dudes {
	return Typicons4Dudes{
		Typicons: icons,
	}
}

// Typicons4Dudes checks for wednesdays.
type Typicons4Dudes struct {
	Typicons
}

// Calendar overrides the typicons calendar.
func (dudes Typicons4Dudes) Calendar(now time.Time) *pango.Node {
	if now.Weekday() == time.Wednesday {
		return pango.Text("🐸")
	}

	return dudes.Typicons.Calendar(now)
}
