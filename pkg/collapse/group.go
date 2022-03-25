package collapse

import (
	"barista.run/bar"
	"barista.run/group/collapsing"
	"github.com/abnt713/cappuccino/pkg"
)

// NewGroup creates a new group instance.
func NewGroup(
	modules pkg.Modules,
	collapseSymbol,
	expandSymbol pkg.Output,
) *Group {
	return &Group{
		modules:        modules,
		collapseSymbol: collapseSymbol,
		expandSymbol:   expandSymbol,
	}
}

// Group is a collapsable group.
type Group struct {
	modules        pkg.Modules
	collapseSymbol pkg.Output
	expandSymbol   pkg.Output

	ctrl collapsing.Controller
}

// GenerateBaristaModule generates the group.
func (gr *Group) GenerateBaristaModule() (bar.Module, error) {
	baristaMods := make([]bar.Module, 0, len(gr.modules))
	for _, mod := range gr.modules {
		baristaMod, err := mod.GenerateBaristaModule()
		if err != nil {
			return nil, err
		}
		baristaMods = append(baristaMods, baristaMod)
	}

	var baristaGroup bar.Module
	baristaGroup, gr.ctrl = collapsing.Group(baristaMods...)
	gr.ctrl.ButtonFunc(gr.buttonFunc)
	return baristaGroup, nil
}

func (gr *Group) buttonFunc(ctrl collapsing.Controller) (start, end bar.Output) {
	if ctrl.Expanded() {
		return nil, gr.collapseSymbol.GenerateBaristaOutput()
	}
	return gr.expandSymbol.GenerateBaristaOutput(), nil
}
