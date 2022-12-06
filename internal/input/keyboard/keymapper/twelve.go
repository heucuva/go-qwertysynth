package keymapper

import (
	"bytes"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	rkeyboard "github.com/heucuva/go-qwertysynth/internal/input/keyboard/resources/images/keyboard/twelve"
	rtwelve "github.com/heucuva/go-qwertysynth/internal/input/keyboard/resources/twelve"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/synth"
)

type twelve struct{}

var Twelve = &twelve{}

var twelveMap map[ebiten.Key]KeyOctave

func (twelve) KeyMap() map[ebiten.Key]KeyOctave {
	return twelveMap
}

func (twelve) CenterRowKey() ebiten.Key {
	return ebiten.KeyA
}

func (twelve) KeyRect(k ebiten.Key) (image.Rectangle, bool) {
	return rtwelve.KeyRect(k)
}

func (m twelve) ShowHelp(currentOctave scale.Octave, s synth.Synth) {
	tuning := s.Machine().Tuning()
	rowStartHelp := func(key ebiten.Key) {
		ko := m.KeyMap()[key]
		rowNote := s.Note(currentOctave+scale.Octave(ko.Octave), tuning.Key(ko.Key), 0)
		fmt.Printf("%s-row starts with %v\n", key, rowNote)
	}

	rowStartHelp(ebiten.KeyQ)
	rowStartHelp(ebiten.KeyA)
	rowStartHelp(ebiten.KeyZ)

	fmt.Println()

	fmt.Println("Hold keys to sustain notes; release them to decay them")
	fmt.Println("Release keys while holding Shift to cut/stop them")
	fmt.Println("PageUp to increase keyboard octave; PageDown to decrease keyboard octave")
	fmt.Println("Note: US English keyboard layout works best")
}

var twelveKeyboardImage *ebiten.Image

func (twelve) KeyboardImage() *ebiten.Image {
	return twelveKeyboardImage
}

func init() {
	twelveMap = make(map[ebiten.Key]KeyOctave)
	twelveMap[ebiten.KeyQ] = KeyOctave{Octave: 1, Key: 0}
	twelveMap[ebiten.KeyW] = KeyOctave{Octave: 1, Key: 1}
	twelveMap[ebiten.KeyE] = KeyOctave{Octave: 1, Key: 2}
	twelveMap[ebiten.KeyR] = KeyOctave{Octave: 1, Key: 3}
	twelveMap[ebiten.KeyT] = KeyOctave{Octave: 1, Key: 4}
	twelveMap[ebiten.KeyY] = KeyOctave{Octave: 1, Key: 5}
	twelveMap[ebiten.KeyU] = KeyOctave{Octave: 1, Key: 6}
	twelveMap[ebiten.KeyI] = KeyOctave{Octave: 1, Key: 7}
	twelveMap[ebiten.KeyO] = KeyOctave{Octave: 1, Key: 8}
	twelveMap[ebiten.KeyP] = KeyOctave{Octave: 1, Key: 9}
	twelveMap[ebiten.KeyBracketLeft] = KeyOctave{Octave: 1, Key: 10}
	twelveMap[ebiten.KeyBracketRight] = KeyOctave{Octave: 1, Key: 11}

	twelveMap[ebiten.KeyA] = KeyOctave{Octave: 0, Key: 0}
	twelveMap[ebiten.KeyS] = KeyOctave{Octave: 0, Key: 1}
	twelveMap[ebiten.KeyD] = KeyOctave{Octave: 0, Key: 2}
	twelveMap[ebiten.KeyF] = KeyOctave{Octave: 0, Key: 3}
	twelveMap[ebiten.KeyG] = KeyOctave{Octave: 0, Key: 4}
	twelveMap[ebiten.KeyH] = KeyOctave{Octave: 0, Key: 5}
	twelveMap[ebiten.KeyJ] = KeyOctave{Octave: 0, Key: 6}
	twelveMap[ebiten.KeyK] = KeyOctave{Octave: 0, Key: 7}
	twelveMap[ebiten.KeyL] = KeyOctave{Octave: 0, Key: 8}
	twelveMap[ebiten.KeySemicolon] = KeyOctave{Octave: 0, Key: 9}
	twelveMap[ebiten.KeyQuote] = KeyOctave{Octave: 0, Key: 10}
	twelveMap[ebiten.KeyEnter] = KeyOctave{Octave: 0, Key: 11}

	twelveMap[ebiten.KeyZ] = KeyOctave{Octave: -1, Key: 0}
	twelveMap[ebiten.KeyX] = KeyOctave{Octave: -1, Key: 1}
	twelveMap[ebiten.KeyC] = KeyOctave{Octave: -1, Key: 2}
	twelveMap[ebiten.KeyV] = KeyOctave{Octave: -1, Key: 3}
	twelveMap[ebiten.KeyB] = KeyOctave{Octave: -1, Key: 4}
	twelveMap[ebiten.KeyN] = KeyOctave{Octave: -1, Key: 5}
	twelveMap[ebiten.KeyM] = KeyOctave{Octave: -1, Key: 6}
	twelveMap[ebiten.KeyComma] = KeyOctave{Octave: -1, Key: 7}
	twelveMap[ebiten.KeyPeriod] = KeyOctave{Octave: -1, Key: 8}
	twelveMap[ebiten.KeySlash] = KeyOctave{Octave: -1, Key: 9}

	img, _, err := image.Decode(bytes.NewReader(rkeyboard.Keyboard_png))
	if err != nil {
		log.Fatal(err)
	}

	twelveKeyboardImage = ebiten.NewImageFromImage(img)
}
