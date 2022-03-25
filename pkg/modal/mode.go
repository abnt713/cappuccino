package modal

import (
	"barista.run/bar"
	"barista.run/group/modal"
	"github.com/abnt713/cappuccino/pkg"
)

// NewMode creates a new Mode instance.
func NewMode(
	name string,
	button pkg.Segment,
	actions ...ModeAction,
) Mode {
	return Mode{
		name:    name,
		button:  button,
		actions: actions,
	}
}

// Mode builds a set of modules.
type Mode struct {
	name    string
	button  pkg.Segment
	actions []ModeAction
}

// Apply applies the mode to a modal.
func (mo Mode) Apply(baristaModal ModeCreator) error {
	barMode := baristaModal.CreateMode(mo.name)
	if mo.button != nil {
		barMode = barMode.SetOutput(mo.button.GenerateBaristaSegment())
	}
	for _, act := range mo.actions {
		var err error
		barMode, err = act.apply(barMode)
		if err != nil {
			return err
		}
	}
	return nil
}

// ModeAction applies a module to a mode.
type ModeAction interface {
	apply(barMode *modal.Mode) (*modal.Mode, error)
}

// AddAction creates an add action.
func AddAction(mod pkg.Module) Action {
	return Action{
		mod: mod,
		applyFn: func(barMode *modal.Mode, mod bar.Module) *modal.Mode {
			return barMode.Add(mod)
		},
	}
}

// DetailAction creates a detail action.
func DetailAction(mod pkg.Module) Action {
	return Action{
		mod: mod,
		applyFn: func(barMode *modal.Mode, mod bar.Module) *modal.Mode {
			return barMode.Add(mod)
		},
	}
}

// SummaryAction creates a summary action.
func SummaryAction(mod pkg.Module) Action {
	return Action{
		mod: mod,
		applyFn: func(barMode *modal.Mode, mod bar.Module) *modal.Mode {
			return barMode.Summary(mod)
		},
	}
}

// Action acts uppon a module towards a barMode.
type Action struct {
	mod     pkg.Module
	applyFn func(*modal.Mode, bar.Module) *modal.Mode
}

func (aa Action) apply(barMode *modal.Mode) (*modal.Mode, error) {
	mod, err := aa.mod.GenerateBaristaModule()
	if err != nil {
		return nil, err
	}
	return aa.applyFn(barMode, mod), nil
}

// ModeCreator creates a modal mode.
type ModeCreator interface {
	CreateMode(name string) *modal.Mode
}
