package command

import (
	"errors"
	"fmt"
	"time"

	"github.com/heucuva/go-qwertysynth/internal/machine"
	"github.com/heucuva/go-qwertysynth/internal/output"
	"github.com/heucuva/go-qwertysynth/internal/standards/keyoctave"
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/synth"
	"github.com/spf13/cobra"
)

// flags
var (
	playTrackerBpm  int  = 300
	playTrackerLoop bool = false
)

func init() {
	output.Setup()

	cflags := playTrackerCmd.Flags()
	cflags.IntVarP(&playTrackerBpm, "bpm", "B", playTrackerBpm, "beats per minute")
	cflags.BoolVarP(&playTrackerLoop, "loop", "l", playTrackerLoop, "loop song")

	playCmd.AddCommand(playTrackerCmd)
}

var playTrackerCmd = &cobra.Command{
	Use:   "tracker [flags]",
	Short: "Play synth sounds using a tracker",
	Long:  "Play synth sounds using a tracker.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		mach := playGetMachine()
		chopsticks := trackerGenPattern(mach)

		t := newTracker(playTrackerBpm, playTrackerLoop, chopsticks)

		return playSynth(mach, t.Tick, false)
	},
}

type column struct {
	n note.Note
	a synth.KeyAction
}

type row []column

func trackerGenPattern(mach machine.Machine) []row {
	var (
		g2 = mach.Note(2, keyoctave.KeyG, 0)
		d3 = mach.Note(3, keyoctave.KeyD, 0)
		c3 = mach.Note(3, keyoctave.KeyC, 0)
		//c4 = mach.Note(4, keyoctave.KeyC, 0)
		d4 = mach.Note(4, keyoctave.KeyD, 0)
		e4 = mach.Note(4, keyoctave.KeyE, 0)
		f4 = mach.Note(4, keyoctave.KeyF, 0)
		g4 = mach.Note(4, keyoctave.KeyG, 0)
		a4 = mach.Note(4, keyoctave.KeyA, 0)
		b4 = mach.Note(4, keyoctave.KeyB, 0)
		c5 = mach.Note(5, keyoctave.KeyC, 0)
	)

	var chopsticks []row

	addNotes := func(a synth.KeyAction, notes ...note.Note) {
		var r row
		for _, n := range notes {
			r = append(r, column{
				n: n,
				a: a,
			})
		}
		chopsticks = append(chopsticks, r)
	}

	addNotes(synth.KeyActionOn, f4, g4, d3)
	addNotes(synth.KeyActionOff, f4, g4)
	addNotes(synth.KeyActionOn, f4, g4)
	addNotes(synth.KeyActionOff, f4, g4)
	addNotes(synth.KeyActionOn, f4, g4)
	addNotes(synth.KeyActionOff, f4, g4, d3)

	addNotes(synth.KeyActionOn, f4, g4, g2)
	addNotes(synth.KeyActionOff, f4, g4)
	addNotes(synth.KeyActionOn, f4, g4)
	addNotes(synth.KeyActionOff, f4, g4)
	addNotes(synth.KeyActionOn, f4, g4)
	addNotes(synth.KeyActionOff, f4, g4, g2)

	addNotes(synth.KeyActionOn, e4, g4, c3)
	addNotes(synth.KeyActionOff, e4, g4)
	addNotes(synth.KeyActionOn, e4, g4)
	addNotes(synth.KeyActionOff, e4, g4)
	addNotes(synth.KeyActionOn, e4, g4)
	addNotes(synth.KeyActionOff, e4, g4, c3)

	addNotes(synth.KeyActionOn, e4, g4, g2)
	addNotes(synth.KeyActionOff, e4, g4)
	addNotes(synth.KeyActionOn, e4, g4)
	addNotes(synth.KeyActionOff, e4, g4)
	addNotes(synth.KeyActionOn, e4, g4)
	addNotes(synth.KeyActionOff, e4, g4, g2)
	//l2
	addNotes(synth.KeyActionOn, f4, b4, d3)
	addNotes(synth.KeyActionOff, f4, b4)
	addNotes(synth.KeyActionOn, f4, b4)
	addNotes(synth.KeyActionOff, f4, b4)
	addNotes(synth.KeyActionOn, f4, b4)
	addNotes(synth.KeyActionOff, f4, b4, d3)

	addNotes(synth.KeyActionOn, f4, b4, g2)
	addNotes(synth.KeyActionOff, f4, b4)
	addNotes(synth.KeyActionOn, f4, a4)
	addNotes(synth.KeyActionOff, f4, a4)
	addNotes(synth.KeyActionOn, f4, b4)
	addNotes(synth.KeyActionOff, f4, b4, g2)

	addNotes(synth.KeyActionOn, e4, g4, c3)
	addNotes(synth.KeyActionOff, e4, g4)
	addNotes(synth.KeyActionOn, e4, g4)
	addNotes(synth.KeyActionOff, e4, g4)
	addNotes(synth.KeyActionOn, e4, g4)
	addNotes(synth.KeyActionOff, e4, g4, c3)

	addNotes(synth.KeyActionOn, e4, c5, g2)
	addNotes(synth.KeyActionOff, e4, c5)
	addNotes(synth.KeyActionOn, e4, c5)
	addNotes(synth.KeyActionOff, e4, c5)
	addNotes(synth.KeyActionOn, e4, c5)
	addNotes(synth.KeyActionOff, e4, c5, g2)
	//l3
	addNotes(synth.KeyActionOn, f4, g4, d3)
	addNotes(synth.KeyActionOff, f4, g4)
	addNotes(synth.KeyActionOn, f4, g4)
	addNotes(synth.KeyActionOff, f4, g4)
	addNotes(synth.KeyActionOn, f4, g4)
	addNotes(synth.KeyActionOff, f4, g4, d3)

	addNotes(synth.KeyActionOn, f4, g4, g2)
	addNotes(synth.KeyActionOff, f4, g4)
	addNotes(synth.KeyActionOn, f4, g4)
	addNotes(synth.KeyActionOff, f4, g4)
	addNotes(synth.KeyActionOn, f4, g4)
	addNotes(synth.KeyActionOff, f4, g4, g2)

	addNotes(synth.KeyActionOn, e4, g4, c3)
	addNotes(synth.KeyActionOff, e4, g4)
	addNotes(synth.KeyActionOn, e4, g4)
	addNotes(synth.KeyActionOff, e4, g4)
	addNotes(synth.KeyActionOn, e4, g4)
	addNotes(synth.KeyActionOff, e4, g4, c3)

	addNotes(synth.KeyActionOn, e4, g4, g2)
	addNotes(synth.KeyActionOff, e4, g4)
	addNotes(synth.KeyActionOn, e4, g4)
	addNotes(synth.KeyActionOff, e4, g4)
	addNotes(synth.KeyActionOn, e4, g4)
	addNotes(synth.KeyActionOff, e4, g4, g2)
	//l4
	addNotes(synth.KeyActionOn, f4, b4, d3)
	addNotes(synth.KeyActionOff, f4, b4)
	addNotes(synth.KeyActionOn, f4, b4)
	addNotes(synth.KeyActionOff, f4, b4)
	addNotes(synth.KeyActionOn, f4, b4)
	addNotes(synth.KeyActionOff, f4, b4, d3)

	addNotes(synth.KeyActionOn, f4, b4, g2)
	addNotes(synth.KeyActionOff, f4, b4)
	addNotes(synth.KeyActionOn, f4, a4)
	addNotes(synth.KeyActionOff, f4, a4)
	addNotes(synth.KeyActionOn, f4, b4)
	addNotes(synth.KeyActionOff, f4, b4, g2)

	addNotes(synth.KeyActionOn, e4, c5, c3)
	addNotes(synth.KeyActionOff, e4, c5, c3)
	addNotes(synth.KeyActionCut, e4, c5, c3)
	addNotes(synth.KeyActionFadeout)
	addNotes(synth.KeyActionOn, f4, b4, g2)
	addNotes(synth.KeyActionOff, f4, b4, g2)

	addNotes(synth.KeyActionOn, e4, c5, c3)
	addNotes(synth.KeyActionOff, e4, c5, c3)
	addNotes(synth.KeyActionCut, e4, c5, c3)
	addNotes(synth.KeyActionFadeout)
	addNotes(synth.KeyActionOn, b4, d4)
	addNotes(synth.KeyActionOff, b4, d4)

	return chopsticks[:]
}

type tracker struct {
	pattern []row
	r       int
	rowDur  time.Duration
	cur     time.Duration
	loop    bool
}

func newTracker(bpm int, loop bool, pattern []row) *tracker {
	rowDur := time.Minute / time.Duration(bpm)
	return &tracker{
		pattern: pattern,
		rowDur:  rowDur,
		r:       -1,
		loop:    loop,
	}
}

func (t *tracker) Tick(s synth.Synth, dur time.Duration) error {
	t.cur += dur
	for t.cur >= t.rowDur {
		t.cur -= t.rowDur
		if err := t.advanceRow(s); err != nil {
			return err
		}
	}
	return nil
}

var (
	errTrackerDone = errors.New("done")
)

func (t *tracker) advanceRow(s synth.Synth) error {
	t.r++
	if t.r >= len(t.pattern) {
		if !t.loop {
			return errTrackerDone
		} else {
			for t.r >= len(t.pattern) {
				t.r -= len(t.pattern)
			}
		}
	}

	printedOne := false
	for _, c := range t.pattern[t.r] {
		if c.a == synth.KeyActionOn {
			if printedOne {
				fmt.Print("\t")
			}
			fmt.Printf("%v", c.n)
			printedOne = true
		}
		s.KeyAction(c.n, c.a)
	}
	if printedOne {
		fmt.Println()
	}

	return nil
}
