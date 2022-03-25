package cappuccino

import (
	"barista.run/bar"
	"barista.run/modules/battery"
	"barista.run/outputs"
)

// NewBatteryViewer creates a new battery viewer instance.
func NewBatteryViewer(batteryName string) BatteryViewer {
	return BatteryViewer{
		batteryName: batteryName,
	}
}

// BatteryViewer displays battery information
type BatteryViewer struct {
	batteryName string
}

// GenerateBaristaModule generates a battery viewer barista module.
func (ba BatteryViewer) GenerateBaristaModule() (bar.Module, error) {
	mod := ba.getBatteryModule().Output(func(i battery.Info) bar.Output {
		sep := "+"
		if i.Discharging() {
			sep = "-"
		}

		return outputs.Textf("%d%% [%s]", i.RemainingPct(), sep)
	})
	return mod, nil
}

func (ba BatteryViewer) getBatteryModule() *battery.Module {
	if ba.batteryName != "" {
		return battery.Named(ba.batteryName)
	}

	return battery.All()
}
