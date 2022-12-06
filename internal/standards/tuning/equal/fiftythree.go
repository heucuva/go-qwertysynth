package equal

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
)

const (
	FiftyThree_C4Frequency                       = FiftyThree_CSharp4Frequency / fiftyThirdRoot2
	FiftyThree_CSharp4Frequency                  = FiftyThree_CDoubleSharp4Frequency / fiftyThirdRoot2
	FiftyThree_CDoubleSharp4Frequency            = FiftyThree_CSharpDoubleSharp4Frequency / fiftyThirdRoot2
	FiftyThree_CSharpDoubleSharp4Frequency       = FiftyThree_CDoubleSharpDoubleSharp4Frequency / fiftyThirdRoot2
	FiftyThree_CDoubleSharpDoubleSharp4Frequency = FiftyThree_DDoubleFlatDoubleFlat4Frequency / fiftyThirdRoot2
	FiftyThree_DDoubleFlatDoubleFlat4Frequency   = FiftyThree_DTripleFlat4Frequency / fiftyThirdRoot2
	FiftyThree_DTripleFlat4Frequency             = FiftyThree_DDoubleFlat4Frequency / fiftyThirdRoot2
	FiftyThree_DDoubleFlat4Frequency             = FiftyThree_DFlat / fiftyThirdRoot2
	FiftyThree_DFlat                             = FiftyThree_D4Frequency / fiftyThirdRoot2
	FiftyThree_D4Frequency                       = FiftyThree_DSharp4Frequency / fiftyThirdRoot2
	FiftyThree_DSharp4Frequency                  = FiftyThree_DDoubleSharp4Frequency / fiftyThirdRoot2
	FiftyThree_DDoubleSharp4Frequency            = FiftyThree_DSharpDoubleSharp4Frequency / fiftyThirdRoot2
	FiftyThree_DSharpDoubleSharp4Frequency       = FiftyThree_DDoubleSharpDoubleSharp4Frequency / fiftyThirdRoot2
	FiftyThree_DDoubleSharpDoubleSharp4Frequency = FiftyThree_EDoubleFlatDoubleFlat4Frequency / fiftyThirdRoot2
	FiftyThree_EDoubleFlatDoubleFlat4Frequency   = FiftyThree_ETripleFlat4Frequency / fiftyThirdRoot2
	FiftyThree_ETripleFlat4Frequency             = FiftyThree_EDoubleFlat4Frequency / fiftyThirdRoot2
	FiftyThree_EDoubleFlat4Frequency             = FiftyThree_EFlat4Frequency / fiftyThirdRoot2
	FiftyThree_EFlat4Frequency                   = FiftyThree_E4Frequency / fiftyThirdRoot2
	FiftyThree_E4Frequency                       = FiftyThree_ESharp4Frequency / fiftyThirdRoot2
	FiftyThree_ESharp4Frequency                  = FiftyThree_EDoubleSharp4Frequency / fiftyThirdRoot2
	FiftyThree_EDoubleSharp4Frequency            = FiftyThree_FFlat4Frequency / fiftyThirdRoot2
	FiftyThree_FFlat4Frequency                   = FiftyThree_F4Frequency / fiftyThirdRoot2
	FiftyThree_F4Frequency                       = FiftyThree_FSharp4Frequency / fiftyThirdRoot2
	FiftyThree_FSharp4Frequency                  = FiftyThree_FDoubleSharp4Frequency / fiftyThirdRoot2
	FiftyThree_FDoubleSharp4Frequency            = FiftyThree_FSharpDoubleSharp4Frequency / fiftyThirdRoot2
	FiftyThree_FSharpDoubleSharp4Frequency       = FiftyThree_FDoubleSharpDoubleSharp4Frequency / fiftyThirdRoot2
	FiftyThree_FDoubleSharpDoubleSharp4Frequency = FiftyThree_GDoubleFlatDoubleFlat4Frequency / fiftyThirdRoot2
	FiftyThree_GDoubleFlatDoubleFlat4Frequency   = FiftyThree_GTripleFlat4Frequency / fiftyThirdRoot2
	FiftyThree_GTripleFlat4Frequency             = FiftyThree_GDoubleFlat4Frequency / fiftyThirdRoot2
	FiftyThree_GDoubleFlat4Frequency             = FiftyThree_GFlat4Frequency / fiftyThirdRoot2
	FiftyThree_GFlat4Frequency                   = FiftyThree_G4Frequency / fiftyThirdRoot2
	FiftyThree_G4Frequency                       = FiftyThree_GSharp4Frequency / fiftyThirdRoot2
	FiftyThree_GSharp4Frequency                  = FiftyThree_GDoubleSharp4Frequency / fiftyThirdRoot2
	FiftyThree_GDoubleSharp4Frequency            = FiftyThree_GSharpDoubleSharp4Frequency / fiftyThirdRoot2
	FiftyThree_GSharpDoubleSharp4Frequency       = FiftyThree_GDoubleSharpDoubleSharp4Frequency / fiftyThirdRoot2
	FiftyThree_GDoubleSharpDoubleSharp4Frequency = FiftyThree_ADoubleFlatDoubleFlat4Frequency / fiftyThirdRoot2
	FiftyThree_ADoubleFlatDoubleFlat4Frequency   = FiftyThree_ATripleFlat4Frequency / fiftyThirdRoot2
	FiftyThree_ATripleFlat4Frequency             = FiftyThree_ADoubleFlat4Frequency / fiftyThirdRoot2
	FiftyThree_ADoubleFlat4Frequency             = FiftyThree_AFlat4Frequency / fiftyThirdRoot2
	FiftyThree_AFlat4Frequency                   = FiftyThree_A4Frequency / fiftyThirdRoot2
	FiftyThree_A4Frequency                       = 440.0
	FiftyThree_ASharp4Frequency                  = FiftyThree_A4Frequency * fiftyThirdRoot2
	FiftyThree_ADoubleSharp4Frequency            = FiftyThree_ASharp4Frequency * fiftyThirdRoot2
	FiftyThree_ASharpDoubleSharp4Frequency       = FiftyThree_ADoubleSharp4Frequency * fiftyThirdRoot2
	FiftyThree_ADoubleSharpDoubleSharp4Frequency = FiftyThree_ASharpDoubleSharp4Frequency * fiftyThirdRoot2
	FiftyThree_BDoubleFlatDoubleFlat4Frequency   = FiftyThree_ADoubleSharpDoubleSharp4Frequency * fiftyThirdRoot2
	FiftyThree_BTripleFlat4Frequency             = FiftyThree_BDoubleFlatDoubleFlat4Frequency * fiftyThirdRoot2
	FiftyThree_BDoubleFlat4Frequency             = FiftyThree_BTripleFlat4Frequency * fiftyThirdRoot2
	FiftyThree_BFlat4Frequency                   = FiftyThree_BDoubleFlat4Frequency * fiftyThirdRoot2
	FiftyThree_B4Frequency                       = FiftyThree_BFlat4Frequency * fiftyThirdRoot2
	FiftyThree_BSharp4Frequency                  = FiftyThree_B4Frequency * fiftyThirdRoot2
	FiftyThree_BDoubleSharp4Frequency            = FiftyThree_BSharp4Frequency * fiftyThirdRoot2
	FiftyThree_CFlat4Frequency                   = FiftyThree_BDoubleSharp4Frequency * fiftyThirdRoot2
)

type fiftyThree struct{}

var FiftyThree tuning.Tuning = &fiftyThree{}

var fiftyThree_scale = [FiftyThreeKeysPerOctave]float64{
	FiftyThree_C4Frequency,
	FiftyThree_CSharp4Frequency,
	FiftyThree_CDoubleSharp4Frequency,
	FiftyThree_CSharpDoubleSharp4Frequency,
	FiftyThree_CDoubleSharpDoubleSharp4Frequency,
	FiftyThree_DDoubleFlatDoubleFlat4Frequency,
	FiftyThree_DTripleFlat4Frequency,
	FiftyThree_DDoubleFlat4Frequency,
	FiftyThree_DFlat,
	FiftyThree_D4Frequency,
	FiftyThree_DSharp4Frequency,
	FiftyThree_DDoubleSharp4Frequency,
	FiftyThree_DSharpDoubleSharp4Frequency,
	FiftyThree_DDoubleSharpDoubleSharp4Frequency,
	FiftyThree_EDoubleFlatDoubleFlat4Frequency,
	FiftyThree_ETripleFlat4Frequency,
	FiftyThree_EDoubleFlat4Frequency,
	FiftyThree_EFlat4Frequency,
	FiftyThree_E4Frequency,
	FiftyThree_ESharp4Frequency,
	FiftyThree_EDoubleSharp4Frequency,
	FiftyThree_FFlat4Frequency,
	FiftyThree_F4Frequency,
	FiftyThree_FSharp4Frequency,
	FiftyThree_FDoubleSharp4Frequency,
	FiftyThree_FSharpDoubleSharp4Frequency,
	FiftyThree_FDoubleSharpDoubleSharp4Frequency,
	FiftyThree_GDoubleFlatDoubleFlat4Frequency,
	FiftyThree_GTripleFlat4Frequency,
	FiftyThree_GDoubleFlat4Frequency,
	FiftyThree_GFlat4Frequency,
	FiftyThree_G4Frequency,
	FiftyThree_GSharp4Frequency,
	FiftyThree_GDoubleSharp4Frequency,
	FiftyThree_GSharpDoubleSharp4Frequency,
	FiftyThree_GDoubleSharpDoubleSharp4Frequency,
	FiftyThree_ADoubleFlatDoubleFlat4Frequency,
	FiftyThree_ATripleFlat4Frequency,
	FiftyThree_ADoubleFlat4Frequency,
	FiftyThree_AFlat4Frequency,
	FiftyThree_A4Frequency,
	FiftyThree_ASharp4Frequency,
	FiftyThree_ADoubleSharp4Frequency,
	FiftyThree_ASharpDoubleSharp4Frequency,
	FiftyThree_ADoubleSharpDoubleSharp4Frequency,
	FiftyThree_BDoubleFlatDoubleFlat4Frequency,
	FiftyThree_BTripleFlat4Frequency,
	FiftyThree_BDoubleFlat4Frequency,
	FiftyThree_BFlat4Frequency,
	FiftyThree_B4Frequency,
	FiftyThree_BSharp4Frequency,
	FiftyThree_BDoubleSharp4Frequency,
	FiftyThree_CFlat4Frequency,
}

func (fiftyThree) ToFrequency(ko tuning.KeyOctave) float64 {
	k, o := ko.Split(FiftyThree)
	freq := fiftyThree_scale[k.Index()]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}

func (fiftyThree) Key(index int) scale.Key {
	return FiftyThreeKey(index)
}

func (fiftyThree) BaseKey() (scale.Key, scale.Octave) {
	return FiftyThreeKeyA, 4
}

func (fiftyThree) KeysPerOctave() int {
	return FiftyThreeKeysPerOctave
}
