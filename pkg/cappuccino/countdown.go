package cappuccino

import (
	"fmt"
	"strings"
	"time"

	"barista.run/bar"
	"barista.run/modules/funcs"
	"barista.run/outputs"
	"barista.run/pango"
)

// NewCountdown creates a new countdown instance.
func NewCountdown(
	name string,
	date time.Time,
	refresh time.Duration,
	icons CountdownIcons,
) Countdown {
	return Countdown{
		name:     name,
		date:     date,
		interval: refresh,
		icons:    icons,
	}
}

// Countdown counts down to an specific date.
type Countdown struct {
	name     string
	date     time.Time
	interval time.Duration
	icons    CountdownIcons
}

// HandleStream handles the called stream.
func (c Countdown) HandleStream(s bar.Sink) {
	now := time.Now()
	diff := c.date.Sub(now)
	if diff < 0 {
		s.Output(outputs.Text(fmt.Sprintf("!! %s !!", strings.ToUpper(c.name))))
		return
	}

	cd := fmt.Sprintf("%s até %s", fmtDuration(diff), c.name)
	s.Output(outputs.Pango(c.icons.Stopwatch(), space, pango.Text(cd)))
}

func fmtDuration(until time.Duration) string {
	shouldDisplaySeconds := (until < 30*time.Minute)

	dur := until.Round(time.Second)

	dateFmt := []string{}
	contents := []any{}
	day := dur / (24 * time.Hour)
	if day > 0 {
		dateFmt = append(dateFmt, "%02dd")
		contents = append(contents, day)
		dur -= day * (24 * time.Hour)
	}

	hrs := dur / time.Hour
	if hrs > 0 {
		dateFmt = append(dateFmt, "%02dh")
		contents = append(contents, hrs)
		dur -= hrs * time.Hour
	}

	min := dur / time.Minute
	if min > 0 {
		dateFmt = append(dateFmt, "%02dm")
		contents = append(contents, min)
		dur -= min * time.Minute
	}

	if shouldDisplaySeconds {
		sec := dur / time.Second
		dateFmt = append(dateFmt, "%02ds")
		contents = append(contents, sec)
	}

	return fmt.Sprintf(strings.Join(dateFmt, " "), contents...)
}

// GenerateBaristaModule generates the countdown module.
func (c Countdown) GenerateBaristaModule() (bar.Module, error) {
	return funcs.Every(c.interval, c.HandleStream), nil
}

// CountdownIcons contains all countdown related icons.
type CountdownIcons interface {
	Stopwatch() *pango.Node
}
