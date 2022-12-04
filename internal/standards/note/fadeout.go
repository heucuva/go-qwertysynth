package note

import "github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"

type fadeout struct{}

func (fadeout) Split() (keyoctave.Octave, keyoctave.Key, keyoctave.Semitone) {
	return 0, 0, 0
}

func (fadeout) KeyOctave() keyoctave.KeyOctave {
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

func (fadeout) ToFrequency() float64 {
	return 0.0
}

func (f fadeout) AddSemitones(s keyoctave.Semitone) Note {
	return f
}
