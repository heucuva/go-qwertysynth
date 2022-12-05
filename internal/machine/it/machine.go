package it

import (
	"github.com/heucuva/go-qwertysynth/internal/machine"
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	"github.com/heucuva/go-qwertysynth/internal/synth/wave"
)

type itMachine struct{}

var Machine = &itMachine{}

func (m itMachine) Default() machine.Default {
	return &m
}

func (itMachine) WaveformFrequency() float64 {
	return itBaseFrequency
}

func (itMachine) Generate(tuning tuning.Tuning, generator wave.Generator, opts ...wave.GeneratorParam) (wave.Wave, error) {
	var machineOpts []wave.GeneratorParam
	machineOpts = append(machineOpts,
		wave.SetParameterByName("frequency", baseNote.ToFrequency(tuning)),
		wave.SetParameterByName("sampleRate", itBaseFrequency),
	)
	return generator(append(machineOpts, opts...)...)
}

func (itMachine) Tuning() tuning.Tuning {
	return defaultTuning
}

const (
	itMicrotonesPerKey       = 64
	itMinOctave              = 0
	itMaxOctave              = 9
	itMicrotonesPerOctave    = itMicrotonesPerKey * scale.KeysPerOctave
	itTotalOctaves           = itMaxOctave - itMinOctave + 1
	itMaxMicrotones          = itMicrotonesPerOctave * itTotalOctaves
	itMinMicrotone           = 0
	itBaseFrequency          = 8363.0
	itOctaveForBaseFrequency = 5
)

func (itMachine) NoteFromChannelData(d uint8) (note.Note, error) {
	switch {
	case d == 0:
		return note.None, nil
	case d >= 1 && d <= 96:
		v := d - 1
		o := itMinOctave + scale.Octave(v/uint8(scale.KeysPerOctave))
		k := scale.MinKey + scale.Key(v%uint8(scale.KeysPerOctave))
		return Machine.Note(o, k, 0), nil
	case d == 97:
		return note.Cut, nil
	default:
		return note.None, ErrChannelDataInvalid
	}
}

func (itMachine) Note(o scale.Octave, k scale.Key, s scale.Microtone) note.Note {
	return itNote(s) +
		itNote(k)*itMicrotonesPerKey +
		itNote(o)*itNote(itMicrotonesPerOctave)
}

func (itMachine) BaseFrequency() float64 {
	return itBaseFrequency
}

func (itMachine) CenterNote() note.Note {
	return Machine.Note(itOctaveForBaseFrequency, scale.KeyC, 0)
}
