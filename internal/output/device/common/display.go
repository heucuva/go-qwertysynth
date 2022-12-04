package common

import "github.com/heucuva/go-qwertysynth/internal/output/premix"

// WrittenCallback defines the callback for when a premix buffer is mixed/rendered and output on the device
type WrittenCallback func(deviceCommon Kind, premix *premix.PremixData)
