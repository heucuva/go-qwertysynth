package keymapper

import (
	"bytes"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	rfiftythree "github.com/heucuva/go-qwertysynth/internal/input/keyboard/resources/fiftythree"
	rkeyboard "github.com/heucuva/go-qwertysynth/internal/input/keyboard/resources/images/keyboard/fiftythree"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/synth"
)

type fiftyThree struct{}

var FiftyThree = &fiftyThree{}

var fiftyThreeMap map[ebiten.Key]KeyOctave

func (fiftyThree) KeyMap() map[ebiten.Key]KeyOctave {
	return fiftyThreeMap
}

func (fiftyThree) CenterRowKey() ebiten.Key {
	return ebiten.KeyZ
}

func (fiftyThree) KeyRect(k ebiten.Key) (image.Rectangle, bool) {
	return rfiftythree.KeyRect(k)
}

func (m fiftyThree) ShowHelp(currentOctave scale.Octave, s synth.Synth) {
	rowStartHelp := func(key ebiten.Key) {
		ko := m.KeyMap()[key]
		rowNote := s.Note(currentOctave+scale.Octave(ko.Octave), s.Machine().Tuning().Key(ko.Key), 0)
		fmt.Printf("%s-row starts with %v\n", key, rowNote)
	}

	rowStartHelp(ebiten.KeyF1)
	rowStartHelp(ebiten.Key1)
	rowStartHelp(ebiten.KeyQ)
	rowStartHelp(ebiten.KeyA)
	rowStartHelp(ebiten.KeyZ)

	fmt.Println()

	fmt.Println("Hold keys to sustain notes; release them to decay them")
	fmt.Println("Release keys while holding Shift to cut/stop them")
	fmt.Println("PageUp to increase keyboard octave; PageDown to decrease keyboard octave")
	fmt.Println("Note: US English keyboard layout works best")
}

var fiftyThreeKeyboardImage *ebiten.Image

func (fiftyThree) KeyboardImage() *ebiten.Image {
	return fiftyThreeKeyboardImage
}

func init() {
	fiftyThreeMap = make(map[ebiten.Key]KeyOctave)
	fiftyThreeMap[ebiten.KeyZ] = KeyOctave{Key: 0}
	fiftyThreeMap[ebiten.KeyX] = KeyOctave{Key: 1}
	fiftyThreeMap[ebiten.KeyC] = KeyOctave{Key: 2}
	fiftyThreeMap[ebiten.KeyV] = KeyOctave{Key: 3}
	fiftyThreeMap[ebiten.KeyB] = KeyOctave{Key: 4}
	fiftyThreeMap[ebiten.KeyN] = KeyOctave{Key: 5}
	fiftyThreeMap[ebiten.KeyM] = KeyOctave{Key: 6}
	fiftyThreeMap[ebiten.KeyComma] = KeyOctave{Key: 7}
	fiftyThreeMap[ebiten.KeyPeriod] = KeyOctave{Key: 8}
	fiftyThreeMap[ebiten.KeySlash] = KeyOctave{Key: 9}
	fiftyThreeMap[ebiten.KeyA] = KeyOctave{Key: 10}
	fiftyThreeMap[ebiten.KeyS] = KeyOctave{Key: 11}
	fiftyThreeMap[ebiten.KeyD] = KeyOctave{Key: 12}
	fiftyThreeMap[ebiten.KeyF] = KeyOctave{Key: 13}
	fiftyThreeMap[ebiten.KeyG] = KeyOctave{Key: 14}
	fiftyThreeMap[ebiten.KeyH] = KeyOctave{Key: 15}
	fiftyThreeMap[ebiten.KeyJ] = KeyOctave{Key: 16}
	fiftyThreeMap[ebiten.KeyK] = KeyOctave{Key: 17}
	fiftyThreeMap[ebiten.KeyL] = KeyOctave{Key: 18}
	fiftyThreeMap[ebiten.KeySemicolon] = KeyOctave{Key: 19}
	fiftyThreeMap[ebiten.KeyQuote] = KeyOctave{Key: 20}
	fiftyThreeMap[ebiten.KeyQ] = KeyOctave{Key: 21}
	fiftyThreeMap[ebiten.KeyW] = KeyOctave{Key: 22}
	fiftyThreeMap[ebiten.KeyE] = KeyOctave{Key: 23}
	fiftyThreeMap[ebiten.KeyR] = KeyOctave{Key: 24}
	fiftyThreeMap[ebiten.KeyT] = KeyOctave{Key: 25}
	fiftyThreeMap[ebiten.KeyY] = KeyOctave{Key: 26}
	fiftyThreeMap[ebiten.KeyU] = KeyOctave{Key: 27}
	fiftyThreeMap[ebiten.KeyI] = KeyOctave{Key: 28}
	fiftyThreeMap[ebiten.KeyO] = KeyOctave{Key: 29}
	fiftyThreeMap[ebiten.KeyP] = KeyOctave{Key: 30}
	fiftyThreeMap[ebiten.KeyBracketLeft] = KeyOctave{Key: 31}
	fiftyThreeMap[ebiten.KeyBracketRight] = KeyOctave{Key: 32}
	fiftyThreeMap[ebiten.Key1] = KeyOctave{Key: 33}
	fiftyThreeMap[ebiten.Key2] = KeyOctave{Key: 34}
	fiftyThreeMap[ebiten.Key3] = KeyOctave{Key: 35}
	fiftyThreeMap[ebiten.Key4] = KeyOctave{Key: 36}
	fiftyThreeMap[ebiten.Key5] = KeyOctave{Key: 37}
	fiftyThreeMap[ebiten.Key6] = KeyOctave{Key: 38}
	fiftyThreeMap[ebiten.Key7] = KeyOctave{Key: 39}
	fiftyThreeMap[ebiten.Key8] = KeyOctave{Key: 40}
	fiftyThreeMap[ebiten.Key9] = KeyOctave{Key: 41}
	fiftyThreeMap[ebiten.Key0] = KeyOctave{Key: 42}
	fiftyThreeMap[ebiten.KeyMinus] = KeyOctave{Key: 43}
	fiftyThreeMap[ebiten.KeyEqual] = KeyOctave{Key: 44}
	fiftyThreeMap[ebiten.KeyF1] = KeyOctave{Key: 45}
	fiftyThreeMap[ebiten.KeyF2] = KeyOctave{Key: 46}
	fiftyThreeMap[ebiten.KeyF3] = KeyOctave{Key: 47}
	fiftyThreeMap[ebiten.KeyF4] = KeyOctave{Key: 48}
	fiftyThreeMap[ebiten.KeyF5] = KeyOctave{Key: 49}
	fiftyThreeMap[ebiten.KeyF6] = KeyOctave{Key: 50}
	fiftyThreeMap[ebiten.KeyF7] = KeyOctave{Key: 51}
	fiftyThreeMap[ebiten.KeyF8] = KeyOctave{Key: 52}

	img, _, err := image.Decode(bytes.NewReader(rkeyboard.Keyboard_png))
	if err != nil {
		log.Fatal(err)
	}

	fiftyThreeKeyboardImage = ebiten.NewImageFromImage(img)
}
