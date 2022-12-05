package xm

import (
	"fmt"
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	equalTuning "github.com/heucuva/go-qwertysynth/internal/standards/tuning/equal"
)

type xmNote int

func (n xmNote) Split() (scale.Octave, scale.Key, scale.Microtone) {
	ko := n / xmMicrotonesPerKey
	o := scale.Octave(ko) / scale.Octave(scale.KeysPerOctave)
	k := scale.Key(ko) % scale.Key(scale.KeysPerOctave)
	s := scale.Microtone(n) % xmMicrotonesPerKey
	return o, k, s
}

func (n xmNote) KeyOctave() scale.KeyOctave {
	return scale.KeyOctave(n) / xmMicrotonesPerKey
}

func (xmNote) IsCut() bool {
	return false
}

func (xmNote) IsFadeout() bool {
	return false
}

func (n xmNote) Microtones() scale.Microtone {
	return scale.Microtone(n)
}

func (xmNote) Kind() note.Kind {
	return Machine
}

func (n xmNote) String() string {
	o, k, s := n.Split()
	if s != 0 {
		return fmt.Sprintf("%v%d (%d)", k, o, s)
	}

	return fmt.Sprintf("%v%d", k, o)
}

const (
	microtonesPerKey    float64 = xmMicrotonesPerKey
	microtonesPerOctave float64 = float64(xmMicrotonesPerOctave)
	totalOctaves        float64 = float64(scale.NumOctaves)
	totalMicrotones     float64 = totalOctaves * microtonesPerOctave
	divisorOctave       float64 = totalOctaves - xmOctaveForBaseFrequency
	divisorMicrotone    float64 = divisorOctave * microtonesPerOctave
)

var (
	defaultTuning = equalTuning.A440
	baseNote      = Machine.Note(xmOctaveForBaseFrequency, scale.KeyC, 0)
)

func (n xmNote) ToFrequency(tuning tuning.Tuning) float64 {
	if tuning == nil {
		tuning = defaultTuning
	}

	o, k, m := n.Split()
	frequency := tuning.ToFrequency(scale.NewKeyOctave(k, o)) *
		math.Pow(2.0, float64(m)/microtonesPerOctave)
	return frequency
}

func (n xmNote) AddMicrotones(s scale.Microtone) note.Note {
	o, k, st := n.Split()
	return xmNote(s+st) +
		xmNote(k)*xmMicrotonesPerKey +
		xmNote(o)*xmNote(xmMicrotonesPerOctave)
}
