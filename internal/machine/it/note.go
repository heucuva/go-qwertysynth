package it

import (
	"fmt"
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	tuningPkg "github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	equalTuning "github.com/heucuva/go-qwertysynth/internal/standards/tuning/equal"
)

type itNote struct {
	oks    int
	tuning tuningPkg.Tuning
}

func (n itNote) Split() (scale.Octave, scale.Key, scale.Microtone) {
	o, k, s := n.split()
	return scale.Octave(o), n.tuning.Key(k), scale.Microtone(s)
}

func (n itNote) split() (int, int, int) {
	keysPerOctave := n.tuning.KeysPerOctave()
	ko := int(n.oks) / itMicrotonesPerKey
	o := ko / keysPerOctave
	k := ko % keysPerOctave
	s := int(n.oks) % itMicrotonesPerKey
	return o, k, s
}

func (n itNote) KeyOctave() tuningPkg.KeyOctave {
	o, k, _ := n.split()
	ko := tuningPkg.KeyOctave(k) |
		tuning.KeyOctave(o)<<8
	return ko
}

func (itNote) IsCut() bool {
	return false
}

func (itNote) IsFadeout() bool {
	return false
}

func (n itNote) Microtones() scale.Microtone {
	return scale.Microtone(n.oks)
}

func (itNote) Kind() note.Kind {
	return Machine(defaultTuning)
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
)

func (n itNote) ToFrequency() float64 {
	_, _, m := n.split()
	frequency := n.tuning.ToFrequency(n.KeyOctave()) *
		math.Pow(2.0, float64(m)/microtonesPerOctave)
	return frequency
}

func (n itNote) AddMicrotones(s scale.Microtone) note.Note {
	_, _, st := n.split()
	n.oks -= st + int(s)
	if n.oks < 0 {
		n.oks = 0
	}
	return n
}
