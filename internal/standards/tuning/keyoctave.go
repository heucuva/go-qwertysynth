package tuning

import "github.com/heucuva/go-qwertysynth/internal/standards/scale"

type KeyOctave int

func NewKeyOctave(k scale.Key, o scale.Octave) KeyOctave {
	ko := KeyOctave(o<<8) | KeyOctave(k.Index())
	return ko
}

func (ko KeyOctave) Split(tuning Tuning) (scale.Key, scale.Octave) {
	k := tuning.Key(int(ko) & 255)
	o := scale.Octave(int(ko) >> 8)
	return k, o
}
