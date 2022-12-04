package it

import (
	"github.com/heucuva/go-qwertysynth/internal/machine"
	"github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
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

func (itMachine) Generate(generator wave.Generator, opts ...wave.GeneratorParam) (wave.Wave, error) {
	var machineOpts []wave.GeneratorParam
	machineOpts = append(machineOpts,
		wave.SetParameterByName("frequency", c5Freq),
		wave.SetParameterByName("sampleRate", itBaseFrequency),
	)
	return generator(append(machineOpts, opts...)...)
}

const (
	itSemitonesPerKey        = 64
	itMinOctave              = 0
	itMaxOctave              = 9
	itSemitonesPerOctave     = itSemitonesPerKey * keyoctave.KeysPerOctave
	itTotalOctaves           = itMaxOctave - itMinOctave + 1
	itMaxSemitones           = itSemitonesPerOctave * itTotalOctaves
	itMinSemitone            = 0
	itBaseFrequency          = 8363.0
	itOctaveForBaseFrequency = 5
)

func (itMachine) NoteFromChannelData(d uint8) (note.Note, error) {
	switch {
	case d == 0:
		return note.None, nil
	case d >= 1 && d <= 96:
		v := d - 1
		o := itMinOctave + keyoctave.Octave(v/uint8(keyoctave.KeysPerOctave))
		k := keyoctave.MinKey + keyoctave.Key(v%uint8(keyoctave.KeysPerOctave))
		return Machine.Note(o, k, 0), nil
	case d == 97:
		return note.Cut, nil
	default:
		return note.None, ErrChannelDataInvalid
	}
}

func (itMachine) Note(o keyoctave.Octave, k keyoctave.Key, s keyoctave.Semitone) note.Note {
	return itNote(s) +
		itNote(k)*itSemitonesPerKey +
		itNote(o)*itNote(itSemitonesPerOctave)
}

func (itMachine) BaseFrequency() float64 {
	return itBaseFrequency
}

func (itMachine) CenterNote() note.Note {
	return Machine.Note(itOctaveForBaseFrequency, keyoctave.KeyC, 0)
}
