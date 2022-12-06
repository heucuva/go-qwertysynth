package it

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/heucuva/go-qwertysynth/internal/machine"
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	tuningPkg "github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	"github.com/heucuva/go-qwertysynth/internal/synth/wave"
)

type itMachine struct {
	tuning tuningPkg.Tuning
}

func Machine(tuning tuningPkg.Tuning) machine.Machine {
	if tuning == nil {
		tuning = defaultTuning
	}
	return &itMachine{
		tuning: tuning,
	}
}

func (itMachine) Default() machine.Default {
	return itDefault
}

func (m itMachine) Generate(generator wave.Generator, opts ...wave.GeneratorParam) (wave.Wave, error) {
	k, o := m.tuning.BaseKey()
	baseNote := m.Note(o, k, 0)
	var machineOpts []wave.GeneratorParam
	machineOpts = append(machineOpts,
		wave.SetParameterByName("frequency", baseNote.ToFrequency()),
		wave.SetParameterByName("sampleRate", itBaseFrequency),
	)
	return generator(append(machineOpts, opts...)...)
}

const (
	itMicrotonesPerKey       = 64
	itMinOctave              = 0
	itMaxOctave              = 9
	itKeysPerOctave          = 12
	itMicrotonesPerOctave    = itMicrotonesPerKey * itKeysPerOctave
	itTotalOctaves           = itMaxOctave - itMinOctave + 1
	itMaxMicrotones          = itMicrotonesPerOctave * itTotalOctaves
	itMinMicrotone           = 0
	itBaseFrequency          = 8363.0
	itOctaveForBaseFrequency = 5
)

func (m itMachine) NoteFromChannelData(d uint8) (note.Note, error) {
	switch {
	case d == 0:
		return note.None, nil
	case d >= 1 && d <= 96:
		v := d - 1
		o := itMinOctave + scale.Octave(v/uint8(itKeysPerOctave))
		krat := float64(m.tuning.KeysPerOctave()) / float64(itKeysPerOctave)
		kv := v % uint8(itKeysPerOctave)
		k := m.tuning.Key(int(math.Round(float64(kv) * krat)))
		return m.Note(o, k, 0), nil
	case d == 97:
		return note.Cut, nil
	default:
		return note.None, ErrChannelDataInvalid
	}
}

func (m itMachine) Note(o scale.Octave, k scale.Key, s scale.Microtone) note.Note {
	keysPerOctave := m.tuning.KeysPerOctave()
	microtonesPerOctave := itMicrotonesPerKey * keysPerOctave
	oks := int(s) +
		k.Index()*itMicrotonesPerKey +
		int(o)*microtonesPerOctave
	return itNote{
		oks:    oks,
		tuning: m.tuning,
	}
}

func (itMachine) BaseFrequency() float64 {
	return itBaseFrequency
}

func (m itMachine) Tuning() tuningPkg.Tuning {
	return m.tuning
}

func (m itMachine) BaseNote() note.Note {
	k, o := m.tuning.BaseKey()
	baseNote := m.Note(o, k, 0)
	return baseNote
}

func (m itMachine) ParseNote(str string) (note.Note, error) {
	if len(str) < 3 {
		return nil, errors.New("insufficient characters in string")
	}

	if str == "---" {
		return note.Cut, nil
	}

	ks, os := str[0:2], str[2:]

	o, err := strconv.ParseInt(os, 10, 64)
	if err != nil {
		return nil, err
	}

	switch strings.ToLower(ks) {
	case "c-":
		return m.noteFromPieces(int(o), 0)
	case "c#":
		return m.noteFromPieces(int(o), 1)
	case "d-":
		return m.noteFromPieces(int(o), 2)
	case "d#":
		return m.noteFromPieces(int(o), 3)
	case "e-":
		return m.noteFromPieces(int(o), 4)
	case "f-":
		return m.noteFromPieces(int(o), 5)
	case "f#":
		return m.noteFromPieces(int(o), 6)
	case "g-":
		return m.noteFromPieces(int(o), 7)
	case "g#":
		return m.noteFromPieces(int(o), 8)
	case "a-":
		return m.noteFromPieces(int(o), 9)
	case "a#":
		return m.noteFromPieces(int(o), 10)
	case "b-":
		return m.noteFromPieces(int(o), 11)
	}

	return nil, errors.New("invalid note string")
}

func (m itMachine) noteFromPieces(o int, k int) (note.Note, error) {
	var data uint8 = 1
	data += uint8(k)
	data += uint8(o-itMinOctave) * itKeysPerOctave
	return m.NoteFromChannelData(data)
}
