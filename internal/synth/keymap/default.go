package keymap

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	tuningPkg "github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	"github.com/heucuva/go-qwertysynth/internal/synth/wavetable"
)

func Default(tuning tuningPkg.Tuning) KeyMap {
	km := make(KeyMap)
	var i wavetable.Index
	for o := scale.MinOctave; o <= scale.MaxOctave; o++ {
		for k := 0; k < tuning.KeysPerOctave(); k++ {
			ko := tuningPkg.NewKeyOctave(tuning.Key(k), o)
			km[ko] = i
			i++
		}
	}
	return km
}
