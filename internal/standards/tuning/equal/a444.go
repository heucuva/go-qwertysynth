package equal

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
)

const (
	A444_C4Frequency      = A444_CSharp4Frequency / twelfthRoot2
	A444_CSharp4Frequency = A444_D4Frequency / twelfthRoot2
	A444_D4Frequency      = A444_DSharp4Frequency / twelfthRoot2
	A444_DSharp4Frequency = A444_E4Frequency / twelfthRoot2
	A444_E4Frequency      = A444_F4Frequency / twelfthRoot2
	A444_F4Frequency      = A444_FSharp4Frequency / twelfthRoot2
	A444_FSharp4Frequency = A444_G4Frequency / twelfthRoot2
	A444_G4Frequency      = A444_GSharp4Frequency / twelfthRoot2
	A444_GSharp4Frequency = A444_A4Frequency / twelfthRoot2
	A444_A4Frequency      = 444.0
	A444_ASharp4Frequency = A444_A4Frequency * twelfthRoot2
	A444_B4Frequency      = A444_ASharp4Frequency * twelfthRoot2
)

type a444 struct{}

var A444 tuning.Tuning = &a444{}

var a444_scale = [TwelveKeysPerOctave]float64{
	A444_C4Frequency,
	A444_CSharp4Frequency,
	A444_D4Frequency,
	A444_DSharp4Frequency,
	A444_E4Frequency,
	A444_F4Frequency,
	A444_FSharp4Frequency,
	A444_G4Frequency,
	A444_GSharp4Frequency,
	A444_A4Frequency,
	A444_ASharp4Frequency,
	A444_B4Frequency,
}

func (a444) ToFrequency(ko tuning.KeyOctave) float64 {
	k, o := ko.Split(A444)
	freq := a444_scale[k.Index()]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}

func (a444) Key(index int) scale.Key {
	return TwelveKey(index)
}

func (a444) BaseKey() (scale.Key, scale.Octave) {
	return TwelveKeyA, 4
}

func (a444) KeysPerOctave() int {
	return TwelveKeysPerOctave
}
