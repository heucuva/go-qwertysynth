package keymap

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	"github.com/heucuva/go-qwertysynth/internal/synth/wavetable"
)

type KeyMap map[tuning.KeyOctave]wavetable.Index
