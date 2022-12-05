package wavetable

import (
	"github.com/heucuva/go-qwertysynth/internal/machine"
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	"github.com/heucuva/go-qwertysynth/internal/synth/envelope"
	"github.com/heucuva/go-qwertysynth/internal/synth/pwm"
	"github.com/heucuva/go-qwertysynth/internal/synth/voice"
	"github.com/heucuva/go-qwertysynth/internal/synth/wave"
	"github.com/pkg/errors"
)

type Component struct {
	Generator wave.Generator
	Options   []wave.GeneratorParam
	Envelope  envelope.ADSR
}

type Operator struct {
	Amplitude Component
	Frequency Component
}

type Op struct {
	Wave     wave.Wave
	Envelope envelope.ADSR
}

type Item struct {
	AM       Op
	FM       Op
	BaseNote note.Note
	Tuning   tuning.Tuning
}

func (i Item) Voice() voice.Voice {
	env := envelope.NewADSR(i.AM.Envelope)
	v := voice.NewVoice(pwm.NewModulator(i.AM.Wave, i.AM.Wave.SampleRate()), env, i.BaseNote, i.Tuning)
	if i.FM.Wave != nil {
		env := envelope.NewADSR(i.FM.Envelope)
		v.SetFM(pwm.NewModulator(i.FM.Wave, i.FM.Wave.SampleRate()), env)
	}
	return v
}

func Generate(mach machine.Machine, op Operator, tuning tuning.Tuning) ([]*Item, error) {
	if tuning == nil {
		tuning = mach.Default().Tuning()
	}

	am, err := mach.Generate(tuning, op.Amplitude.Generator, op.Amplitude.Options...)
	if err != nil {
		return nil, errors.Wrap(err, "amplitude modulator")
	}

	var fm wave.Wave
	if op.Frequency.Generator != nil {
		fm, err = mach.Generate(tuning, op.Frequency.Generator, op.Frequency.Options...)
		if err != nil {
			return nil, errors.Wrap(err, "frequency modulator")
		}
	}

	it := &Item{
		AM: Op{
			Wave:     am,
			Envelope: op.Amplitude.Envelope,
		},
		FM: Op{
			Wave:     fm,
			Envelope: op.Frequency.Envelope,
		},
		BaseNote: mach.Default().CenterNote(),
		Tuning:   tuning,
	}

	t := make([]*Item, 0, scale.TotalKeyOctaves)
	for o := scale.MinOctave; o <= scale.MaxOctave; o++ {
		for k := scale.MinKey; k <= scale.MaxKey; k++ {
			t = append(t, it)
		}
	}
	return t, nil
}
