package keymap

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/synth/wavetable"
)

var Default = KeyMap{}

func init() {
	var i wavetable.Index
	for o := scale.MinOctave; o <= scale.MaxOctave; o++ {
		for k := scale.MinKey; k <= scale.MaxKey; k++ {
			Default[scale.KeyOctave(i)] = i
			i++
		}
	}
}
