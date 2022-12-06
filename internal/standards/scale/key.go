package scale

import "fmt"

type Key interface {
	fmt.Stringer
	KeysPerOctave() int
	Index() int
}
