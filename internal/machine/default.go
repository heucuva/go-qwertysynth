package machine

import (
	"github.com/heucuva/go-qwertysynth/internal/standards/note"
	"github.com/heucuva/go-qwertysynth/internal/standards/tuning"
)

type Default interface {
	WaveformFrequency() float64
	CenterNote() note.Note
	Tuning() tuning.Tuning
}
