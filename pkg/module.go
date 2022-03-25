package pkg

import (
	"barista.run/bar"
)

// Module is an app module.
type Module interface {
	GenerateBaristaModule() (bar.Module, error)
}

// Modules is a collection of modules.
type Modules []Module

// GenerateBaristaModules generates all collection modules.
func (m Modules) GenerateBaristaModules() ([]bar.Module, error) {
	barMods := make([]bar.Module, 0, len(m))
	for _, mod := range m {
		barMod, err := mod.GenerateBaristaModule()
		if err != nil {
			return nil, err
		}
		barMods = append(barMods, barMod)
	}

	return barMods, nil
}

// Output generates an output.
type Output interface {
	GenerateBaristaOutput() bar.Output
}

// Segment generates a segment.
type Segment interface {
	GenerateBaristaSegment() *bar.Segment
}
