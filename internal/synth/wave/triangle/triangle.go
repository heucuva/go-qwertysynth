package triangle

import (
	"math"
	"time"

	"github.com/heucuva/go-qwertysynth/internal/synth/wave"
	"github.com/pkg/errors"
)

type triangle struct {
	data       []float32
	sampleRate float64
}

func (w triangle) Data() []float32 {
	return w.data
}

func (w triangle) SampleRate() float64 {
	return w.sampleRate
}

var _ wave.Generator = Triangle

func Triangle(opts ...wave.GeneratorParam) (wave.Wave, error) {
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
	w := &triangle{
		data:       make([]float32, length),
		sampleRate: s.sampleRate,
	}

	periodRadians := 2.0 * math.Pi
	radConverter := periodRadians / samplesPerPeriod

	a := float64(s.amp)
	p := periodRadians / 2.0

	theta := 1.5*p + float64(s.phase)
	for i := range w.data {
		mod := math.Mod(theta, p*2.0)
		y := a - (2.0*a/p)*(p-math.Abs(mod-p))

		w.data[i] = float32(y)
		theta += radConverter
	}

	return w, nil
}
