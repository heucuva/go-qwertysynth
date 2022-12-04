package wave

type Wave interface {
	Data() []float32
	SampleRate() float64
}
