package note

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
)

type Note interface {
	Split() (scale.Octave, scale.Key, scale.Microtone)
	KeyOctave() scale.KeyOctave
	IsCut() bool
	IsFadeout() bool
	Kind() Kind
	ToFrequency(tuning tuning.Tuning) float64
	AddMicrotones(s scale.Microtone) Note
}

var (
	None    Note = nil
	Cut     Note = &cut{}
	Fadeout Note = &fadeout{}
	Base    Note = &baseNote{}
)
