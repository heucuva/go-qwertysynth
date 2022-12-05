package it

import (
	"fmt"
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	equalTuning "github.com/heucuva/go-qwertysynth/internal/standards/tuning/equal"
)

type itNote int

func (n itNote) Split() (scale.Octave, scale.Key, scale.Microtone) {
	ko := n / itMicrotonesPerKey
	o := scale.Octave(ko) / scale.Octave(scale.KeysPerOctave)
	k := scale.Key(ko) % scale.Key(scale.KeysPerOctave)
	s := scale.Microtone(n) % itMicrotonesPerKey
	return o, k, s
}

func (n itNote) KeyOctave() scale.KeyOctave {
	return scale.KeyOctave(n) / itMicrotonesPerKey
}

func (itNote) IsCut() bool {
	return false
}

func (itNote) IsFadeout() bool {
	return false
}

func (n itNote) Microtones() scale.Microtone {
	return scale.Microtone(n)
}

func (itNote) Kind() note.Kind {
	return Machine
}

func (n itNote) String() string {
	o, k, s := n.Split()
	if s != 0 {
		return fmt.Sprintf("%v%d (%d)", k, o, s)
	}

	return fmt.Sprintf("%v%d", k, o)
}

const (
	microtonesPerKey    float64 = itMicrotonesPerKey
	microtonesPerOctave float64 = float64(itMicrotonesPerOctave)
	totalOctaves        float64 = float64(scale.NumOctaves)
	totalMicrotones     float64 = totalOctaves * microtonesPerOctave
	divisorOctave       float64 = totalOctaves - itOctaveForBaseFrequency
	divisorMicrotone    float64 = divisorOctave * microtonesPerOctave
)

var (
	defaultTuning = equalTuning.A440
	baseNote      = Machine.Note(itOctaveForBaseFrequency, scale.KeyC, 0)
)

func (n itNote) ToFrequency(tuning tuning.Tuning) float64 {
	if tuning == nil {
		tuning = defaultTuning
	}

	o, k, m := n.Split()
	frequency := tuning.ToFrequency(scale.NewKeyOctave(k, o-1)) *
		math.Pow(2.0, float64(m)/microtonesPerOctave)
	return frequency
}

func (n itNote) AddMicrotones(s scale.Microtone) note.Note {
	o, k, st := n.Split()
	return itNote(s+st) +
		itNote(k)*itMicrotonesPerKey +
		itNote(o)*itNote(itMicrotonesPerOctave)
}
