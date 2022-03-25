package main

import (
	"github.com/abnt713/cappuccino/pkg"
	"github.com/abnt713/cappuccino/pkg/cappuccino"
)

func main() {
	modules := pkg.Modules{
		cappuccino.NewNetworkManagerViewer(),
		cappuccino.NewPulseAudioViewer(),
		cappuccino.NewBatteryViewer(""),
		cappuccino.NewDateViewer(),
		cappuccino.NewClock(nil),
	}
	app := cappuccino.NewApp(modules)
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
