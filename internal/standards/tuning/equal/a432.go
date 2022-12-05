package equal

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
)

const (
	A432_C4Frequency      = A432_CSharp4Frequency / twelfthRoot2
	A432_CSharp4Frequency = A432_D4Frequency / twelfthRoot2
	A432_D4Frequency      = A432_DSharp4Frequency / twelfthRoot2
	A432_DSharp4Frequency = A432_E4Frequency / twelfthRoot2
	A432_E4Frequency      = A432_F4Frequency / twelfthRoot2
	A432_F4Frequency      = A432_FSharp4Frequency / twelfthRoot2
	A432_FSharp4Frequency = A432_G4Frequency / twelfthRoot2
	A432_G4Frequency      = A432_GSharp4Frequency / twelfthRoot2
	A432_GSharp4Frequency = A432_A4Frequency / twelfthRoot2
	A432_A4Frequency      = 432.0
	A432_ASharp4Frequency = A432_A4Frequency * twelfthRoot2
	A432_B4Frequency      = A432_ASharp4Frequency * twelfthRoot2
)

type a432 struct{}

var A432 tuning.Tuning = &a432{}

var a432_scale = [scale.KeysPerOctave]float64{
	A432_C4Frequency,
	A432_CSharp4Frequency,
	A432_D4Frequency,
	A432_DSharp4Frequency,
	A432_E4Frequency,
	A432_F4Frequency,
	A432_FSharp4Frequency,
	A432_G4Frequency,
	A432_GSharp4Frequency,
	A432_A4Frequency,
	A432_ASharp4Frequency,
	A432_B4Frequency,
}

func (a432) ToFrequency(ko scale.KeyOctave) float64 {
	k, o := ko.Split()
	freq := a432_scale[int(k)]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}
