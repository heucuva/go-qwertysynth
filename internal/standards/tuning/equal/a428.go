package equal

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
)

const (
	A428_C4Frequency      = A428_CSharp4Frequency / twelfthRoot2
	A428_CSharp4Frequency = A428_D4Frequency / twelfthRoot2
	A428_D4Frequency      = A428_DSharp4Frequency / twelfthRoot2
	A428_DSharp4Frequency = A428_E4Frequency / twelfthRoot2
	A428_E4Frequency      = A428_F4Frequency / twelfthRoot2
	A428_F4Frequency      = A428_FSharp4Frequency / twelfthRoot2
	A428_FSharp4Frequency = A428_G4Frequency / twelfthRoot2
	A428_G4Frequency      = A428_GSharp4Frequency / twelfthRoot2
	A428_GSharp4Frequency = A428_A4Frequency / twelfthRoot2
	A428_A4Frequency      = 428.0
	A428_ASharp4Frequency = A428_A4Frequency * twelfthRoot2
	A428_B4Frequency      = A428_ASharp4Frequency * twelfthRoot2
)

type a428 struct{}

var A428 tuning.Tuning = &a428{}

var a428_scale = [TwelveKeysPerOctave]float64{
	A428_C4Frequency,
	A428_CSharp4Frequency,
	A428_D4Frequency,
	A428_DSharp4Frequency,
	A428_E4Frequency,
	A428_F4Frequency,
	A428_FSharp4Frequency,
	A428_G4Frequency,
	A428_GSharp4Frequency,
	A428_A4Frequency,
	A428_ASharp4Frequency,
	A428_B4Frequency,
}

func (a428) ToFrequency(ko tuning.KeyOctave) float64 {
	k, o := ko.Split(A428)
	freq := a428_scale[k.Index()]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}

func (a428) Key(index int) scale.Key {
	return TwelveKey(index)
}

func (a428) BaseKey() (scale.Key, scale.Octave) {
	return TwelveKeyA, 4
}

func (a428) KeysPerOctave() int {
	return TwelveKeysPerOctave
}
