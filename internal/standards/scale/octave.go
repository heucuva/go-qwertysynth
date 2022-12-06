package scale

const (
	MinOctave  Octave = 0
	MaxOctave  Octave = 9
	NumOctaves        = int(MaxOctave - MinOctave + 1)
)

type Octave int
