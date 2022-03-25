package cappuccino

import (
	"fmt"
	"strings"

	"barista.run/bar"
	"barista.run/base/value"
	"barista.run/outputs"
	"github.com/Wifx/gonetworkmanager"
)

// NewNetworkManagerViewer creates a new networkmanager viewer instance.
func NewNetworkManagerViewer() NetworkManagerViewer {
	nm, err := gonetworkmanager.NewNetworkManager()
	if err != nil {
		return NetworkManagerViewer{
			NetworkManager: nil,
		}
	}

	return NetworkManagerViewer{
		NetworkManager: nm,
	}
}

// NetworkManagerViewer is a viewer for the networkmanager module.
type NetworkManagerViewer struct {
	gonetworkmanager.NetworkManager

	formatFunc value.Value
}

// DBUSTargetEvent is the targeted dbus event.
const DBUSTargetEvent = "org.freedesktop.NetworkManager.VPN.Connection.VpnStateChanged"

// Stream receives bar sink for data streaming.
func (nm NetworkManagerViewer) Stream(s bar.Sink) {
	if nm.NetworkManager == nil {
		s.Error(fmt.Errorf("Failed to create client"))
		return
	}

	nm.fillVPNInfo(s)
	updates := nm.Subscribe()
	for {
		update := <-updates
		if update.Name != DBUSTargetEvent {
			continue
		}

		nm.fillVPNInfo(s)
	}
}

func (nm NetworkManagerViewer) fillVPNInfo(s bar.Sink) {
	conns, err := nm.GetPropertyActiveConnections()
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

	vpnStr := "---"
	if len(vpns) > 0 {
		vpnStr = strings.Join(vpns, ",")
	}

	s.Output(outputs.Text(vpnStr))
}

// GenerateBaristaModule generates a networkmanager viewer barista module.
func (nm NetworkManagerViewer) GenerateBaristaModule() (bar.Module, error) {
	return nm, nil
}
