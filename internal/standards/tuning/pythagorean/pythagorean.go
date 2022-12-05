package pythagorean

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	equalTuning "github.com/heucuva/go-qwertysynth/internal/standards/tuning/equal"
)

const (
	Pythagorean_C4Frequency      = Pythagorean_D4Frequency * MinorSeventh
	Pythagorean_CSharp4Frequency = Pythagorean_D4Frequency * MajorSeventh
	Pythagorean_D4Frequency      = equalTuning.A432_D4Frequency * Unison
	Pythagorean_DSharp4Frequency = Pythagorean_D4Frequency * MinorSecond
	Pythagorean_E4Frequency      = Pythagorean_D4Frequency * MajorSecond
	Pythagorean_F4Frequency      = Pythagorean_D4Frequency * MinorThird
	Pythagorean_FSharp4Frequency = Pythagorean_D4Frequency * MajorThird
	Pythagorean_G4Frequency      = Pythagorean_D4Frequency * PerfectFourth
	Pythagorean_GSharp4Frequency = Pythagorean_D4Frequency * AugmentedFourth
	Pythagorean_A4Frequency      = Pythagorean_D4Frequency * PerfectFifth
	Pythagorean_ASharp4Frequency = Pythagorean_D4Frequency * MinorSixth
	Pythagorean_B4Frequency      = Pythagorean_D4Frequency * MajorSixth
)

type pythagorean struct{}

var Pythagorean tuning.Tuning = &pythagorean{}

var pythagorean_scale = [scale.KeysPerOctave]float64{
	Pythagorean_C4Frequency,
	Pythagorean_CSharp4Frequency,
	Pythagorean_D4Frequency,
	Pythagorean_DSharp4Frequency,
	Pythagorean_E4Frequency,
	Pythagorean_F4Frequency,
	Pythagorean_FSharp4Frequency,
	Pythagorean_G4Frequency,
	Pythagorean_GSharp4Frequency,
	Pythagorean_A4Frequency,
	Pythagorean_ASharp4Frequency,
	Pythagorean_B4Frequency,
}

func (pythagorean) ToFrequency(ko scale.KeyOctave) float64 {
	k, o := ko.Split()
	freq := pythagorean_scale[int(k)]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}
