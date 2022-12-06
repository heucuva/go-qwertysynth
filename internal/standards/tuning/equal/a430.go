package equal

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
)

const (
	A430_C4Frequency      = A430_CSharp4Frequency / twelfthRoot2
	A430_CSharp4Frequency = A430_D4Frequency / twelfthRoot2
	A430_D4Frequency      = A430_DSharp4Frequency / twelfthRoot2
	A430_DSharp4Frequency = A430_E4Frequency / twelfthRoot2
	A430_E4Frequency      = A430_F4Frequency / twelfthRoot2
	A430_F4Frequency      = A430_FSharp4Frequency / twelfthRoot2
	A430_FSharp4Frequency = A430_G4Frequency / twelfthRoot2
	A430_G4Frequency      = A430_GSharp4Frequency / twelfthRoot2
	A430_GSharp4Frequency = A430_A4Frequency / twelfthRoot2
	A430_A4Frequency      = 430.0
	A430_ASharp4Frequency = A430_A4Frequency * twelfthRoot2
	A430_B4Frequency      = A430_ASharp4Frequency * twelfthRoot2
)

type a430 struct{}

var A430 tuning.Tuning = &a430{}

var a430_scale = [TwelveKeysPerOctave]float64{
	A430_C4Frequency,
	A430_CSharp4Frequency,
	A430_D4Frequency,
	A430_DSharp4Frequency,
	A430_E4Frequency,
	A430_F4Frequency,
	A430_FSharp4Frequency,
	A430_G4Frequency,
	A430_GSharp4Frequency,
	A430_A4Frequency,
	A430_ASharp4Frequency,
	A430_B4Frequency,
}

func (a430) ToFrequency(ko tuning.KeyOctave) float64 {
	k, o := ko.Split(A430)
	freq := a430_scale[k.Index()]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}

func (a430) Key(index int) scale.Key {
	return TwelveKey(index)
}

func (a430) BaseKey() (scale.Key, scale.Octave) {
	return TwelveKeyA, 4
}

func (a430) KeysPerOctave() int {
	return TwelveKeysPerOctave
}
