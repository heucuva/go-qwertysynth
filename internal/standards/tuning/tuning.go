package tuning

import "github.com/heucuva/go-qwertysynth/internal/standards/scale"

type Tuning interface {
	ToFrequency(ko KeyOctave) float64
	Key(index int) scale.Key
	BaseKey() (scale.Key, scale.Octave)
	KeysPerOctave() int
}
