package scale

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"
)

const (
	Scientific_C4Frequency      = 256.0
	Scientific_CSharp4Frequency = Scientific_C4Frequency * twelfthRoot2
	Scientific_D4Frequency      = Scientific_CSharp4Frequency * twelfthRoot2
	Scientific_DSharp4Frequency = Scientific_D4Frequency * twelfthRoot2
	Scientific_E4Frequency      = Scientific_DSharp4Frequency * twelfthRoot2
	Scientific_F4Frequency      = Scientific_E4Frequency * twelfthRoot2
	Scientific_FSharp4Frequency = Scientific_F4Frequency * twelfthRoot2
	Scientific_G4Frequency      = Scientific_FSharp4Frequency * twelfthRoot2
	Scientific_GSharp4Frequency = Scientific_G4Frequency * twelfthRoot2
	Scientific_A4Frequency      = Scientific_GSharp4Frequency * twelfthRoot2
	Scientific_ASharp4Frequency = Scientific_A4Frequency * twelfthRoot2
	Scientific_B4Frequency      = Scientific_ASharp4Frequency * twelfthRoot2
)

type scientific struct{}

var Scientific Scale = &scientific{}

var scientific_scale = [keyoctave.KeysPerOctave]float64{
	Scientific_C4Frequency,
	Scientific_CSharp4Frequency,
	Scientific_D4Frequency,
	Scientific_DSharp4Frequency,
	Scientific_E4Frequency,
	Scientific_F4Frequency,
	Scientific_FSharp4Frequency,
	Scientific_G4Frequency,
	Scientific_GSharp4Frequency,
	Scientific_A4Frequency,
	Scientific_ASharp4Frequency,
	Scientific_B4Frequency,
}

func (scientific) ToFrequency(ko keyoctave.KeyOctave) float64 {
	k, o := ko.Split()
	freq := scientific_scale[int(k)]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}
