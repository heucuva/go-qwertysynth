package keyboard

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
)

type action int

const (
	actionKeyOn action = iota
	actionKeyOff
)

func (g *Keyboard) processNote(n note.Note, a action, shift bool) {
	switch a {
	case actionKeyOn:
		g.KeyOn(n)
	case actionKeyOff:
		if shift {
			g.KeyCut(n)
		} else {
			g.KeyOff(n)
		}
	}
}

func (g Keyboard) keyNote(k ebiten.Key) note.Note {
	if ko, found := g.keyMap.KeyMap()[k]; found {
		return g.s.Note(g.currentOctave+scale.Octave(ko.Octave), g.s.Machine().Tuning().Key(ko.Key), 0)
	}

	return nil
}

func (g *Keyboard) processKey(k ebiten.Key, a action, shift bool) {
	switch k {
	case ebiten.KeyPageUp:
		if a == actionKeyOn {
			g.SetCurrentOctave(g.currentOctave + 1)
		}

	case ebiten.KeyPageDown:
		if a == actionKeyOn {
			g.SetCurrentOctave(g.currentOctave - 1)
		}

	case ebiten.KeyEscape:
		if a == actionKeyOn {
			g.wantStop = true
		}

	default:
		if n := g.keyNote(k); n != nil {
			g.processNote(n, a, shift)
		}
	}
}
