package voice

import (
	"time"

	"github.com/heucuva/go-qwertysynth/internal/synth/envelope"
	"github.com/heucuva/go-qwertysynth/internal/synth/pwm"
)

type voiceOp struct {
	mod pwm.Modulator
	env envelope.Envelope
}

func (v voiceOp) Get() (float32, bool) {
	if v.mod == nil || v.env == nil {
		return 0, false
	}

	mod, playing := v.env.Get()
	if !playing {
		return 0, false
	}

	pre := v.mod.Get()
	return pre * mod, true
}

func (v voiceOp) IsPlaying() bool {
	return v.env != nil && v.env.IsPlaying()
}

func (v *voiceOp) Advance(sampDur time.Duration) {
	if v.mod != nil {
		v.mod.Advance()
	}
	if v.env != nil {
		v.env.Advance(sampDur)
	}
}

func (v *voiceOp) SetNoteRatio(ratio float64) {
	if v.mod != nil {
		v.mod.SetNoteRatio(ratio)
	}
}

func (v *voiceOp) SetOutputSampleRate(sampleRate float64) {
	if v.mod != nil {
		v.mod.SetOutputSampleRate(sampleRate)
	}
}

func (v *voiceOp) SetFreqModulation(mod float64) {
	if v.mod != nil {
		v.mod.SetFreqModulation(mod)
	}
}

func (v *voiceOp) KeyOn() {
	if v.env != nil {
		v.env.KeyOn()
	}
}

func (v *voiceOp) KeyOff() {
	if v.env != nil {
		v.env.KeyOff()
	}
}

func (v *voiceOp) Cut() {
	if v.env != nil {
		v.env.Cut()
	}
}
