package equal

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
)

const (
	A435_C4Frequency      = A435_CSharp4Frequency / twelfthRoot2
	A435_CSharp4Frequency = A435_D4Frequency / twelfthRoot2
	A435_D4Frequency      = A435_DSharp4Frequency / twelfthRoot2
	A435_DSharp4Frequency = A435_E4Frequency / twelfthRoot2
	A435_E4Frequency      = A435_F4Frequency / twelfthRoot2
	A435_F4Frequency      = A435_FSharp4Frequency / twelfthRoot2
	A435_FSharp4Frequency = A435_G4Frequency / twelfthRoot2
	A435_G4Frequency      = A435_GSharp4Frequency / twelfthRoot2
	A435_GSharp4Frequency = A435_A4Frequency / twelfthRoot2
	A435_A4Frequency      = 435.0
	A435_ASharp4Frequency = A435_A4Frequency * twelfthRoot2
	A435_B4Frequency      = A435_ASharp4Frequency * twelfthRoot2
)

type a435 struct{}

var A435 tuning.Tuning = &a435{}

var a435_scale = [TwelveKeysPerOctave]float64{
	A435_C4Frequency,
	A435_CSharp4Frequency,
	A435_D4Frequency,
	A435_DSharp4Frequency,
	A435_E4Frequency,
	A435_F4Frequency,
	A435_FSharp4Frequency,
	A435_G4Frequency,
	A435_GSharp4Frequency,
	A435_A4Frequency,
	A435_ASharp4Frequency,
	A435_B4Frequency,
}

func (a435) ToFrequency(ko tuning.KeyOctave) float64 {
	k, o := ko.Split(A435)
	freq := a435_scale[k.Index()]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}

func (a435) Key(index int) scale.Key {
	return TwelveKey(index)
}

func (a435) BaseKey() (scale.Key, scale.Octave) {
	return TwelveKeyA, 4
}

func (a435) KeysPerOctave() int {
	return TwelveKeysPerOctave
}
