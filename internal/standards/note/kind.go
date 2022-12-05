package note

import "github.com/heucuva/go-qwertysynth/internal/standards/scale"

type Kind interface {
	Note(octave scale.Octave, key scale.Key, microtone scale.Microtone) Note
	BaseFrequency() float64
	CenterNote() Note
}

type specialKind struct{}

func (specialKind) Note(o scale.Octave, k scale.Key, s scale.Microtone) Note {
	return None
}

func (specialKind) BaseFrequency() float64 {
	return 0.0
}

func (specialKind) CenterNote() Note {
	return None
}

var special Kind = &specialKind{}
