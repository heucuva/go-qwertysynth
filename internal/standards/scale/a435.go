package scale

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"
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

var A435 Scale = &a435{}

var a435_scale = [keyoctave.KeysPerOctave]float64{
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

func (a435) ToFrequency(ko keyoctave.KeyOctave) float64 {
	k, o := ko.Split()
	freq := a435_scale[int(k)]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}
