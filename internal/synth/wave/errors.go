package wave

import "errors"

var (
	ErrInvalidParameterValue    = errors.New("invalid parameter value")
	ErrNotValidForThisGenerator = errors.New("not valid for this generator")
)
