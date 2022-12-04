package envelope

import "time"

type Envelope interface {
	KeyOn()
	KeyOff()
	Cut()
	IsPlaying() bool
	Get() (float32, bool)
	Advance(dur time.Duration)
}
