package output

import (
	"errors"

	"github.com/heucuva/go-qwertysynth/internal/output/device"
	deviceCommon "github.com/heucuva/go-qwertysynth/internal/output/device/common"
)

type devicePriority int

// the further down the list, the higher the priority
const (
	devicePriorityNone = devicePriority(iota)
	devicePriorityPulseAudio
	devicePriorityWinmm
	devicePriorityDirectSound
)

var (
	// DefaultOutputDeviceName is the default device name
	DefaultOutputDeviceName = "none"

	devicePriorityMap = make(map[string]devicePriority)
)

func calculateOptimalDefaultOutputDeviceName() string {
	preferredPriority := devicePriority(0)
	preferredName := "none"
	for name := range device.Map {
		if priority, ok := devicePriorityMap[name]; ok && priority > preferredPriority {
			preferredName = name
			preferredPriority = priority
		}
	}

	return preferredName
}

// CreateOutputDevice creates an output device based on the provided settings
func CreateOutputDevice(settings deviceCommon.Settings) (device.Device, error) {
	d, err := device.CreateOutputDevice(settings)
	if err != nil {
		return nil, err
	}

	if d == nil {
		return nil, errors.New("could not create output device")
	}

	return d, nil
}

// Setup finalizes the output device preference system
func Setup() {
	DefaultOutputDeviceName = calculateOptimalDefaultOutputDeviceName()
}

// DeviceInfo returns information about a device
type DeviceInfo struct {
	Priority int
	Kind     deviceCommon.Kind
}

func GetOutputDevices() map[string]DeviceInfo {
	m := make(map[string]DeviceInfo)
	for k, v := range devicePriorityMap {
		if d, ok := device.Map[k]; ok {
			m[k] = DeviceInfo{
				Priority: int(v),
				Kind:     d.Kind,
			}
		}
	}
	return m
}

func init() {
	_ = devicePriorityNone // lint
	devicePriorityMap["pulseaudio"] = devicePriorityPulseAudio
	devicePriorityMap["winmm"] = devicePriorityWinmm
	devicePriorityMap["directsound"] = devicePriorityDirectSound
}
