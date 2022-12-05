package tuning

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
)

type Tuning interface {
	ToFrequency(ko scale.KeyOctave) float64
}
