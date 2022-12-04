package scale

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"
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

var A466 Scale = &a466{}

var a466_scale = [keyoctave.KeysPerOctave]float64{
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

func (a466) ToFrequency(ko keyoctave.KeyOctave) float64 {
	k, o := ko.Split()
	freq := a466_scale[int(k)]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}
