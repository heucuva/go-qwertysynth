package square

import (
	"math"
	"time"

	"github.com/heucuva/go-qwertysynth/internal/synth/wave"
	"github.com/pkg/errors"
)

type square struct {
	data       []float32
	sampleRate float64
}

func (w square) Data() []float32 {
	return w.data
}

func (w square) SampleRate() float64 {
	return w.sampleRate
}

var _ wave.Generator = Square

func Square(opts ...wave.GeneratorParam) (wave.Wave, error) {
	s := settings{
		amp:        defaultAmp,
		phase:      defaultPhase,
		frequency:  defaultFrequency,
		sampleRate: defaultSampleRate,
		duration:   defaultDuration,
	}

	for _, opt := range opts {
		if err := opt(&s); err != nil {
			return nil, err
		}
	}

	if s.frequency == 0 {
		return nil, errors.Wrap(wave.ErrInvalidParameterValue, "frequency")
	}

	if s.duration == 0 {
		onePeriod := float64(time.Second) / s.frequency
		s.duration = time.Duration(onePeriod)
	}
	length := int(math.Round(s.duration.Seconds() * s.sampleRate))
	samplesPerPeriod := s.sampleRate / s.frequency
	w := &square{
		data:       make([]float32, length),
		sampleRate: s.sampleRate,
	}

	radConverter := 2.0 * math.Pi / samplesPerPeriod

	theta := float64(s.phase)
	for i := range w.data {
		v := float32(math.Sin(theta))
		if v > 0 {
			v = 1.0
		} else if v < 0 {
			v = -1.0
		}
		w.data[i] = s.amp * v
		theta += radConverter
	}

	return w, nil
}
