package note

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
)

type fadeout struct{}

func (fadeout) Split() (scale.Octave, scale.Key, scale.Microtone) {
	return 0, 0, 0
}

func (fadeout) KeyOctave() scale.KeyOctave {
	return 0
}

func (fadeout) IsCut() bool {
	return false
}

func (fadeout) IsFadeout() bool {
	return true
}

func (fadeout) Kind() Kind {
	return special
}

func (fadeout) ToFrequency(tuning tuning.Tuning) float64 {
	return 0.0
}

func (f fadeout) AddMicrotones(s scale.Microtone) Note {
	return f
}
