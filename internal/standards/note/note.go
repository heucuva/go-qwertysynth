package note

import "github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"

type Note interface {
	Split() (keyoctave.Octave, keyoctave.Key, keyoctave.Semitone)
	KeyOctave() keyoctave.KeyOctave
	IsCut() bool
	IsFadeout() bool
	Kind() Kind
	ToFrequency() float64
	AddSemitones(s keyoctave.Semitone) Note
}

var (
	None    Note = nil
	Cut     Note = &cut{}
	Fadeout Note = &fadeout{}
	Base    Note = &baseNote{}
)
