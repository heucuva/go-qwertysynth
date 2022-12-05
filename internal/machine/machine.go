package machine

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	"github.com/heucuva/go-qwertysynth/internal/synth/wave"
)

type Machine interface {
	Default() Default
	Note(o scale.Octave, k scale.Key, s scale.Microtone) note.Note
	Generate(tuning tuning.Tuning, generator wave.Generator, opts ...wave.GeneratorParam) (wave.Wave, error)
}
