package cappuccino

import (
	"time"

	"barista.run/bar"
	"barista.run/base/click"
	"barista.run/modules/clock"
	"barista.run/outputs"
)

// NewClock creates a new clock.
func NewClock(toggler ModeToggler) Clock {
	return Clock{
		toggler: toggler,
	}
}

// Clock is the statusbar clock.
type Clock struct {
	toggler ModeToggler
}

// GenerateBaristaModule generates the clock barista module.
func (cl Clock) GenerateBaristaModule() (bar.Module, error) {
	localtime := clock.Local().Output(
		time.Second,
		func(now time.Time) bar.Output {
			return outputs.Text(now.Format("15:04:05")).
				OnClick(click.Left(func() {
					if cl.toggler != nil {
						cl.toggler.ToggleMode("date")
					}
				}))
		},
	)

	return localtime, nil
}

// ModeToggler toggles a mode.
type ModeToggler interface {
	ToggleMode(name string)
}
