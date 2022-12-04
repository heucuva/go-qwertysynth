package note

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
)

type baseNote struct{}

func (baseNote) Split() (keyoctave.Octave, keyoctave.Key, keyoctave.Semitone) {
	return 4, keyoctave.KeyA, 0
}

func (baseNote) KeyOctave() keyoctave.KeyOctave {
	return keyoctave.NewKeyOctave(keyoctave.KeyA, 4)
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

func (baseNote) ToFrequency() float64 {
	return scale.A440_A4Frequency
}

func (b baseNote) AddSemitones(s keyoctave.Semitone) Note {
	return b
}
