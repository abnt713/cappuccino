package cappuccino

import (
	"fmt"
	"image/color"
	"strings"
	"time"

	"barista.run/bar"
	"barista.run/modules/funcs"
	"barista.run/outputs"
	"barista.run/pango"
)

// NewCountdown creates a new countdown instance.
func NewCountdown(
	event CountdownEvent,
	icons CountdownIcons,
	colors CountdownColors,
) Countdown {
	return Countdown{
		event:  event,
		icons:  icons,
		colors: colors,
	}
}

// Countdown counts down to an specific date.
type Countdown struct {
	event    CountdownEvent
	interval time.Duration
	icons    CountdownIcons
	colors   CountdownColors
}

// HandleStream handles the called stream.
func (c Countdown) HandleStream(s bar.Sink) {
	now := time.Now()
	diff := c.event.Date.Sub(now)
	if diff < 0 {
		out := outputs.Text(
			fmt.Sprintf("!! %s !!", strings.ToUpper(c.event.Name)),
		).Color(c.colors.Stopwatch(c.event, diff))
		s.Output(out)
		return
	}

	cd := fmt.Sprintf("%s até %s", fmtDuration(diff), c.event.Name)
	out := outputs.Pango(c.icons.Stopwatch(), space, pango.Text(cd))
	s.Output(out.Color(c.colors.Stopwatch(c.event, diff)))
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
	return funcs.Every(time.Second, c.HandleStream), nil
}

// CountdownEvent is the event for which the clock counts towards.
type CountdownEvent struct {
	Name               string            `json:"name"`
	Date               time.Time         `json:"date"`
	IsDeadline         bool              `json:"is_deadline"`
	UrgentWithLessThan CountdownDuration `json:"urgent_with_less_than"`
}

// CountdownDuration wraps a countdown for better decoding.
type CountdownDuration time.Duration

// UnmarshalJSON customizes JSON unmarshalling process.
func (cd *CountdownDuration) UnmarshalJSON(data []byte) error {
	content := strings.Trim(string(data), `"`)
	dur, err := time.ParseDuration(content)
	if err != nil {
		return err
	}

	*cd = CountdownDuration(dur)
	return nil
}

// IsClose tells if the event is close.
func (evt CountdownEvent) IsClose(remaining time.Duration) bool {
	return remaining < time.Duration(evt.UrgentWithLessThan)
}

// CountdownIcons contains all countdown related icons.
type CountdownIcons interface {
	Stopwatch() *pango.Node
}

// CountdownColors contains all countdown related colors.
type CountdownColors interface {
	Stopwatch(evt CountdownEvent, remaining time.Duration) color.Color
}
