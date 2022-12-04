package wave

type Generator func(opts ...GeneratorParam) (Wave, error)
