package keyboard

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/heucuva/go-qwertysynth/internal/input/keyboard/keymapper"
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/synth"
)

var (
	ErrWantStop = errors.New("want stop")
)

type Keyboard struct {
	keys          []ebiten.Key
	currentOctave scale.Octave
	lastTick      time.Time
	wantStop      bool
	onTick        TickFunc
	s             synth.Synth
	ctx           context.Context
	keyMap        keymapper.KeyMapper
}

type TickFunc func(g *Keyboard, amt time.Duration) error

func NewKeyboard(ctx context.Context, s synth.Synth, onTick TickFunc, showHelp bool) *Keyboard {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Keyboard (qwertysynth)")

	tuning := s.Machine().Tuning()
	_, o := tuning.BaseKey()

	var keyMap keymapper.KeyMapper
	keysPerOctave := tuning.KeysPerOctave()
	switch keysPerOctave {
	case 12:
		keyMap = keymapper.Twelve
	case 53:
		keyMap = keymapper.FiftyThree
	default:
		panic("unhandled number of keys per octave")
	}

	g := &Keyboard{
		lastTick:      time.Now(),
		currentOctave: o,
		onTick:        onTick,
		s:             s,
		ctx:           ctx,
		keyMap:        keyMap,
	}

	if showHelp {
		g.showHelp()
	}

	fmt.Println("Press Escape to quit")

	return g
}

func (g Keyboard) showHelp() {
	g.keyMap.ShowHelp(g.currentOctave, g.s)

	fmt.Println()

	fmt.Println("Ready")
}

func (g *Keyboard) KeyOn(n note.Note) {
	g.s.KeyAction(n, synth.KeyActionOn)
	fmt.Printf("%v\n", n)
}

func (g *Keyboard) KeyOff(n note.Note) {
	g.s.KeyAction(n, synth.KeyActionOff)
}

func (g *Keyboard) KeyCut(n note.Note) {
	g.s.KeyAction(n, synth.KeyActionCut)
}

func (g *Keyboard) SetCurrentOctave(o scale.Octave) {
	if o > scale.MaxOctave-1 {
		o = scale.MaxOctave - 1
	}
	if o < scale.MinOctave+1 {
		o = scale.MinOctave + 1
	}
	g.currentOctave = o
	fmt.Printf("%s-row octave: %d\n", g.keyMap.CenterRowKey(), g.currentOctave)
}

func (Keyboard) isShiftPressed() bool {
	return inpututil.KeyPressDuration(ebiten.KeyShift) > 0 ||
		inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 ||
		inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0
}

func (g *Keyboard) Update() error {
	select {
	case <-g.ctx.Done():
		return ErrWantStop
	default:
		if g.wantStop {
			return ErrWantStop
		}
	}

	currentKeys := inpututil.AppendPressedKeys(nil)
	shift := g.isShiftPressed()

	for _, key := range currentKeys {
		if inpututil.IsKeyJustPressed(key) {
			g.processKey(key, actionKeyOn, shift)
		}
	}

	for _, key := range g.keys {
		if inpututil.IsKeyJustReleased(key) {
			g.processKey(key, actionKeyOff, shift)
		}
	}

	g.keys = currentKeys

	tickDur := time.Since(g.lastTick)
	g.lastTick = time.Now()

	if g.onTick != nil {
		g.onTick(g, tickDur)
	}

	return nil
}

func (g *Keyboard) Draw(screen *ebiten.Image) {
	const (
		offsetX = 24
		offsetY = 40
	)

	// Draw the base (grayed) keyboard image.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(offsetX, offsetY)
	op.ColorM.Scale(0.5, 0.5, 0.5, 1)
	screen.DrawImage(g.keyMap.KeyboardImage(), op)

	// Draw the highlighted keys.
	op = &ebiten.DrawImageOptions{}
	for _, p := range g.keys {
		op.GeoM.Reset()
		r, ok := g.keyMap.KeyRect(p)
		if !ok {
			continue
		}
		op.GeoM.Translate(float64(r.Min.X), float64(r.Min.Y))
		op.GeoM.Translate(offsetX, offsetY)
		screen.DrawImage(g.keyMap.KeyboardImage().SubImage(r).(*ebiten.Image), op)
	}

	keyStrs := []string{}
	for _, k := range g.keys {
		if n := g.keyNote(k); n != nil {
			keyStrs = append(keyStrs, fmt.Sprint(n))
		}
	}
	ebitenutil.DebugPrint(screen, strings.Join(keyStrs, ", "))
}

func (g *Keyboard) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

const (
	screenWidth  = 320
	screenHeight = 240
)
