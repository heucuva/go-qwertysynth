package keyoctave

type KeyOctave int

const (
	TotalKeyOctaves = KeysPerOctave * NumOctaves
)

func NewKeyOctave(k Key, o Octave) KeyOctave {
	ko := KeyOctave(o*Octave(KeysPerOctave)) + KeyOctave(k)
	return ko
}

func (ko KeyOctave) Split() (Key, Octave) {
	k := Key(int(ko) % KeysPerOctave)
	o := Octave(int(ko) / KeysPerOctave)
	return k, o
}
