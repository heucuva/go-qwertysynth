package xm

import (
	"github.com/heucuva/go-qwertysynth/internal/machine"
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
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

func (xmMachine) Generate(tuning tuning.Tuning, generator wave.Generator, opts ...wave.GeneratorParam) (wave.Wave, error) {
	var machineOpts []wave.GeneratorParam
	machineOpts = append(machineOpts,
		wave.SetParameterByName("frequency", baseNote.ToFrequency(tuning)),
		wave.SetParameterByName("sampleRate", xmBaseFrequency),
	)
	return generator(append(machineOpts, opts...)...)
}

func (xmMachine) Tuning() tuning.Tuning {
	return defaultTuning
}

const (
	xmMicrotonesPerKey       = 64
	xmMinOctave              = 1
	xmMaxOctave              = 8
	xmMicrotonesPerOctave    = xmMicrotonesPerKey * scale.KeysPerOctave
	xmTotalOctaves           = xmMaxOctave - xmMinOctave + 1
	xmMaxMicrotones          = xmMicrotonesPerOctave * xmTotalOctaves
	xmMinMicrotone           = 0
	xmBaseFrequency          = 8363.0
	xmOctaveForBaseFrequency = 4
)

func (xmMachine) NoteFromChannelData(d uint8) (note.Note, error) {
	switch {
	case d == 0:
		return note.None, nil
	case d >= 1 && d <= 96:
		v := d - 1
		o := xmMinOctave + scale.Octave(v/uint8(scale.KeysPerOctave))
		k := scale.MinKey + scale.Key(v%uint8(scale.KeysPerOctave))
		return Machine.Note(o, k, 0), nil
	case d == 97:
		return note.Cut, nil
	default:
		return note.None, ErrChannelDataInvalid
	}
}

func (xmMachine) Note(o scale.Octave, k scale.Key, s scale.Microtone) note.Note {
	return xmNote(s) +
		xmNote(k)*xmMicrotonesPerKey +
		xmNote(o)*xmNote(xmMicrotonesPerOctave)
}

func (xmMachine) BaseFrequency() float64 {
	return xmBaseFrequency
}

func (xmMachine) CenterNote() note.Note {
	return Machine.Note(xmOctaveForBaseFrequency, scale.KeyC, 0)
}
