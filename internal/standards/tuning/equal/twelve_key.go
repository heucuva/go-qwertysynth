package equal

type TwelveKey int

const (
	TwelveKeyC = TwelveKey(0 + iota)
	TwelveKeyCSharp
	TwelveKeyD
	TwelveKeyDSharp
	TwelveKeyE
	TwelveKeyF
	TwelveKeyFSharp
	TwelveKeyG
	TwelveKeyGSharp
	TwelveKeyA
	TwelveKeyASharp
	TwelveKeyB
	cMaxTwelveKeys
	MaxTwelveKey = TwelveKeyB
	MinTwelveKey = TwelveKeyC
)

const twelveKeyStr = "C-C#D-D#E-F-F#G-G#A-A#B-"

const TwelveKeysPerOctave int = int(cMaxTwelveKeys)

func (k TwelveKey) String() string {
	switch {
	case k >= MinTwelveKey && k <= MaxTwelveKey:
		idx := k.Index() * 2
		return twelveKeyStr[idx : idx+2]
	default:
		return "??"
	}
}

func (TwelveKey) KeysPerOctave() int {
	return TwelveKeysPerOctave
}

func (k TwelveKey) Index() int {
	return int(k)
}
