package wavetable

import (
	"github.com/heucuva/go-qwertysynth/internal/machine"
	"github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/synth/voice"
)

type WaveTable interface {
	Get(n note.Note) voice.Voice
	Default() machine.Default
	Note(o keyoctave.Octave, k keyoctave.Key, s keyoctave.Semitone) note.Note
}
