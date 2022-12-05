package keymap

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/synth/wavetable"
)

type KeyMap [scale.TotalKeyOctaves]wavetable.Index
