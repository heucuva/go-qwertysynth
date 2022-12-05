package note

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
)

type cut struct{}

func (cut) Split() (scale.Octave, scale.Key, scale.Microtone) {
	return 0, 0, 0
}

func (cut) KeyOctave() scale.KeyOctave {
	return 0
}

func (cut) IsCut() bool {
	return true
}

func (cut) IsFadeout() bool {
	return false
}

func (cut) Kind() Kind {
	return special
}

func (cut) ToFrequency(tuning tuning.Tuning) float64 {
	return 0.0
}

func (c cut) AddMicrotones(s scale.Microtone) Note {
	return c
}
