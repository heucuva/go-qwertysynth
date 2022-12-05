package wavetable

import (
	"github.com/heucuva/go-qwertysynth/internal/machine"
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/synth/voice"
)

type WaveTable interface {
	Get(n note.Note) voice.Voice
	Default() machine.Default
	Note(o scale.Octave, k scale.Key, s scale.Microtone) note.Note
}
