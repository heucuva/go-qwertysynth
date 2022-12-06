package just

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	equalTuning "github.com/heucuva/go-qwertysynth/internal/standards/tuning/equal"
)

const (
	Harmonic_C4Frequency      = equalTuning.A440_C4Frequency * HarmonicUnison
	Harmonic_CSharp4Frequency = Harmonic_C4Frequency * HarmonicMinorSecond
	Harmonic_D4Frequency      = Harmonic_C4Frequency * HarmonicMajorSecond
	Harmonic_DSharp4Frequency = Harmonic_C4Frequency * HarmonicMinorThird
	Harmonic_E4Frequency      = Harmonic_C4Frequency * HarmonicMajorThird
	Harmonic_F4Frequency      = Harmonic_C4Frequency * HarmonicFourth
	Harmonic_FSharp4Frequency = Harmonic_C4Frequency * HarmonicDiminishedFourth
	Harmonic_G4Frequency      = Harmonic_C4Frequency * HarmonicFifth
	Harmonic_GSharp4Frequency = Harmonic_C4Frequency * HarmonicMinorSixth
	Harmonic_A4Frequency      = Harmonic_C4Frequency * HarmonicMajorSixth
	Harmonic_ASharp4Frequency = Harmonic_C4Frequency * HarmonicMinorSeventh
	Harmonic_B4Frequency      = Harmonic_C4Frequency * HarmonicMajorSeventh
)

type harmonic struct{}

var Harmonic tuning.Tuning = &harmonic{}

var harmonic_scale = [HarmonicKeysPerOctave]float64{
	Harmonic_C4Frequency,
	Harmonic_CSharp4Frequency,
	Harmonic_D4Frequency,
	Harmonic_DSharp4Frequency,
	Harmonic_E4Frequency,
	Harmonic_F4Frequency,
	Harmonic_FSharp4Frequency,
	Harmonic_G4Frequency,
	Harmonic_GSharp4Frequency,
	Harmonic_A4Frequency,
	Harmonic_ASharp4Frequency,
	Harmonic_B4Frequency,
}

func (harmonic) ToFrequency(ko tuning.KeyOctave) float64 {
	k, o := ko.Split(Harmonic)
	freq := harmonic_scale[k.Index()]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}

func (harmonic) Key(index int) scale.Key {
	return HarmonicKey(index)
}

func (harmonic) BaseKey() (scale.Key, scale.Octave) {
	return HarmonicKeyA, 4
}

func (harmonic) KeysPerOctave() int {
	return HarmonicKeysPerOctave
}
