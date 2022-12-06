package note

import (
	"errors"

	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
)

type Kind interface {
	Note(octave scale.Octave, key scale.Key, microtone scale.Microtone) Note
	BaseFrequency() float64
	BaseNote() Note
	ParseNote(str string) (Note, error)
}

type specialKind struct{}

func (specialKind) Note(o scale.Octave, k scale.Key, s scale.Microtone) Note {
	return None
}

func (specialKind) BaseFrequency() float64 {
	return 0.0
}

func (specialKind) BaseNote() Note {
	return None
}

func (specialKind) ParseNote(str string) (Note, error) {
	return None, errors.New("cannot construct special note through ParseNote")
}

var special Kind = &specialKind{}
