package xm

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
)

type xmMachineDefault struct{}

func (xmMachineDefault) WaveformFrequency() float64 {
	return xmBaseFrequency
}

func (xmMachineDefault) Tuning() tuning.Tuning {
	return defaultTuning
}

func (m xmMachineDefault) BaseNote() note.Note {
	tuning := m.Tuning()
	return Machine(tuning).Note(xmOctaveForBaseFrequency, tuning.Key(0), 0)
}

var xmDefault = &xmMachineDefault{}
