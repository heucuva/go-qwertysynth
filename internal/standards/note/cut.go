package note

import "github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"

type cut struct{}

func (cut) Split() (keyoctave.Octave, keyoctave.Key, keyoctave.Semitone) {
	return 0, 0, 0
}

func (cut) KeyOctave() keyoctave.KeyOctave {
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

func (cut) ToFrequency() float64 {
	return 0.0
}

func (c cut) AddSemitones(s keyoctave.Semitone) Note {
	return c
}
