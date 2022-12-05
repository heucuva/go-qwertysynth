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
	switch k {
	case ebiten.KeyQ:
		return g.s.Note(g.currentOctave+1, scale.KeyC, 0)
	case ebiten.KeyW:
		return g.s.Note(g.currentOctave+1, scale.KeyCSharp, 0)
	case ebiten.KeyE:
		return g.s.Note(g.currentOctave+1, scale.KeyD, 0)
	case ebiten.KeyR:
		return g.s.Note(g.currentOctave+1, scale.KeyDSharp, 0)
	case ebiten.KeyT:
		return g.s.Note(g.currentOctave+1, scale.KeyE, 0)
	case ebiten.KeyY:
		return g.s.Note(g.currentOctave+1, scale.KeyF, 0)
	case ebiten.KeyU:
		return g.s.Note(g.currentOctave+1, scale.KeyFSharp, 0)
	case ebiten.KeyI:
		return g.s.Note(g.currentOctave+1, scale.KeyG, 0)
	case ebiten.KeyO:
		return g.s.Note(g.currentOctave+1, scale.KeyGSharp, 0)
	case ebiten.KeyP:
		return g.s.Note(g.currentOctave+1, scale.KeyA, 0)
	case ebiten.KeyBracketLeft:
		return g.s.Note(g.currentOctave+1, scale.KeyASharp, 0)
	case ebiten.KeyBracketRight:
		return g.s.Note(g.currentOctave+1, scale.KeyB, 0)

	case ebiten.KeyA:
		return g.s.Note(g.currentOctave, scale.KeyC, 0)
	case ebiten.KeyS:
		return g.s.Note(g.currentOctave, scale.KeyCSharp, 0)
	case ebiten.KeyD:
		return g.s.Note(g.currentOctave, scale.KeyD, 0)
	case ebiten.KeyF:
		return g.s.Note(g.currentOctave, scale.KeyDSharp, 0)
	case ebiten.KeyG:
		return g.s.Note(g.currentOctave, scale.KeyE, 0)
	case ebiten.KeyH:
		return g.s.Note(g.currentOctave, scale.KeyF, 0)
	case ebiten.KeyJ:
		return g.s.Note(g.currentOctave, scale.KeyFSharp, 0)
	case ebiten.KeyK:
		return g.s.Note(g.currentOctave, scale.KeyG, 0)
	case ebiten.KeyL:
		return g.s.Note(g.currentOctave, scale.KeyGSharp, 0)
	case ebiten.KeySemicolon:
		return g.s.Note(g.currentOctave, scale.KeyA, 0)
	case ebiten.KeyQuote:
		return g.s.Note(g.currentOctave, scale.KeyASharp, 0)
	case ebiten.KeyEnter:
		return g.s.Note(g.currentOctave, scale.KeyB, 0)

	case ebiten.KeyZ:
		return g.s.Note(g.currentOctave-1, scale.KeyC, 0)
	case ebiten.KeyX:
		return g.s.Note(g.currentOctave-1, scale.KeyCSharp, 0)
	case ebiten.KeyC:
		return g.s.Note(g.currentOctave-1, scale.KeyD, 0)
	case ebiten.KeyV:
		return g.s.Note(g.currentOctave-1, scale.KeyDSharp, 0)
	case ebiten.KeyB:
		return g.s.Note(g.currentOctave-1, scale.KeyE, 0)
	case ebiten.KeyN:
		return g.s.Note(g.currentOctave-1, scale.KeyF, 0)
	case ebiten.KeyM:
		return g.s.Note(g.currentOctave-1, scale.KeyFSharp, 0)
	case ebiten.KeyComma:
		return g.s.Note(g.currentOctave-1, scale.KeyG, 0)
	case ebiten.KeyPeriod:
		return g.s.Note(g.currentOctave-1, scale.KeyGSharp, 0)
	case ebiten.KeySlash:
		return g.s.Note(g.currentOctave-1, scale.KeyA, 0)
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
