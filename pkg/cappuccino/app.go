package cappuccino

import (
	"barista.run"
	"github.com/abnt713/cappuccino/pkg"
)

// App represents the statusbar struct.
type App struct {
	modules pkg.Modules
}

// NewApp creates a new app instance.
func NewApp(modules pkg.Modules) App {
	return App{
		modules: modules,
	}
}

// Run executes the app.
func (app App) Run() error {
	mods, err := app.modules.GenerateBaristaModules()
	if err != nil {
		return err
	}
	return barista.Run(mods...)
}
