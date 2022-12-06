package just

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	equalTuning "github.com/heucuva/go-qwertysynth/internal/standards/tuning/equal"
)

const (
	Pythagorean_DFrequency      = equalTuning.A432_D4Frequency * PythagoreanUnison
	Pythagorean_EFlatFrequency  = Pythagorean_DFrequency * PythagoreanMinorSecond
	Pythagorean_EFrequency      = Pythagorean_DFrequency * PythagoreanMajorSecond
	Pythagorean_FFrequency      = Pythagorean_DFrequency * PythagoreanMinorThird
	Pythagorean_FSharpFrequency = Pythagorean_DFrequency * PythagoreanMajorThird
	Pythagorean_GFrequency      = Pythagorean_DFrequency * PythagoreanPerfectFourth
	Pythagorean_AFlatFrequency  = Pythagorean_DFrequency * PythagoreanDiminishedFifth
	Pythagorean_AFrequency      = Pythagorean_DFrequency * PythagoreanPerfectFifth
	Pythagorean_BFlatFrequency  = Pythagorean_DFrequency * PythagoreanMinorSixth
	Pythagorean_BFrequency      = Pythagorean_DFrequency * PythagoreanMajorSixth
	Pythagorean_CFrequency      = Pythagorean_DFrequency * PythagoreanMinorSeventh
	Pythagorean_CSharpFrequency = Pythagorean_DFrequency * PythagoreanMajorSeventh
)

const (
	PythagoreanUnison          = 1.0
	PythagoreanMinorSecond     = 256.0 / 243.0
	PythagoreanMajorSecond     = 9.0 / 8.0
	PythagoreanMinorThird      = 32.0 / 27.0
	PythagoreanMajorThird      = 81.0 / 64.0
	PythagoreanPerfectFourth   = 4.0 / 3.0
	PythagoreanDiminishedFifth = 1024.0 / 729.0
	PythagoreanPerfectFifth    = 3.0 / 2.0
	PythagoreanMinorSixth      = 128.0 / 81.0
	PythagoreanMajorSixth      = 27.0 / 16.0
	PythagoreanMinorSeventh    = 16.0 / 9.0
	PythagoreanMajorSeventh    = 243.0 / 128.0
)

type pythagorean struct{}

var Pythagorean tuning.Tuning = &pythagorean{}

var pythagorean_scale = [PythagoreanKeysPerOctave]float64{
	Pythagorean_DFrequency,
	Pythagorean_EFlatFrequency,
	Pythagorean_EFrequency,
	Pythagorean_FFrequency,
	Pythagorean_FSharpFrequency,
	Pythagorean_GFrequency,
	Pythagorean_AFlatFrequency,
	Pythagorean_AFrequency,
	Pythagorean_BFlatFrequency,
	Pythagorean_BFrequency,
	Pythagorean_CFrequency,
	Pythagorean_CSharpFrequency,
}

func (pythagorean) ToFrequency(ko tuning.KeyOctave) float64 {
	k, o := ko.Split(Pythagorean)
	freq := pythagorean_scale[k.Index()]
	freq *= math.Pow(2.0, float64(o)-4.0)
	return freq
}

func (pythagorean) Key(index int) scale.Key {
	return PythagoreanKey(index)
}

func (pythagorean) BaseKey() (scale.Key, scale.Octave) {
	return PythagoreanKeyD, 4
}

func (pythagorean) KeysPerOctave() int {
	return PythagoreanKeysPerOctave
}
