package cappuccino

import (
	"barista.run/bar"
	"barista.run/colors"
	"barista.run/modules/netinfo"
	"barista.run/outputs"
)

// NetViewMode sets the view mode for the module.
type NetViewMode string

// All possible view modes.
const (
	NetViewModeDefault = NetViewMode("default")
	NetViewModeVPN     = NetViewMode("vpn")
)

// NewNetViewer creates a new net viewer instance.
func NewNetViewer(viewMode NetViewMode) NetViewer {
	return NetViewer{
		viewMode: viewMode,
	}
}

// NetViewer views Net information.
type NetViewer struct {
	viewMode NetViewMode
}

// GenerateBaristaModule generates the VPN viewer module.
func (nv NetViewer) GenerateBaristaModule() (bar.Module, error) {
	switch nv.viewMode {
	case NetViewModeVPN:
		return netinfo.Prefix("tun").Output(nv.vpnOutput), nil
	default:
		return netinfo.New().Output(nv.networkOutput), nil
	}
}

func (nv NetViewer) networkOutput(s netinfo.State) bar.Output {
	if len(s.IPs) < 1 {
		return outputs.Text("No network").Color(colors.Scheme("bad"))
	}
	return outputs.Textf("%s: %v", s.Name, s.IPs[0])
}

func (nv NetViewer) vpnOutput(s netinfo.State) bar.Output {
	if s.Connected() {
		return outputs.Text("VPN").Color(colors.Scheme("good"))
	}

	if s.Connecting() {
		return outputs.Text("...").Color(colors.Scheme("good"))
	}

	return outputs.Text("---").Color(colors.Scheme("good"))
}
