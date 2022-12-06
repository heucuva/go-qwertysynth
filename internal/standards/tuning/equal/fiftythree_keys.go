package equal

type FiftyThreeKey int

const (
	FiftyThreeKeyC = FiftyThreeKey(0 + iota)
	FiftyThreeKeyCSharp
	FiftyThreeKeyCDoubleSharp
	FiftyThreeKeyCSharpDoubleSharp
	FiftyThreeKeyCDoubleSharpDoubleSharp
	FiftyThreeKeyDDoubleFlatDoubleFlat
	FiftyThreeKeyDTripleFlat
	FiftyThreeKeyDDoubleFlat
	FiftyThreeKeyDFlat
	FiftyThreeKeyD
	FiftyThreeKeyDSharp
	FiftyThreeKeyDDoubleSharp
	FiftyThreeKeyDSharpDoubleSharp
	FiftyThreeKeyDDoubleSharpDoubleSharp
	FiftyThreeKeyEDoubleFlatDoubleFlat
	FiftyThreeKeyETripleFlat
	FiftyThreeKeyEDoubleFlat
	FiftyThreeKeyEFlat
	FiftyThreeKeyE
	FiftyThreeKeyESharp
	FiftyThreeKeyEDoubleSharp
	FiftyThreeKeyFFlat
	FiftyThreeKeyF
	FiftyThreeKeyFSharp
	FiftyThreeKeyFDoubleSharp
	FiftyThreeKeyFSharpDoubleSharp
	FiftyThreeKeyFDoubleSharpDoubleSharp
	FiftyThreeKeyGDoubleFlatDoubleFlat
	FiftyThreeKeyGTripleFlat
	FiftyThreeKeyGDoubleFlat
	FiftyThreeKeyGFlat
	FiftyThreeKeyG
	FiftyThreeKeyGSharp
	FiftyThreeKeyGDoubleSharp
	FiftyThreeKeyGSharpDoubleSharp
	FiftyThreeKeyGDoubleSharpDoubleSharp
	FiftyThreeKeyADoubleFlatDoubleFlat
	FiftyThreeKeyATripleFlat
	FiftyThreeKeyADoubleFlat
	FiftyThreeKeyAFlat
	FiftyThreeKeyA
	FiftyThreeKeyASharp
	FiftyThreeKeyADoubleSharp
	FiftyThreeKeyASharpDoubleSharp
	FiftyThreeKeyADoubleSharpDoubleSharp
	FiftyThreeKeyBDoubleFlatDoubleFlat
	FiftyThreeKeyBTripleFlat
	FiftyThreeKeyBDoubleFlat
	FiftyThreeKeyBFlat
	FiftyThreeKeyB
	FiftyThreeKeyBSharp
	FiftyThreeKeyBDoubleSharp
	FiftyThreeKeyCFlat
	cMaxFiftyThreeKeys
	MaxFiftyThreeKey = FiftyThreeKeyCFlat
	MinFiftyThreeKey = FiftyThreeKeyC
)

const fiftyThreeKeyStr = "C-C#C*C+C`D_DvDBDb" +
	"D-D#D*D+D`E_EvEBEb" +
	"E-E#E*Fb" +
	"F-F#F*F+F`G_GvGBGb" +
	"G-G#G*G+G`A_AvABAb" +
	"A-A#A*A+A`B_BvBBBb" +
	"B-B#B*Cb"

const FiftyThreeKeysPerOctave int = int(cMaxFiftyThreeKeys)

func (k FiftyThreeKey) String() string {
	switch {
	case k >= MinFiftyThreeKey && k <= MaxFiftyThreeKey:
		idx := k.Index() * 2
		return fiftyThreeKeyStr[idx : idx+2]
	default:
		return "??"
	}
}

func (FiftyThreeKey) KeysPerOctave() int {
	return FiftyThreeKeysPerOctave
}

func (k FiftyThreeKey) Index() int {
	return int(k)
}
