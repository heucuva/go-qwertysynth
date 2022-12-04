package keymap

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"
	"github.com/heucuva/go-qwertysynth/internal/synth/wavetable"
)

var Default = KeyMap{}

func init() {
	var i wavetable.Index
	for o := keyoctave.MinOctave; o <= keyoctave.MaxOctave; o++ {
		for k := keyoctave.MinKey; k <= keyoctave.MaxKey; k++ {
			Default[keyoctave.KeyOctave(i)] = i
			i++
		}
	}
}
