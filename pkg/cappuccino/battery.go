package cappuccino

import (
	"barista.run/bar"
	"barista.run/modules/battery"
	"barista.run/outputs"
	"barista.run/pango"
)

// NewBatteryViewer creates a new battery viewer instance.
func NewBatteryViewer(batteryName string, icon BatteryIcons) BatteryViewer {
	return BatteryViewer{
		batteryName: batteryName,
		icon:        icon,
	}
}

// BatteryViewer displays battery information
type BatteryViewer struct {
	batteryName string
	icon        BatteryIcons
}

// GenerateBaristaModule generates a battery viewer barista module.
func (ba BatteryViewer) GenerateBaristaModule() (bar.Module, error) {
	mod := ba.getBatteryModule().Output(func(i battery.Info) bar.Output {
		percentage := i.RemainingPct()
		batIcon := ba.icon.Battery(pctToBatteryLevel(percentage))
		if !i.Discharging() {
			batIcon = ba.icon.BatteryCharging()
		}

		return outputs.Pango(
			batIcon,
			space,
			pango.Textf("%d%%", percentage),
		)
	})
	return mod, nil
}

func (ba BatteryViewer) getBatteryModule() *battery.Module {
	if ba.batteryName != "" {
		return battery.Named(ba.batteryName)
	}

	return battery.All()
}

// BatteryIcons provides the battery icons.
type BatteryIcons interface {
	Battery(level BatteryLevel) *pango.Node
	BatteryCharging() *pango.Node
}

// BatteryLevel is a battery energy level.
type BatteryLevel string

// All possible battery levels.
const (
	BatteryLevelLow      = "low"
	BatteryLevelMedium   = "mid"
	BatteryLevelHigh     = "high"
	BatteryLevelFull     = "full"
	BatteryLevelCharging = "charging"
)

func pctToBatteryLevel(pct int) BatteryLevel {
	if pct <= 20 {
		return BatteryLevelLow
	}

	if pct <= 50 {
		return BatteryLevelMedium
	}

	if pct <= 90 {
		return BatteryLevelHigh
	}

	return BatteryLevelFull
}
