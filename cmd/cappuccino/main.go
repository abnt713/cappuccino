package main

import (
	"time"

	"github.com/abnt713/cappuccino/cfg"
	"github.com/abnt713/cappuccino/pkg"
	"github.com/abnt713/cappuccino/pkg/cappuccino"
	"github.com/abnt713/cappuccino/pkg/icons"
	"github.com/abnt713/cappuccino/pkg/log"
)

func main() {
	conf := cfg.GetConfig()
	icons := icons.NewTypicons(conf.TypiconsPath)

	countdowns := make(pkg.Modules, 0, len(conf.Events))
	for _, evt := range conf.Events {
		evtTime, err := time.Parse(time.RFC3339, evt.Date)
		if err != nil {
			panic(err)
		}
		countdowns = append(
			countdowns,
			cappuccino.NewCountdown(evt.Name, evtTime, evt.Rate, icons),
		)
	}
	logger := log.NewSTDOut()

	modules := pkg.Modules{
		cappuccino.NewNetworkManagerViewer(icons),
		cappuccino.NewPulseAudioViewer(logger, icons),
		cappuccino.NewBatteryViewer("", icons),
		cappuccino.NewClock(icons),
	}

	modules = append(countdowns, modules...)

	app := cappuccino.NewApp(modules)
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
