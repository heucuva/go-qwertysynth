package scale

type Key int

const (
	KeyC = Key(0 + iota)
	KeyCSharp
	KeyD
	KeyDSharp
	KeyE
	KeyF
	KeyFSharp
	KeyG
	KeyGSharp
	KeyA
	KeyASharp
	KeyB
	cMaxKeys
	MaxKey = KeyB
	MinKey = KeyC
)

const keyStr = "C-C#D-D#E-F-F#G-G#A-A#B-"

func (k Key) String() string {
	switch {
	case k >= MinKey && k <= MaxKey:
		idx := (k - MinKey) * 2
		return keyStr[idx : idx+2]
	default:
		return "??"
	}
}
