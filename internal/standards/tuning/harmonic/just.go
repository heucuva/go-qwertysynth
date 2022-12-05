package harmonic

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	equalTuning "github.com/heucuva/go-qwertysynth/internal/standards/tuning/equal"
)

const (
	Just_C4Frequency      = equalTuning.A440_C4Frequency * Unison
	Just_CSharp4Frequency = Just_C4Frequency * MinorSecond
	Just_D4Frequency      = Just_C4Frequency * MajorSecond
	Just_DSharp4Frequency = Just_C4Frequency * MinorThird
	Just_E4Frequency      = Just_C4Frequency * MajorThird
	Just_F4Frequency      = Just_C4Frequency * Fourth
	Just_FSharp4Frequency = Just_C4Frequency * DiminishedFourth
	Just_G4Frequency      = Just_C4Frequency * Fifth
	Just_GSharp4Frequency = Just_C4Frequency * MinorSixth
	Just_A4Frequency      = Just_C4Frequency * MajorSixth
	Just_ASharp4Frequency = Just_C4Frequency * MinorSeventh
	Just_B4Frequency      = Just_C4Frequency * MajorSeventh
)

type just struct{}

var Just tuning.Tuning = &just{}

var just_scale = [scale.KeysPerOctave]float64{
	Just_C4Frequency,
	Just_CSharp4Frequency,
	Just_D4Frequency,
	Just_DSharp4Frequency,
	Just_E4Frequency,
	Just_F4Frequency,
	Just_FSharp4Frequency,
	Just_G4Frequency,
	Just_GSharp4Frequency,
	Just_A4Frequency,
	Just_ASharp4Frequency,
	Just_B4Frequency,
}

func (just) ToFrequency(ko scale.KeyOctave) float64 {
	k, o := ko.Split()
	freq := just_scale[int(k)]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}
