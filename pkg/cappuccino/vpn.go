package cappuccino

import (
	"fmt"
	"image/color"
	"strings"

	"barista.run/bar"
	"barista.run/base/value"
	"barista.run/outputs"
	"barista.run/pango"
	"github.com/Wifx/gonetworkmanager"
)

// NewVPNViewer creates a new vpn viewer instance.
func NewVPNViewer(nm gonetworkmanager.NetworkManager, icons VPNIcons, colors VPNColors) VPNViewer {
	return VPNViewer{
		NetworkManager: nm,
		icons:          icons,
		colors:         colors,
	}
}

// VPNViewer is a viewer for the networkmanager module which displays vpn informations.
type VPNViewer struct {
	gonetworkmanager.NetworkManager
	icons  VPNIcons
	colors VPNColors

	formatFunc value.Value
}

// DBUSTargetEvent is the targeted dbus event.
const DBUSTargetEvent = "org.freedesktop.NetworkManager.VPN.Connection.VpnStateChanged"

// Stream receives bar sink for data streaming.
func (vv VPNViewer) Stream(s bar.Sink) {
	if vv.NetworkManager == nil {
		s.Error(fmt.Errorf("Failed to create client"))
		return
	}

	vv.fillVPNInfo(s)
	updates := vv.Subscribe()
	for {
		update := <-updates
		if update.Name != DBUSTargetEvent {
			continue
		}

		vv.fillVPNInfo(s)
	}
}

func (vv VPNViewer) fillVPNInfo(s bar.Sink) {
	conns, err := vv.GetPropertyActiveConnections()
	if err != nil {
		s.Error(err)
		return
	}

	vpns := make([]string, 0)
	for _, conn := range conns {
		isVPN, err := conn.GetPropertyVPN()
		if err != nil {
			continue
		}

		if !isVPN {
			continue
		}

		vpnName, err := conn.GetPropertyID()
		if err != nil {
			continue
		}
		vpns = append(vpns, vpnName)
	}

	hasActiveVPNs := len(vpns) > 0
	outputParts := make([]interface{}, 0, 3)
	outputParts = append(outputParts, vv.icons.VPN(hasActiveVPNs))
	if hasActiveVPNs {
		outputParts = append(outputParts, space)
		outputParts = append(outputParts, strings.Join(vpns, ","))
	}
	s.Output(outputs.Pango(outputParts...).Color(vv.colors.VPN(hasActiveVPNs)))
}

// GenerateBaristaModule generates a networkmanager viewer barista module.
func (vv VPNViewer) GenerateBaristaModule() (bar.Module, error) {
	return vv, nil
}

// VPNIcons contains all vpn related icons.
type VPNIcons interface {
	VPN(on bool) *pango.Node
}

// VPNColors provides all vpn related colors.
type VPNColors interface {
	VPN(on bool) color.Color
}
