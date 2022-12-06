package equal

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
)

const (
	A429_C4Frequency      = A429_CSharp4Frequency / twelfthRoot2
	A429_CSharp4Frequency = A429_D4Frequency / twelfthRoot2
	A429_D4Frequency      = A429_DSharp4Frequency / twelfthRoot2
	A429_DSharp4Frequency = A429_E4Frequency / twelfthRoot2
	A429_E4Frequency      = A429_F4Frequency / twelfthRoot2
	A429_F4Frequency      = A429_FSharp4Frequency / twelfthRoot2
	A429_FSharp4Frequency = A429_G4Frequency / twelfthRoot2
	A429_G4Frequency      = A429_GSharp4Frequency / twelfthRoot2
	A429_GSharp4Frequency = A429_A4Frequency / twelfthRoot2
	A429_A4Frequency      = 429.0
	A429_ASharp4Frequency = A429_A4Frequency * twelfthRoot2
	A429_B4Frequency      = A429_ASharp4Frequency * twelfthRoot2
)

type a429 struct{}

var A429 tuning.Tuning = &a429{}

var a429_scale = [TwelveKeysPerOctave]float64{
	A429_C4Frequency,
	A429_CSharp4Frequency,
	A429_D4Frequency,
	A429_DSharp4Frequency,
	A429_E4Frequency,
	A429_F4Frequency,
	A429_FSharp4Frequency,
	A429_G4Frequency,
	A429_GSharp4Frequency,
	A429_A4Frequency,
	A429_ASharp4Frequency,
	A429_B4Frequency,
}

func (a429) ToFrequency(ko tuning.KeyOctave) float64 {
	k, o := ko.Split(A429)
	freq := a429_scale[k.Index()]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}

func (a429) Key(index int) scale.Key {
	return TwelveKey(index)
}

func (a429) BaseKey() (scale.Key, scale.Octave) {
	return TwelveKeyA, 4
}

func (a429) KeysPerOctave() int {
	return TwelveKeysPerOctave
}
