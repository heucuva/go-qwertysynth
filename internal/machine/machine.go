package machine

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/synth/wave"
)

type Machine interface {
	Default() Default
	Note(o keyoctave.Octave, k keyoctave.Key, s keyoctave.Semitone) note.Note
	Generate(generator wave.Generator, opts ...wave.GeneratorParam) (wave.Wave, error)
}
