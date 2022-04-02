package cappuccino

import (
	"image/color"
	"time"

	"barista.run/bar"
	"barista.run/modules/clock"
	"barista.run/outputs"
	"barista.run/pango"
)

// NewClock creates a new clock.
func NewClock(icons ClockIcons, colors ClockColors) Clock {
	return Clock{
		icons:  icons,
		colors: colors,
	}
}

// Clock is the statusbar clock.
type Clock struct {
	icons  ClockIcons
	colors ClockColors
}

// GenerateBaristaModule generates the clock barista module.
func (cl Clock) GenerateBaristaModule() (bar.Module, error) {
	localtime := clock.Local().Output(
		time.Second,
		func(now time.Time) bar.Output {
			date := now.Format("Mon Jan 2")
			time := now.Format("15:04:05")

			calendarNode := pango.New(
				cl.icons.Calendar(now),
				space,
				pango.Text(date),
			).Color(cl.colors.Calendar(now))

			clockNode := pango.New(
				cl.icons.Clock(now),
				space,
				pango.Text(time),
			).Color(cl.colors.Clock(now))

			return outputs.Pango(
				calendarNode,
				space,
				clockNode,
			)
		},
	)

	return localtime, nil
}

// ClockIcons contains all icons related to time.
type ClockIcons interface {
	Calendar(time.Time) *pango.Node
	Clock(time.Time) *pango.Node
}

// ClockColors defines the clock colorscheme.
type ClockColors interface {
	Calendar(time.Time) color.Color
	Clock(time.Time) color.Color
}
