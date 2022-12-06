package machine

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	"github.com/heucuva/go-qwertysynth/internal/synth/wave"
)

type Machine interface {
	Default() Default
	Generate(generator wave.Generator, opts ...wave.GeneratorParam) (wave.Wave, error)
	Tuning() tuning.Tuning
	note.Kind
}
