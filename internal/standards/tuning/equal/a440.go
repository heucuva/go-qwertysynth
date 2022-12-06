package equal

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
)

const (
	A440_C4Frequency      = A440_CSharp4Frequency / twelfthRoot2
	A440_CSharp4Frequency = A440_D4Frequency / twelfthRoot2
	A440_D4Frequency      = A440_DSharp4Frequency / twelfthRoot2
	A440_DSharp4Frequency = A440_E4Frequency / twelfthRoot2
	A440_E4Frequency      = A440_F4Frequency / twelfthRoot2
	A440_F4Frequency      = A440_FSharp4Frequency / twelfthRoot2
	A440_FSharp4Frequency = A440_G4Frequency / twelfthRoot2
	A440_G4Frequency      = A440_GSharp4Frequency / twelfthRoot2
	A440_GSharp4Frequency = A440_A4Frequency / twelfthRoot2
	A440_A4Frequency      = 440.0
	A440_ASharp4Frequency = A440_A4Frequency * twelfthRoot2
	A440_B4Frequency      = A440_ASharp4Frequency * twelfthRoot2
)

type a440 struct{}

var A440 tuning.Tuning = &a440{}

var a440_scale = [TwelveKeysPerOctave]float64{
	A440_C4Frequency,
	A440_CSharp4Frequency,
	A440_D4Frequency,
	A440_DSharp4Frequency,
	A440_E4Frequency,
	A440_F4Frequency,
	A440_FSharp4Frequency,
	A440_G4Frequency,
	A440_GSharp4Frequency,
	A440_A4Frequency,
	A440_ASharp4Frequency,
	A440_B4Frequency,
}

func (a440) ToFrequency(ko tuning.KeyOctave) float64 {
	k, o := ko.Split(A440)
	freq := a440_scale[k.Index()]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}

func (a440) Key(index int) scale.Key {
	return TwelveKey(index)
}

func (a440) BaseKey() (scale.Key, scale.Octave) {
	return TwelveKeyA, 4
}

func (a440) KeysPerOctave() int {
	return TwelveKeysPerOctave
}
