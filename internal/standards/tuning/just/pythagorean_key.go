package just

type PythagoreanKey int

const (
	PythagoreanKeyD = PythagoreanKey(0 + iota)
	PythagoreanKeyEFlat
	PythagoreanKeyE
	PythagoreanKeyF
	PythagoreanKeyFSharp
	PythagoreanKeyG
	PythagoreanKeyAFlat
	PythagoreanKeyA
	PythagoreanKeyBFlat
	PythagoreanKeyB
	PythagoreanKeyC
	PythagoreanKeyCSharp
	cPythagoreanMaxKeys
	PythagoreanMaxKey = PythagoreanKeyCSharp
	PythagoreanMinKey = PythagoreanKeyD
)

const PythagoreanKeysPerOctave int = int(cPythagoreanMaxKeys)

const pythagoreankeyStr = "D-EbE-F-F#G-AbA-BbB-C-C#"

func (k PythagoreanKey) String() string {
	switch {
	case k >= PythagoreanMinKey && k <= PythagoreanMaxKey:
		idx := k.Index() * 2
		return pythagoreankeyStr[idx : idx+2]
	default:
		return "??"
	}
}

func (PythagoreanKey) KeysPerOctave() int {
	return PythagoreanKeysPerOctave
}

func (k PythagoreanKey) Index() int {
	return int(k)
}
