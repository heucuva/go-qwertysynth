package wave

type GeneratorParam func(o interface{}) error

type setParameterByName interface {
	SetParameterByName(name string, value any) error
}

func SetParameterByName(name string, value any) GeneratorParam {
	return func(o interface{}) error {
		iface, ok := o.(setParameterByName)
		if !ok {
			return ErrNotValidForThisGenerator
		}

		return iface.SetParameterByName(name, value)
	}
}
