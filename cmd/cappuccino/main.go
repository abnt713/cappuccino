package main

import (
	"fmt"

	"github.com/Wifx/gonetworkmanager"
	"github.com/abnt713/cappuccino/cfg"
	"github.com/abnt713/cappuccino/pkg"
	"github.com/abnt713/cappuccino/pkg/cappuccino"
	"github.com/abnt713/cappuccino/pkg/colorschemes/palenight"
	"github.com/abnt713/cappuccino/pkg/icons"
	"github.com/abnt713/cappuccino/pkg/log"
)

func main() {
	conf, err := cfg.ReadJSONConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", conf)
	nm, err := gonetworkmanager.NewNetworkManager()
	if err != nil {
		panic(err)
	}

	logger, err := log.NewFile(conf.LogFilePath)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	icons := icons.NewTypicons4Dudes(icons.NewTypicons(conf.TypiconsPath))
	colorscheme := palenight.New()

	countdowns := make(pkg.Modules, 0, len(conf.Events))
	for _, evt := range conf.Events {
		countdowns = append(
			countdowns,
			cappuccino.NewCountdown(evt, icons, colorscheme),
		)
	}

	modules := pkg.Modules{
		cappuccino.NewVPNViewer(nm, icons, colorscheme),
		cappuccino.NewPulseAudioViewer(logger, icons, colorscheme),
		cappuccino.NewBatteryViewer("", conf.BatteryLevels, icons, colorscheme),
		cappuccino.NewClock(icons, colorscheme),
	}

	modules = append(countdowns, modules...)

	app := cappuccino.NewApp(modules)
	err = app.Run()
	if err != nil {
		panic(err)
	}
}
