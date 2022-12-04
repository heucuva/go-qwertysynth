package machine

import "github.com/heucuva/go-qwertysynth/internal/standards/note"

type Default interface {
	WaveformFrequency() float64
	CenterNote() note.Note
}
