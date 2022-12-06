package keyboard

import (
	"github.com/heucuva/go-qwertysynth/internal/machine"
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/synth/keymap"
	"github.com/heucuva/go-qwertysynth/internal/synth/voice"
	"github.com/heucuva/go-qwertysynth/internal/synth/wavetable"
)

type keyboard struct {
	mach   machine.Machine
	voices []*wavetable.Item
	km     keymap.KeyMap
}

func NewKeyboard(mach machine.Machine, km keymap.KeyMap, voices []*wavetable.Item) wavetable.WaveTable {
	return &keyboard{
		mach:   mach,
		voices: voices,
		km:     km,
	}
}

func (t keyboard) Get(n note.Note) voice.Voice {
	if t.voices == nil {
		return nil
	}

	ko := n.KeyOctave()

	idx := t.km[ko]

	if len(t.voices) <= int(idx) {
		return nil
	}

	w := t.voices[idx]
	if w == nil {
		return nil
	}

	v := w.Voice()
	v.SetNote(n)

	return v
}

func (t keyboard) Machine() machine.Machine {
	return t.mach
}

func (t keyboard) Note(o scale.Octave, k scale.Key, s scale.Microtone) note.Note {
	return t.mach.Note(o, k, s)
}
