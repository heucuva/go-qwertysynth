package it

import (
	"fmt"
	"math"

	"github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
)

type itNote int

func (n itNote) Split() (keyoctave.Octave, keyoctave.Key, keyoctave.Semitone) {
	ko := n / itSemitonesPerKey
	o := keyoctave.Octave(ko) / keyoctave.Octave(keyoctave.KeysPerOctave)
	k := keyoctave.Key(ko) % keyoctave.Key(keyoctave.KeysPerOctave)
	s := keyoctave.Semitone(n) % itSemitonesPerKey
	return o, k, s
}

func (n itNote) KeyOctave() keyoctave.KeyOctave {
	return keyoctave.KeyOctave(n) / itSemitonesPerKey
}

func (itNote) IsCut() bool {
	return false
}

func (itNote) IsFadeout() bool {
	return false
}

func (n itNote) Semitones() keyoctave.Semitone {
	return keyoctave.Semitone(n)
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
	semitonesPerKey    float64 = itSemitonesPerKey
	semitonesPerOctave float64 = float64(itSemitonesPerOctave)
	totalOctaves       float64 = float64(keyoctave.NumOctaves)
	totalSemitones     float64 = totalOctaves * semitonesPerOctave
	divisorOctave      float64 = totalOctaves - itOctaveForBaseFrequency
	divisorSemitone    float64 = divisorOctave * semitonesPerOctave
	c5Freq             float64 = scale.A440_C4Frequency
)

func (n itNote) ToFrequency() float64 {
	semitones := float64(n)
	period := totalSemitones - semitones
	frequency := c5Freq * math.Pow(2.0, (divisorSemitone-period)/semitonesPerOctave)
	return frequency
}

func (n itNote) AddSemitones(s keyoctave.Semitone) note.Note {
	o, k, st := n.Split()
	return itNote(s+st) +
		itNote(k)*itSemitonesPerKey +
		itNote(o)*itNote(itSemitonesPerOctave)
}
