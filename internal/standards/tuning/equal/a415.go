package equal

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
)

const (
	A415_C4Frequency      = A415_CSharp4Frequency / twelfthRoot2
	A415_CSharp4Frequency = A415_D4Frequency / twelfthRoot2
	A415_D4Frequency      = A415_DSharp4Frequency / twelfthRoot2
	A415_DSharp4Frequency = A415_E4Frequency / twelfthRoot2
	A415_E4Frequency      = A415_F4Frequency / twelfthRoot2
	A415_F4Frequency      = A415_FSharp4Frequency / twelfthRoot2
	A415_FSharp4Frequency = A415_G4Frequency / twelfthRoot2
	A415_G4Frequency      = A415_GSharp4Frequency / twelfthRoot2
	A415_GSharp4Frequency = A415_A4Frequency / twelfthRoot2
	A415_A4Frequency      = 415.0
	A415_ASharp4Frequency = A415_A4Frequency * twelfthRoot2
	A415_B4Frequency      = A415_ASharp4Frequency * twelfthRoot2
)

type a415 struct{}

var A415 tuning.Tuning = &a415{}

var a415_scale = [scale.KeysPerOctave]float64{
	A415_C4Frequency,
	A415_CSharp4Frequency,
	A415_D4Frequency,
	A415_DSharp4Frequency,
	A415_E4Frequency,
	A415_F4Frequency,
	A415_FSharp4Frequency,
	A415_G4Frequency,
	A415_GSharp4Frequency,
	A415_A4Frequency,
	A415_ASharp4Frequency,
	A415_B4Frequency,
}

func (a415) ToFrequency(ko scale.KeyOctave) float64 {
	k, o := ko.Split()
	freq := a415_scale[int(k)]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}
