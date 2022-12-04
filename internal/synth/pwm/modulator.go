package pwm

import (
	"math"

	"github.com/heucuva/go-qwertysynth/internal/synth/wave"
)

type Modulator interface {
	SetNoteRatio(ratio float64)
	Get() float32
	Advance()
	SetOutputSampleRate(sampleRate float64)
	SetFreqModulation(mod float64)
}

type modulator struct {
	noteRatio     float64
	samplingRatio float64
	fm            float64
	wave          wave.Wave
	pos           float64
}

func NewModulator(w wave.Wave, sampleRate float64) Modulator {
	return &modulator{
		noteRatio:     1.0,
		fm:            1.0,
		wave:          w,
		samplingRatio: w.SampleRate() / sampleRate,
	}
}

func (m *modulator) SetNoteRatio(ratio float64) {
	m.noteRatio = ratio
}

func (m *modulator) SetOutputSampleRate(sampleRate float64) {
	m.samplingRatio = m.wave.SampleRate() / sampleRate
}

func (m modulator) Get() float32 {
	d := m.wave.Data()
	p := int(m.pos)
	if p < 0 || p >= len(d) {
		return 0
	}
	return d[p]
}

func (m *modulator) Advance() {
	m.pos += m.noteRatio * m.fm * m.samplingRatio
	p := int(m.pos)
	d := m.wave.Data()
	l := len(d)
	for p >= l {
		m.pos -= float64(l)
		p = int(m.pos)
	}
}

func (m *modulator) SetFreqModulation(mod float64) {
	m.fm = math.Pow(2.0, mod/10)
}
