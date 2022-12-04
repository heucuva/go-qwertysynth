package xm

import (
	"fmt"
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
)

type xmNote int

func (n xmNote) Split() (keyoctave.Octave, keyoctave.Key, keyoctave.Semitone) {
	ko := n / xmSemitonesPerKey
	o := keyoctave.Octave(ko) / keyoctave.Octave(keyoctave.KeysPerOctave)
	k := keyoctave.Key(ko) % keyoctave.Key(keyoctave.KeysPerOctave)
	s := keyoctave.Semitone(n) % xmSemitonesPerKey
	return o, k, s
}

func (n xmNote) KeyOctave() keyoctave.KeyOctave {
	return keyoctave.KeyOctave(n) / xmSemitonesPerKey
}

func (xmNote) IsCut() bool {
	return false
}

func (xmNote) IsFadeout() bool {
	return false
}

func (n xmNote) Semitones() keyoctave.Semitone {
	return keyoctave.Semitone(n)
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
	semitonesPerKey    float64 = xmSemitonesPerKey
	semitonesPerOctave float64 = float64(xmSemitonesPerOctave)
	totalOctaves       float64 = float64(keyoctave.NumOctaves)
	totalSemitones     float64 = totalOctaves * semitonesPerOctave
	divisorOctave      float64 = totalOctaves - xmOctaveForBaseFrequency
	divisorSemitone    float64 = divisorOctave * semitonesPerOctave
	c4Freq             float64 = scale.A440_C4Frequency
)

func (n xmNote) ToFrequency() float64 {
	semitones := float64(n)
	period := totalSemitones - semitones
	frequency := c4Freq * math.Pow(2.0, (divisorSemitone-period)/semitonesPerOctave)
	return frequency
}

func (n xmNote) AddSemitones(s keyoctave.Semitone) note.Note {
	o, k, st := n.Split()
	return xmNote(s+st) +
		xmNote(k)*xmSemitonesPerKey +
		xmNote(o)*xmNote(xmSemitonesPerOctave)
}
