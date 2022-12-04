package xm

import (
	"github.com/heucuva/go-qwertysynth/internal/machine"
	"github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/synth/wave"
)

type xmMachine struct{}

var Machine = &xmMachine{}

func (m xmMachine) Default() machine.Default {
	return &m
}

func (xmMachine) WaveformFrequency() float64 {
	return xmBaseFrequency
}

func (xmMachine) Generate(generator wave.Generator, opts ...wave.GeneratorParam) (wave.Wave, error) {
	var machineOpts []wave.GeneratorParam
	machineOpts = append(machineOpts,
		wave.SetParameterByName("frequency", c4Freq),
		wave.SetParameterByName("sampleRate", xmBaseFrequency),
	)
	return generator(append(machineOpts, opts...)...)
}

const (
	xmSemitonesPerKey        = 64
	xmMinOctave              = 1
	xmMaxOctave              = 8
	xmSemitonesPerOctave     = xmSemitonesPerKey * keyoctave.KeysPerOctave
	xmTotalOctaves           = xmMaxOctave - xmMinOctave + 1
	xmMaxSemitones           = xmSemitonesPerOctave * xmTotalOctaves
	xmMinSemitone            = 0
	xmBaseFrequency          = 8363.0
	xmOctaveForBaseFrequency = 4
)

func (xmMachine) NoteFromChannelData(d uint8) (note.Note, error) {
	switch {
	case d == 0:
		return note.None, nil
	case d >= 1 && d <= 96:
		v := d - 1
		o := xmMinOctave + keyoctave.Octave(v/uint8(keyoctave.KeysPerOctave))
		k := keyoctave.MinKey + keyoctave.Key(v%uint8(keyoctave.KeysPerOctave))
		return Machine.Note(o, k, 0), nil
	case d == 97:
		return note.Cut, nil
	default:
		return note.None, ErrChannelDataInvalid
	}
}

func (xmMachine) Note(o keyoctave.Octave, k keyoctave.Key, s keyoctave.Semitone) note.Note {
	return xmNote(s) +
		xmNote(k)*xmSemitonesPerKey +
		xmNote(o)*xmNote(xmSemitonesPerOctave)
}

func (xmMachine) BaseFrequency() float64 {
	return xmBaseFrequency
}

func (xmMachine) CenterNote() note.Note {
	return Machine.Note(xmOctaveForBaseFrequency, keyoctave.KeyC, 0)
}
