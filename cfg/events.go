package cfg

import (
	"time"

	"github.com/abnt713/cappuccino/pkg/cappuccino"
)

func newEvent(name string, date string, IsDeadline bool, urgentWithLessThan time.Duration) cappuccino.CountdownEvent {
	evtTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		panic(err)
	}
	return cappuccino.CountdownEvent{
		Name:               name,
		Date:               evtTime,
		IsDeadline:         IsDeadline,
		UrgentWithLessThan: urgentWithLessThan,
	}
}

func newDeadline(name string, date string, urgentWithLessThan time.Duration) cappuccino.CountdownEvent {
	return newEvent(name, date, true, urgentWithLessThan)
}

func newInterest(name string, date string, urgentWithLessThan time.Duration) cappuccino.CountdownEvent {
	return newEvent(name, date, false, urgentWithLessThan)
}
