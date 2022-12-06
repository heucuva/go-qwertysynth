package just

type HarmonicKey int

const (
	HarmonicKeyC = HarmonicKey(0 + iota)
	HarmonicKeyCSharp
	HarmonicKeyD
	HarmonicKeyDSharp
	HarmonicKeyE
	HarmonicKeyF
	HarmonicKeyFSharp
	HarmonicKeyG
	HarmonicKeyGSharp
	HarmonicKeyA
	HarmonicKeyASharp
	HarmonicKeyB
	cHarmonicMaxKeys
	HarmonicMaxKey = HarmonicKeyB
	HarmonicMinKey = HarmonicKeyC
)

const harmonicKeyStr = "C-C#D-D#E-F-F#G-G#A-A#B-"

const HarmonicKeysPerOctave int = int(cHarmonicMaxKeys)

func (k HarmonicKey) String() string {
	switch {
	case k >= HarmonicMinKey && k <= HarmonicMaxKey:
		idx := k.Index() * 2
		return harmonicKeyStr[idx : idx+2]
	default:
		return "??"
	}
}

func (HarmonicKey) KeysPerOctave() int {
	return HarmonicKeysPerOctave
}

func (k HarmonicKey) Index() int {
	return int(k)
}
