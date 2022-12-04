package scale

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"
)

const (
	A427_C4Frequency      = A427_CSharp4Frequency / twelfthRoot2
	A427_CSharp4Frequency = A427_D4Frequency / twelfthRoot2
	A427_D4Frequency      = A427_DSharp4Frequency / twelfthRoot2
	A427_DSharp4Frequency = A427_E4Frequency / twelfthRoot2
	A427_E4Frequency      = A427_F4Frequency / twelfthRoot2
	A427_F4Frequency      = A427_FSharp4Frequency / twelfthRoot2
	A427_FSharp4Frequency = A427_G4Frequency / twelfthRoot2
	A427_G4Frequency      = A427_GSharp4Frequency / twelfthRoot2
	A427_GSharp4Frequency = A427_A4Frequency / twelfthRoot2
	A427_A4Frequency      = 427.0
	A427_ASharp4Frequency = A427_A4Frequency * twelfthRoot2
	A427_B4Frequency      = A427_ASharp4Frequency * twelfthRoot2
)

type a427 struct{}

var A427 Scale = &a427{}

var a427_scale = [keyoctave.KeysPerOctave]float64{
	A427_C4Frequency,
	A427_CSharp4Frequency,
	A427_D4Frequency,
	A427_DSharp4Frequency,
	A427_E4Frequency,
	A427_F4Frequency,
	A427_FSharp4Frequency,
	A427_G4Frequency,
	A427_GSharp4Frequency,
	A427_A4Frequency,
	A427_ASharp4Frequency,
	A427_B4Frequency,
}

func (a427) ToFrequency(ko keyoctave.KeyOctave) float64 {
	k, o := ko.Split()
	freq := a427_scale[int(k)]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}
