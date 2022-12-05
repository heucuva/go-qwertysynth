package note

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	equalTuning "github.com/heucuva/go-qwertysynth/internal/standards/tuning/equal"
)

type baseNote struct{}

func (baseNote) Split() (scale.Octave, scale.Key, scale.Microtone) {
	return 4, scale.KeyA, 0
}

func (baseNote) KeyOctave() scale.KeyOctave {
	return scale.NewKeyOctave(scale.KeyA, 4)
}

func (baseNote) IsCut() bool {
	return false
}

func (baseNote) IsFadeout() bool {
	return false
}

func (baseNote) Kind() Kind {
	return special
}

func (baseNote) ToFrequency(tuning tuning.Tuning) float64 {
	if tuning == nil {
		tuning = equalTuning.A440
	}
	return tuning.ToFrequency(scale.NewKeyOctave(scale.KeyA, 4))
}

func (b baseNote) AddMicrotones(s scale.Microtone) Note {
	return b
}
