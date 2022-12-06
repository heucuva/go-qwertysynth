package voice

import (
	"time"

	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/synth/envelope"
	"github.com/heucuva/go-qwertysynth/internal/synth/pwm"
)

type Voice interface {
	Get() (float32, bool)
	IsPlaying() bool
	Advance(amt time.Duration)
	KeyOn()
	KeyOff()
	Cut()
	SetNote(n note.Note)
	SetOutputSampleRate(sampleRate float64)
	SetFM(mod pwm.Modulator, env envelope.Envelope)
}

type voice struct {
	am       voiceOp
	fm       voiceOp
	baseNote note.Note
	curNote  note.Note
}

func NewVoice(mod pwm.Modulator, env envelope.Envelope, base note.Note) Voice {
	return &voice{
		am: voiceOp{
			mod: mod,
			env: env,
		},
		baseNote: base,
		curNote:  base,
	}
}

func (v *voice) SetFM(mod pwm.Modulator, env envelope.Envelope) {
	v.fm.mod = mod
	v.fm.env = env
}

func (v *voice) SetNote(n note.Note) {
	v.curNote = n
	v.setNoteRatio(n)
}

func (v *voice) setNoteRatio(n note.Note) {
	var ratio float64
	if f0 := v.baseNote.ToFrequency(); f0 != 0 {
		ratio = n.ToFrequency() / f0
	}
	v.am.SetNoteRatio(ratio)
}

func (v *voice) SetOutputSampleRate(sampleRate float64) {
	v.am.SetOutputSampleRate(sampleRate)
}

func (v voice) Get() (float32, bool) {
	return v.am.Get()
}

func (v voice) IsPlaying() bool {
	return v.am.IsPlaying()
}

func (v *voice) Advance(amt time.Duration) {
	mod, used := v.fm.Get()
	if used {
		n := v.curNote.AddMicrotones(scale.Microtone(mod))
		v.setNoteRatio(n)
	}

	v.am.Advance(amt)
	v.fm.Advance(amt)
}

func (v *voice) KeyOn() {
	v.am.KeyOn()
	v.fm.KeyOn()
}

func (v *voice) KeyOff() {
	v.am.KeyOff()
	v.fm.KeyOff()
}

func (v *voice) Cut() {
	v.am.Cut()
	v.fm.Cut()
}
