//go:build !windows || pulseaudio
// +build !windows pulseaudio

package device

import (
	"context"

	"github.com/gotracker/gomixing/mixing"
	"github.com/gotracker/gomixing/sampling"
	"github.com/pkg/errors"

	deviceCommon "github.com/heucuva/go-qwertysynth/internal/output/device/common"
	"github.com/heucuva/go-qwertysynth/internal/output/device/pulseaudio"
	"github.com/heucuva/go-qwertysynth/internal/output/premix"
)

const pulseaudioName = "pulseaudio"

type pulseaudioDevice struct {
	device
	mix      mixing.Mixer
	sampFmt  sampling.Format
	settings deviceCommon.Settings
	pa       *pulseaudio.Client
}

func (pulseaudioDevice) GetKind() deviceCommon.Kind {
	return deviceCommon.KindSoundCard
}

// Name returns the device name
func (pulseaudioDevice) Name() string {
	return pulseaudioName
}

func newPulseAudioDevice(settings deviceCommon.Settings) (Device, error) {
	d := pulseaudioDevice{
		device: device{
			onRowOutput: settings.OnRowOutput,
		},
		mix: mixing.Mixer{
			Channels: settings.Channels,
		},
		settings: settings,
	}

	switch settings.BitsPerSample {
	case 8:
		d.sampFmt = sampling.Format8BitUnsigned
	case 16:
		d.sampFmt = sampling.Format16BitLESigned
	}

	return &d, nil
}

// Play starts the wave output device playing
func (d *pulseaudioDevice) Play(in <-chan *premix.PremixData) error {
	return d.PlayWithCtx(context.Background(), in)
}

// PlayWithCtx starts the wave output device playing
func (d *pulseaudioDevice) PlayWithCtx(ctx context.Context, in <-chan *premix.PremixData) error {
	play, err := pulseaudio.New(ctx, "Music", d.settings.SamplesPerSecond, d.settings.Channels, d.settings.BitsPerSample)
	if err != nil {
		return err
	}
	d.pa = play

	panmixer := mixing.GetPanMixer(d.mix.Channels)
	if panmixer == nil {
		return errors.New("invalid pan mixer - check channel count")
	}

	myCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	for {
		select {
		case <-myCtx.Done():
			return myCtx.Err()
		case row, ok := <-in:
			if !ok {
				return nil
			}
			mixedData := d.mix.Flatten(panmixer, row.SamplesLen, row.Data, row.MixerVolume, d.sampFmt)
			d.pa.Output(mixedData)
			if d.onRowOutput != nil {
				d.onRowOutput(deviceCommon.KindSoundCard, row)
			}
		}
	}
}

// Close closes the wave output device
func (d *pulseaudioDevice) Close() error {
	if d.pa != nil {
		return d.pa.Close()
	}
	return nil
}

func init() {
	Map[pulseaudioName] = deviceDetails{
		create: newPulseAudioDevice,
		Kind:   deviceCommon.KindSoundCard,
	}
}
