package xm

import (
	"fmt"
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	tuningPkg "github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	equalTuning "github.com/heucuva/go-qwertysynth/internal/standards/tuning/equal"
)

type xmNote struct {
	oks    int
	tuning tuningPkg.Tuning
}

func (n xmNote) Split() (scale.Octave, scale.Key, scale.Microtone) {
	o, k, s := n.split()
	return scale.Octave(o), n.tuning.Key(k), scale.Microtone(s)
}

func (n xmNote) split() (int, int, int) {
	keysPerOctave := n.tuning.KeysPerOctave()
	ko := int(n.oks) / xmMicrotonesPerKey
	o := ko / keysPerOctave
	k := ko % keysPerOctave
	s := int(n.oks) % xmMicrotonesPerKey
	return o, k, s
}

func (n xmNote) KeyOctave() tuningPkg.KeyOctave {
	o, k, _ := n.split()
	ko := tuningPkg.KeyOctave(k) |
		tuning.KeyOctave(o)<<8
	return ko
}

func (xmNote) IsCut() bool {
	return false
}

func (xmNote) IsFadeout() bool {
	return false
}

func (n xmNote) Microtones() scale.Microtone {
	return scale.Microtone(n.oks)
}

func (xmNote) Kind() note.Kind {
	return Machine(defaultTuning)
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
)

func (n xmNote) ToFrequency() float64 {
	_, _, m := n.split()
	frequency := n.tuning.ToFrequency(n.KeyOctave()) *
		math.Pow(2.0, float64(m)/microtonesPerOctave)
	return frequency
}

func (n xmNote) AddMicrotones(s scale.Microtone) note.Note {
	_, _, st := n.split()
	n.oks -= st + int(s)
	if n.oks < 0 {
		n.oks = 0
	}
	return n
}
