package equal

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
)

const (
	A466_C4Frequency      = A466_CSharp4Frequency / twelfthRoot2
	A466_CSharp4Frequency = A466_D4Frequency / twelfthRoot2
	A466_D4Frequency      = A466_DSharp4Frequency / twelfthRoot2
	A466_DSharp4Frequency = A466_E4Frequency / twelfthRoot2
	A466_E4Frequency      = A466_F4Frequency / twelfthRoot2
	A466_F4Frequency      = A466_FSharp4Frequency / twelfthRoot2
	A466_FSharp4Frequency = A466_G4Frequency / twelfthRoot2
	A466_G4Frequency      = A466_GSharp4Frequency / twelfthRoot2
	A466_GSharp4Frequency = A466_A4Frequency / twelfthRoot2
	A466_A4Frequency      = 466.0
	A466_ASharp4Frequency = A466_A4Frequency * twelfthRoot2
	A466_B4Frequency      = A466_ASharp4Frequency * twelfthRoot2
)

type a466 struct{}

var A466 tuning.Tuning = &a466{}

var a466_scale = [TwelveKeysPerOctave]float64{
	A466_C4Frequency,
	A466_CSharp4Frequency,
	A466_D4Frequency,
	A466_DSharp4Frequency,
	A466_E4Frequency,
	A466_F4Frequency,
	A466_FSharp4Frequency,
	A466_G4Frequency,
	A466_GSharp4Frequency,
	A466_A4Frequency,
	A466_ASharp4Frequency,
	A466_B4Frequency,
}

func (a466) ToFrequency(ko tuning.KeyOctave) float64 {
	k, o := ko.Split(A466)
	freq := a466_scale[k.Index()]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}

func (a466) Key(index int) scale.Key {
	return TwelveKey(index)
}

func (a466) BaseKey() (scale.Key, scale.Octave) {
	return TwelveKeyA, 4
}

func (a466) KeysPerOctave() int {
	return TwelveKeysPerOctave
}
