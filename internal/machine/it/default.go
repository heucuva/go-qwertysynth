package it

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
)

type itMachineDefault struct{}

func (itMachineDefault) WaveformFrequency() float64 {
	return itBaseFrequency
}

func (itMachineDefault) Tuning() tuning.Tuning {
	return defaultTuning
}

func (m itMachineDefault) BaseNote() note.Note {
	tuning := m.Tuning()
	return Machine(tuning).Note(itOctaveForBaseFrequency, tuning.Key(0), 0)
}

var itDefault = &itMachineDefault{}
