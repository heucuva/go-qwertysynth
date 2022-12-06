package command

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/cobra"

	ebikey "github.com/heucuva/go-qwertysynth/internal/input/keyboard"
	"github.com/heucuva/go-qwertysynth/internal/machine"
	"github.com/heucuva/go-qwertysynth/internal/machine/it"
	"github.com/heucuva/go-qwertysynth/internal/machine/xm"
	"github.com/heucuva/go-qwertysynth/internal/output"
	deviceCommon "github.com/heucuva/go-qwertysynth/internal/output/device/common"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
	equalTuning "github.com/heucuva/go-qwertysynth/internal/standards/tuning/equal"
	justTuning "github.com/heucuva/go-qwertysynth/internal/standards/tuning/just"
	"github.com/heucuva/go-qwertysynth/internal/synth"
	"github.com/heucuva/go-qwertysynth/internal/synth/envelope"
	"github.com/heucuva/go-qwertysynth/internal/synth/keyboard"
	"github.com/heucuva/go-qwertysynth/internal/synth/keymap"
	"github.com/heucuva/go-qwertysynth/internal/synth/wave"
	"github.com/heucuva/go-qwertysynth/internal/synth/wave/sine"
	"github.com/heucuva/go-qwertysynth/internal/synth/wave/square"
	"github.com/heucuva/go-qwertysynth/internal/synth/wave/triangle"
	"github.com/heucuva/go-qwertysynth/internal/synth/wavetable"
)

// persistent flags
var (
	playOutput = deviceCommon.Settings{
		Channels:         2,
		SamplesPerSecond: 44100,
		BitsPerSample:    16,
	}
	playNumPremixBuffers int = 64

	playTickLength time.Duration = time.Millisecond * 150

	playMachineName string = "xm"
	playTuningName  string = "default"

	playAM string = "square,adsr:150ms:50ms:-6.2db:1s"
	playFM string = "sine,adsr:1s:80ms:-12.75db:1s,frequency:1.125,amplitude:16"
)

func init() {
	output.Setup()

	persistFlags := playCmd.PersistentFlags()
	persistFlags.IntVarP(&playOutput.SamplesPerSecond, "sample-rate", "S", playOutput.SamplesPerSecond, "sample rate")
	persistFlags.IntVarP(&playOutput.Channels, "channels", "c", playOutput.Channels, "channels")
	persistFlags.IntVarP(&playOutput.BitsPerSample, "bits-per-sample", "b", playOutput.BitsPerSample, "bits per sample")
	persistFlags.IntVar(&playNumPremixBuffers, "num-buffers", playNumPremixBuffers, "number of premixed buffers")
	persistFlags.StringVarP(&playOutput.Name, "output", "O", output.DefaultOutputDeviceName, "output device")
	persistFlags.StringVarP(&playMachineName, "machine", "m", playMachineName, "name of machine to use [xm, it]")
	persistFlags.DurationVarP(&playTickLength, "tick", "T", playTickLength, "tick interval")
	persistFlags.StringVarP(&playAM, "am", "a", playAM, "Amplitude Modulator settings")
	persistFlags.StringVarP(&playFM, "fm", "f", playFM, "Frequency Modulator settings")
	persistFlags.StringVar(&playTuningName, "tuning", playTuningName, "Tuning system to use")

	rootCmd.AddCommand(playCmd)
}

var playCmd = &cobra.Command{
	Use:   "play [flags]",
	Short: "Play synth sounds using the keyboard",
	Long:  "Play synth sounds using the keyboard.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return playSynth(playGetMachine(), nil, true)
	},
}

func playGetMachine() machine.Machine {
	tuning := playGetTuning()
	switch strings.ToLower(playMachineName) {
	case "it":
		return it.Machine(tuning)
	case "xm":
		fallthrough
	default:
		return xm.Machine(tuning)
	}
}

func playParseGeneratorName(name string) (wave.Generator, error) {
	switch strings.ToLower(name) {
	case "sine", "sin":
		return sine.Sine, nil
	case "square", "sqr":
		return square.Square, nil
	case "triangle", "tri":
		return triangle.Triangle, nil
	default:
		return nil, fmt.Errorf("unsupported waveform name (%q)", name)
	}
}

func playParseADSR(params string) (envelope.ADSR, error) {
	var env envelope.ADSR
	pieces := strings.SplitN(params, ":", 4)
	if len(pieces) < 4 {
		return env, errors.New("incorrect parameter count for envelope")
	}

	a, d, s, r := pieces[0], pieces[1], pieces[2], pieces[3]

	if dur, err := time.ParseDuration(strings.TrimSpace(a)); err != nil {
		return env, err
	} else {
		env.Attack = dur
	}

	if dur, err := time.ParseDuration(strings.TrimSpace(d)); err != nil {
		return env, err
	} else {
		env.Decay = dur
	}

	if vol, err := playParseVol(s); err != nil {
		return env, err
	} else {
		env.Sustain = vol
	}

	if dur, err := time.ParseDuration(strings.TrimSpace(r)); err != nil {
		return env, err
	} else {
		env.Release = dur
	}

	return env, nil
}

func playParseVol(s string) (float32, error) {
	if ss := strings.ToLower(strings.TrimSpace(s)); strings.HasSuffix(ss, "db") {
		dbv, err := strconv.ParseFloat(strings.TrimSuffix(ss, "db"), 64)
		if err != nil {
			return 0, err
		}
		return db(dbv), nil
	}

	vol, err := strconv.ParseFloat(strings.TrimSpace(s), 32)
	return float32(vol), err
}

func playParseComponent(s string) (wavetable.Component, error) {
	var comp wavetable.Component
	tokens := strings.Split(s, ",")
	if len(tokens) < 1 {
		return comp, errors.New("must specify wave generator name as first parameter")
	}

	waveName := strings.TrimSpace(tokens[0])
	if gen, err := playParseGeneratorName(waveName); err != nil {
		return comp, err
	} else {
		comp.Generator = gen
	}

	for _, token := range tokens[1:] {
		pieces := strings.SplitN(token, ":", 2)
		name, value := strings.TrimSpace(pieces[0]), strings.TrimSpace(pieces[1])

		switch ln := strings.ToLower(name); ln {
		case "adsr":
			if env, err := playParseADSR(value); err != nil {
				return comp, err
			} else {
				comp.Envelope = env
			}
		case "amplitude", "amp", "a":
			if vol, err := playParseVol(value); err != nil {
				return comp, err
			} else {
				comp.Options = append(comp.Options, wave.SetParameterByName("amplitude", vol))
			}

		case "frequency", "freq", "f":
			if freq, err := strconv.ParseFloat(value, 64); err != nil {
				return comp, err
			} else {
				comp.Options = append(comp.Options, wave.SetParameterByName("frequency", freq))
			}

		default:
			comp.Options = append(comp.Options, wave.SetParameterByName(name, value))
		}
	}

	return comp, nil
}

func playGetTuning() tuning.Tuning {
	switch strings.ToLower(playTuningName) {
	case "equal-a415", "a415":
		return equalTuning.A415
	case "equal-a427", "a427":
		return equalTuning.A427
	case "equal-a428", "a428":
		return equalTuning.A428
	case "equal-a429", "a429":
		return equalTuning.A429
	case "equal-a430", "a430":
		return equalTuning.A430
	case "equal-a432", "a432":
		return equalTuning.A432
	case "equal-a435", "a435":
		return equalTuning.A435
	case "equal-a440", "a440":
		return equalTuning.A440
	case "equal-a444", "a444":
		return equalTuning.A444
	case "equal-a466", "a466":
		return equalTuning.A466
	case "equal-scientific", "scientific":
		return equalTuning.Scientific
	case "equal-53", "53tet", "53":
		return equalTuning.FiftyThree

	case "just-harmonic", "harmonic":
		return justTuning.Harmonic
	case "just-pythagorean", "pythagorean":
		return justTuning.Pythagorean

	default:
		return nil
	}
}

func playSynth(mach machine.Machine, onTick synth.SynthTickFunc, showHelp bool) error {
	var op wavetable.Operator
	if am, err := playParseComponent(playAM); err != nil {
		return err
	} else {
		op.Amplitude = am
	}

	if op.Amplitude.Generator == nil {
		return errors.New("must specify a valid amplitude modulator")
	}

	if playFM != "" {
		if fm, err := playParseComponent(playFM); err != nil {
			return err
		} else {
			op.Frequency = fm
		}
	}

	wt, err := wavetable.Generate(mach, op)
	if err != nil {
		panic(err)
	}
	voices := keyboard.NewKeyboard(mach, keymap.Default(mach.Tuning()), wt)

	out, err := output.CreateOutputDevice(deviceCommon.Settings{
		Name:             output.DefaultOutputDeviceName,
		Channels:         playOutput.Channels,
		SamplesPerSecond: playOutput.SamplesPerSecond,
		BitsPerSample:    playOutput.BitsPerSample,
	})
	if err != nil {
		panic(err)
	}

	defer out.Close()

	s := synth.NewSynth(playTickLength, voices, playNumPremixBuffers, float64(playOutput.SamplesPerSecond), 0.25)
	defer s.Close()

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer cancel()
		if err := s.Run(onTick); err != nil && !errors.Is(err, errTrackerDone) {
			panic(err)
		}
	}()

	go func() {
		defer cancel()
		if err := out.PlayWithCtx(ctx, s.C()); err != nil && !errors.Is(err, context.Canceled) {
			panic(err)
		}
	}()

	if err := ebiten.RunGame(ebikey.NewKeyboard(ctx, s, nil, showHelp)); err != nil && !errors.Is(err, ebikey.ErrWantStop) {
		panic(err)
	}

	fmt.Printf("closing...\n")
	return nil
}

func db(amt float64) float32 {
	return float32(math.Pow(10.0, amt/20.0))
}

func init() {
	output.Setup()
}
