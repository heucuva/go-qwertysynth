package square

import (
	"time"

	"github.com/heucuva/go-qwertysynth/internal/standards/metric"
	"github.com/heucuva/go-qwertysynth/internal/standards/scale"
	"github.com/heucuva/go-qwertysynth/internal/synth/wave"
	"github.com/pkg/errors"
)

type settings struct {
	amp   float32
	phase metric.Radians

	frequency, sampleRate float64
	duration              time.Duration
}

func (s *settings) SetParameterByName(name string, value any) error {
	switch name {
	case "amplitude", "amp", "a":
		if amp, ok := value.(float32); !ok {
			return errors.Wrap(wave.ErrInvalidParameterValue, name)
		} else {
			s.amp = amp
		}
	case "phase":
		if phase, ok := value.(metric.Radians); !ok {
			return errors.Wrap(wave.ErrInvalidParameterValue, name)
		} else {
			s.phase = phase
		}
	case "frequency", "freq", "f":
		if freq, ok := value.(float64); !ok {
			return errors.Wrap(wave.ErrInvalidParameterValue, name)
		} else {
			s.frequency = freq
		}
	case "sampleRate":
		if sampleRate, ok := value.(float64); !ok {
			return errors.Wrap(wave.ErrInvalidParameterValue, name)
		} else {
			s.sampleRate = sampleRate
		}
	case "duration":
		if duration, ok := value.(time.Duration); !ok {
			return errors.Wrap(wave.ErrInvalidParameterValue, name)
		} else {
			s.duration = duration
		}
	default:
		return errors.Wrap(wave.ErrNotValidForThisGenerator, name)
	}
	return nil
}

const (
	defaultAmp                       = 1.0
	defaultPhase      metric.Radians = 0.0
	defaultFrequency                 = scale.A440_A4Frequency
	defaultSampleRate                = 48000
	defaultDuration   time.Duration  = 0
)

func squareParam(fn func(gen *settings) error) wave.GeneratorParam {
	return func(o interface{}) error {
		gen, ok := o.(*settings)
		if !ok {
			return wave.ErrNotValidForThisGenerator
		}

		return fn(gen)
	}
}

func Amplitude(amp float32) wave.GeneratorParam {
	return squareParam(func(gen *settings) error {
		gen.amp = amp
		return nil
	})
}

func Frequency(freq float64) wave.GeneratorParam {
	return squareParam(func(gen *settings) error {
		gen.frequency = freq
		return nil
	})
}

func Phase(phase metric.Radians) wave.GeneratorParam {
	return squareParam(func(gen *settings) error {
		gen.phase = phase
		return nil
	})
}

func SampleRate(sampleRate float64) wave.GeneratorParam {
	return squareParam(func(gen *settings) error {
		gen.sampleRate = sampleRate
		return nil
	})
}

func Duration(duration time.Duration) wave.GeneratorParam {
	return squareParam(func(gen *settings) error {
		gen.duration = duration
		return nil
	})
}
