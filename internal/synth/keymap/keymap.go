package keymap

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"
	"github.com/heucuva/go-qwertysynth/internal/synth/wavetable"
)

type KeyMap [keyoctave.TotalKeyOctaves]wavetable.Index
