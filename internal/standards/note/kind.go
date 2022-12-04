package note

import "github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"

type Kind interface {
	Note(octave keyoctave.Octave, key keyoctave.Key, semitone keyoctave.Semitone) Note
	BaseFrequency() float64
	CenterNote() Note
}

type specialKind struct{}

func (specialKind) Note(o keyoctave.Octave, k keyoctave.Key, s keyoctave.Semitone) Note {
	return None
}

func (specialKind) BaseFrequency() float64 {
	return 0.0
}

func (specialKind) CenterNote() Note {
	return None
}

var special Kind = &specialKind{}
