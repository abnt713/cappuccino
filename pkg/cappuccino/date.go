package cappuccino

import (
	"time"

	"barista.run/bar"
	"barista.run/base/click"
	"barista.run/modules/clock"
	"barista.run/outputs"
	"barista.run/pango"
)

// DateViewer is the statusbar date.
type DateViewer struct{}

// NewDateViewer creates a new date viewer.
func NewDateViewer() DateViewer {
	return DateViewer{}
}

// GenerateBaristaModule generated the date viewer barista module.
func (cl DateViewer) GenerateBaristaModule() (bar.Module, error) {
	localdate := clock.Local().Output(
		time.Minute,
		func(now time.Time) bar.Output {
			return outputs.Pango(
				pango.Icon("material-today").Alpha(0.6),
				now.Format("Mon Jan 2"),
			).OnClick(click.RunLeft("gsimplecal"))
		},
	)

	return localdate, nil
}
