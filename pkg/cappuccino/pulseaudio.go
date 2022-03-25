package cappuccino

import (
	"fmt"
	"time"

	"barista.run/bar"
	"barista.run/base/value"
	"barista.run/outputs"
	"mrogalski.eu/go/pulseaudio"
)

// PulseAudioViewer shows pulseaudio information.
type PulseAudioViewer struct {
	cli        *pulseaudio.Client
	formatFunc value.Value
}

// NewPulseAudioViewer returns a new pulseaudio viewer instance.
func NewPulseAudioViewer() PulseAudioViewer {
	pv := PulseAudioViewer{}
	pv.formatFunc.Set(func(in string) bar.Output {
		return outputs.Text(in)
	})

	return pv
}

// Stream receives barista informations.
func (pa PulseAudioViewer) Stream(s bar.Sink) {
	const maxAttempts = 3

	var err error
	for attempts := 1; pa.cli != nil && attempts <= maxAttempts; attempts++ {
		pa.cli, err = pulseaudio.NewClient()
		if err != nil && attempts == maxAttempts {
			s.Error(err)
			return
		}
		time.Sleep(2 * time.Second)
	}
	defer pa.cli.Close()
	updates, err := pa.cli.Updates()
	if err != nil {
		s.Error(err)
		return
	}

	info, err := pa.getVolumeInfo()
	if err != nil {
		s.Error(err)
		return
	}

	pa.output(info, s)

	for {
		<-updates
		info, err := pa.getVolumeInfo()
		if err != nil {
			s.Error(err)
			break
		}
		pa.output(info, s)
	}
}

func (pa PulseAudioViewer) getVolumeInfo() (string, error) {
	volume, err := pa.cli.Volume()
	if err != nil {
		return "", err
	}

	outputs, active, err := pa.cli.Outputs()
	if err != nil {
		return "", err
	}

	activeOut := outputs[active]
	return fmt.Sprintf("%.0f%% @ %s", volume*100, activeOut.CardName), nil
}

func (pa PulseAudioViewer) output(in string, s bar.Sink) {
	format := pa.formatFunc.Get().(func(string) bar.Output)
	s.Output(format(in))
}

// GenerateBaristaModule generates the pulseaudio barista module.
func (pa PulseAudioViewer) GenerateBaristaModule() (bar.Module, error) {
	return pa, nil
}
