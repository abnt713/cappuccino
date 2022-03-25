package modal

import (
	"barista.run/bar"
	"barista.run/group/modal"
)

// NewModal creates a new modal instance.
func NewModal() *Modal {
	return &Modal{
		baristaModal: modal.New(),
	}
}

// Modal is a bar modal.
type Modal struct {
	baristaModal *modal.Modal
	modalCtrl    modal.Controller
}

//GenerateBaristaModule generates the barista modal.
func (mo *Modal) GenerateBaristaModule() (bar.Module, error) {
	var mod bar.Module
	mod, mo.modalCtrl = mo.baristaModal.Build()
	return mod, nil
}

// CreateMode creates a new mode.
func (mo *Modal) CreateMode(name string) *modal.Mode {
	return mo.baristaModal.Mode(name)
}

// ToggleMode toggles a modal mode.
func (mo *Modal) ToggleMode(name string) {
	mo.modalCtrl.Toggle(name)
}
