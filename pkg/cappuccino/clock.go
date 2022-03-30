package cappuccino

import (
	"time"

	"barista.run/bar"
	"barista.run/modules/clock"
	"barista.run/outputs"
	"barista.run/pango"
)

// NewClock creates a new clock.
func NewClock(icons ClockIcons) Clock {
	return Clock{
		icons: icons,
	}
}

// Clock is the statusbar clock.
type Clock struct {
	icons ClockIcons
}

// GenerateBaristaModule generates the clock barista module.
func (cl Clock) GenerateBaristaModule() (bar.Module, error) {
	localtime := clock.Local().Output(
		time.Second,
		func(now time.Time) bar.Output {
			date := now.Format("Mon Jan 2")
			time := now.Format("15:04:05")
			return outputs.Pango(
				cl.icons.Calendar(),
				space,
				pango.Text(date),
				space,
				cl.icons.Clock(),
				space,
				pango.Text(time),
			)
		},
	)

	return localtime, nil
}

// ClockIcons contains all icons related to time.
type ClockIcons interface {
	Calendar() *pango.Node
	Clock() *pango.Node
}
