package keymapper

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/synth"
)

type KeyMapper interface {
	KeyMap() map[ebiten.Key]KeyOctave
	ShowHelp(currentOctave scale.Octave, s synth.Synth)
	KeyboardImage() *ebiten.Image
	CenterRowKey() ebiten.Key
	KeyRect(k ebiten.Key) (image.Rectangle, bool)
}

type KeyOctave struct {
	Key    int
	Octave int
}
