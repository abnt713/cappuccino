package cappuccino

import (
	"image/color"

	"barista.run/bar"
	"barista.run/modules/battery"
	"barista.run/outputs"
	"barista.run/pango"
)

// NewBatteryViewer creates a new battery viewer instance.
func NewBatteryViewer(batteryName string, icon BatteryIcons, colors BatteryColors) BatteryViewer {
	return BatteryViewer{
		batteryName: batteryName,
		icon:        icon,
		colors:      colors,
	}
}

// BatteryViewer displays battery information
type BatteryViewer struct {
	batteryName string
	icon        BatteryIcons
	colors      BatteryColors
}

// GenerateBaristaModule generates a battery viewer barista module.
func (ba BatteryViewer) GenerateBaristaModule() (bar.Module, error) {
	mod := ba.getBatteryModule().Output(func(i battery.Info) bar.Output {
		percentage := i.RemainingPct()
		lvl := pctToBatteryLevel(percentage)
		isCharging := !i.Discharging()
		batIcon := ba.icon.Battery(lvl, isCharging)
		batColor := ba.colors.Battery(lvl, isCharging)

		return outputs.Pango(
			batIcon,
			space,
			pango.Textf("%d%%", percentage),
		).Color(batColor)
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
	Battery(level BatteryLevel, isCharging bool) *pango.Node
}

// BatteryColors provides the battery colors.
type BatteryColors interface {
	Battery(level BatteryLevel, isCharging bool) color.Color
}

// BatteryLevel is a battery energy level.
type BatteryLevel string

// All possible battery levels.
const (
	BatteryLevelLow    = BatteryLevel("low")
	BatteryLevelMedium = BatteryLevel("mid")
	BatteryLevelHigh   = BatteryLevel("high")
	BatteryLevelFull   = BatteryLevel("full")
)

func pctToBatteryLevel(pct int) BatteryLevel {
	if pct <= 30 {
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
