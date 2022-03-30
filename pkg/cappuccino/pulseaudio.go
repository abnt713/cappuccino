package cappuccino

import (
	"fmt"
	"time"

	"barista.run/bar"
	"barista.run/base/value"
	"barista.run/outputs"
	"barista.run/pango"
	"mrogalski.eu/go/pulseaudio"
)

// NewPulseAudioViewer returns a new pulseaudio viewer instance.
func NewPulseAudioViewer(logger Logger, icons PulseAudioIcons) PulseAudioViewer {
	pv := PulseAudioViewer{logger: logger, icons: icons}
	pv.formatFunc.Set(func(in string) bar.Output {
		return outputs.Text(in)
	})

	return pv
}

// PulseAudioViewer shows pulseaudio information.
type PulseAudioViewer struct {
	logger Logger
	icons  PulseAudioIcons

	cli        *pulseaudio.Client
	formatFunc value.Value
}

// Stream receives barista informations.
func (pa PulseAudioViewer) Stream(s bar.Sink) {
	var err error
	pa.cli, err = pulseaudio.NewClient()
	if err != nil {
		s.Error(err)
		pa.logger.Error("connection", err)
		return
	}

	defer pa.cli.Close()
	updates, err := pa.cli.Updates()
	if err != nil {
		s.Error(err)
		pa.logger.Error("updates fetch", err)
		return
	}

	s.Output(outputs.Text("BOOTING"))
	err = pa.attemptVolumeRetrieval(3, 1*time.Second)
	if err != nil {
		s.Error(err)
		pa.logger.Error("volume warm up", err)
		return
	}

	info, err := pa.getVolumeInfo()
	if err != nil {
		s.Error(err)
		pa.logger.Error("volume info fetch", err)
		return
	}
	s.Output(info)

	for {
		<-updates
		info, err := pa.getVolumeInfo()
		if err != nil {
			s.Error(err)
			pa.logger.Error("loop volume info fetch", err)
			break
		}
		s.Output(info)
	}
}

func (pa PulseAudioViewer) getVolumeInfo() (*bar.Segment, error) {
	isMuted, err := pa.cli.Mute()
	if err != nil {
		return nil, err
	}

	volume, err := pa.cli.Volume()
	if err != nil {
		return nil, err
	}

	outs, active, err := pa.cli.Outputs()
	if err != nil {
		return nil, err
	}

	activeOut := outs[active]

	icon := pa.icons.Sound(isMuted, volume)
	volumeStr := fmt.Sprintf("%.0f%% @ %s", volume*100, activeOut.CardName)

	return outputs.Pango(icon, space, pango.Text(volumeStr)), nil
}

func (pa PulseAudioViewer) attemptVolumeRetrieval(maxAttempts int, waitTime time.Duration) error {
	var err error
	for attempt := 0; attempt < maxAttempts; attempt++ {
		_, err = pa.cli.Volume()
		if err == nil {
			break
		}
		time.Sleep(waitTime)
	}
	return err
}

// GenerateBaristaModule generates the pulseaudio barista module.
func (pa PulseAudioViewer) GenerateBaristaModule() (bar.Module, error) {
	return pa, nil
}

// Logger logs an error event.
type Logger interface {
	Error(tag string, err error)
}

// PulseAudioIcons contains all pulseaudio related icons.
type PulseAudioIcons interface {
	Sound(muted bool, intensity float32) *pango.Node
}
