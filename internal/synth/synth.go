package synth

import (
	"math"
	"sync"
	"time"

	"github.com/gotracker/gomixing/mixing"
	"github.com/gotracker/gomixing/panning"
	"github.com/gotracker/gomixing/volume"
	"github.com/heucuva/go-qwertysynth/internal/machine"
	"github.com/heucuva/go-qwertysynth/internal/output/premix"
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/synth/voice"
	"github.com/heucuva/go-qwertysynth/internal/synth/wavetable"
)

type KeyAction int

const (
	KeyActionOn KeyAction = iota
	KeyActionOff
	KeyActionCut
	KeyActionFadeout
)

type SynthTickFunc func(s Synth, amt time.Duration) error

type Synth interface {
	Run(onTick SynthTickFunc) error
	Close()
	C() <-chan *premix.PremixData
	Default() machine.Default
	Note(o scale.Octave, k scale.Key, s scale.Microtone) note.Note
	KeyAction(n note.Note, a KeyAction)
}

type synth struct {
	tickInterval time.Duration

	outCh   chan *premix.PremixData
	state   map[scale.KeyOctave]voice.Voice
	stateMu sync.RWMutex

	bufLen     int
	voices     wavetable.WaveTable
	sampleRate float64
	sampleDur  time.Duration
	mixVolume  volume.Volume
}

func NewSynth(tickInterval time.Duration, voices wavetable.WaveTable, numPremixBuffers int, sampleRate float64, mixVolume float32) Synth {
	s := &synth{
		tickInterval: tickInterval,
		outCh:        make(chan *premix.PremixData, numPremixBuffers),
		state:        make(map[scale.KeyOctave]voice.Voice),
		bufLen:       int(math.Ceil(sampleRate * tickInterval.Seconds())),
		voices:       voices,
		sampleRate:   sampleRate,
		sampleDur:    time.Duration(float64(time.Second) / sampleRate),
		mixVolume:    volume.Volume(mixVolume),
	}

	return s
}

func (s *synth) Close() {
	if s.outCh != nil {
		ch := s.outCh
		s.outCh = nil
		close(ch)
	}
}

func (s *synth) C() <-chan *premix.PremixData {
	return s.outCh
}

func (s *synth) Run(onTick SynthTickFunc) error {
	tick := time.NewTicker(s.tickInterval)
	defer tick.Stop()

	for s.outCh != nil {
		if onTick != nil {
			if err := onTick(s, s.tickInterval); err != nil {
				return err
			}
		}

		s.processBuffer()

		<-tick.C
	}

	return nil
}

func (s *synth) Default() machine.Default {
	return s.voices.Default()
}

func (s *synth) Note(o scale.Octave, k scale.Key, st scale.Microtone) note.Note {
	return s.voices.Note(o, k, st)
}

func (s *synth) KeyAction(n note.Note, a KeyAction) {
	switch a {
	case KeyActionOn:
		s.keyOn(n)
	default:
		s.keyAction(n, a)
	}
}

func (s *synth) keyOn(n note.Note) {
	if s.state == nil {
		return
	}

	ko := n.KeyOctave()

	v := s.voices.Get(n)
	v.SetOutputSampleRate(s.sampleRate)
	v.KeyOn()

	s.stateMu.Lock()
	defer s.stateMu.Unlock()

	s.state[ko] = v
}

func (s *synth) keyAction(n note.Note, a KeyAction) {
	if s.state == nil {
		return
	}

	s.stateMu.RLock()
	defer s.stateMu.RUnlock()

	ko := n.KeyOctave()
	v, found := s.state[ko]
	if !found {
		return
	}

	switch a {
	case KeyActionOff:
		v.KeyOff()
	case KeyActionCut:
		v.Cut()
	default:
	}
}

func (s *synth) processBuffer() {
	s.stateMu.Lock()
	defer s.stateMu.Unlock()

	nextState := make(map[scale.KeyOctave]voice.Voice)
	state := s.state
	s.state = nextState

	var voices []voice.Voice
	for k, v := range state {
		if v == nil || !v.IsPlaying() {
			continue
		}

		s.state[k] = v
		voices = append(voices, v)
	}

	go s.mixBuffer(voices)
}

func (s *synth) mixBuffer(voices []voice.Voice) {
	buf := &premix.PremixData{
		SamplesLen:  s.bufLen,
		MixerVolume: 1.0,
	}

	if len(voices) == 0 {
		voices = append(voices, nil)
	}

	for _, v := range voices {
		mix := make(mixing.MixBuffer, s.bufLen)

		for i := 0; i < s.bufLen; i++ {
			mix[i].Channels = 1
			if v != nil {
				if vol, playing := v.Get(); playing {
					mix[i].StaticMatrix[0] = volume.Volume(vol)
				}
				v.Advance(s.sampleDur)
			}
		}

		var cdata mixing.ChannelData
		cdata = append(cdata, mixing.Data{
			Data:       mix,
			Pan:        panning.CenterAhead,
			Volume:     s.mixVolume,
			Pos:        0,
			SamplesLen: s.bufLen,
		})
		buf.Data = append(buf.Data, cdata)
	}

	s.outCh <- buf
}
