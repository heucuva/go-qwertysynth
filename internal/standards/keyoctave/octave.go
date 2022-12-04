package keyoctave

const (
	MinOctave     Octave = 0
	MaxOctave     Octave = 9
	NumOctaves           = int(MaxOctave - MinOctave + 1)
	KeysPerOctave        = int(cMaxKeys)
)

type Octave int
