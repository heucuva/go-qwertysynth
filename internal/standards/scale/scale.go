package scale

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"
)

type Scale interface {
	ToFrequency(ko keyoctave.KeyOctave) float64
}
