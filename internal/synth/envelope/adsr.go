package envelope

import "time"

type ADSR struct {
	Attack  time.Duration
	Decay   time.Duration
	Sustain float32
	Release time.Duration
}

type adsr struct {
	mode    adsrMode
	modeDur time.Duration

	settings ADSR

	startLevel float32
	level      float32
}

var _ Envelope = &adsr{}

type adsrMode int

const (
	adsrModeIdle adsrMode = iota
	adsrModeAttack
	adsrModeDecay
	adsrModeSustain
	adsrModeRelease
)

func NewADSR(settings ADSR) Envelope {
	return &adsr{
		mode:     adsrModeIdle,
		settings: settings,
	}
}

func (e *adsr) KeyOn() {
	e.mode = adsrModeAttack
	e.modeDur = 0
	e.startLevel = e.level
	e.Advance(0)
}

func (e *adsr) KeyOff() {
	switch e.mode {
	case adsrModeAttack, adsrModeDecay, adsrModeSustain:
		e.mode = adsrModeRelease
		e.modeDur = 0
		e.startLevel = e.level
		e.Advance(0)
	default:
		// do nothing
	}
}

func (e *adsr) Cut() {
	e.mode = adsrModeIdle
	e.modeDur = 0
	e.level = 0
	e.startLevel = e.level
}

func (e adsr) IsPlaying() bool {
	return e.mode != adsrModeIdle
}

func (e adsr) Get() (float32, bool) {
	return e.level, e.IsPlaying()
}

func (adsr) lerp(t float64, a, b float32) float32 {
	if t < 0.0 {
		t = 0.0
	} else if t > 1.0 {
		t = 1.0
	}

	difft := float64(b-a) * t
	return a + float32(difft)
}

func (e adsr) modeTimePct(limit time.Duration) float64 {
	if limit == 0 {
		return 1.0
	}

	return e.modeDur.Seconds() / limit.Seconds()
}

func (e adsr) lerpWithTime(limit time.Duration, a, b float32) float32 {
	return e.lerp(e.modeTimePct(limit), a, b)
}

func (e *adsr) Advance(dur time.Duration) {

	switch e.mode {
	case adsrModeAttack:
		e.modeDur += dur
		if e.modeDur < e.settings.Attack {
			e.level = e.lerpWithTime(e.settings.Attack, e.startLevel, 1.0)
			return
		}

		e.mode = adsrModeDecay
		e.modeDur -= e.settings.Attack
		e.startLevel = 1.0
		fallthrough

	case adsrModeDecay:
		e.modeDur += dur
		if e.modeDur < e.settings.Decay {
			e.level = e.lerpWithTime(e.settings.Decay, e.startLevel, e.settings.Sustain)
			return
		}

		e.mode = adsrModeSustain
		e.level = e.settings.Sustain
		fallthrough

	case adsrModeSustain:

	case adsrModeRelease:
		e.modeDur += dur
		if e.modeDur < e.settings.Release {
			e.level = e.lerpWithTime(e.settings.Release, e.startLevel, 0.0)
			return
		}

		e.mode = adsrModeIdle
		e.level = 0.0
		e.startLevel = e.level
		fallthrough

	default:
	}
}
