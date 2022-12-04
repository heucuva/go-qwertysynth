package device

import (
	"context"

	deviceCommon "github.com/heucuva/go-qwertysynth/internal/output/device/common"
	"github.com/heucuva/go-qwertysynth/internal/output/premix"
	"github.com/pkg/errors"
)

var (
	// ErrDeviceNotSupported is returned when the requested device is not supported
	ErrDeviceNotSupported = errors.New("device not supported")
)

// Device is an interface to output device operations
type Device interface {
	Name() string
	Play(in <-chan *premix.PremixData) error
	PlayWithCtx(ctx context.Context, in <-chan *premix.PremixData) error
	Close() error
}

type kindGetter interface {
	GetKind() deviceCommon.Kind
}

type createOutputDeviceFunc func(settings deviceCommon.Settings) (Device, error)

type deviceDetails struct {
	create createOutputDeviceFunc
	Kind   deviceCommon.Kind
}

// GetKind returns the kind for the passed in device
func GetKind(d Device) deviceCommon.Kind {
	if dev, ok := d.(kindGetter); ok {
		return dev.GetKind()
	}
	return deviceCommon.KindNone
}

var (
	// Map is the mapping of device name to device details
	Map = make(map[string]deviceDetails)
)

// CreateOutputDevice creates an output device based on the provided settings
func CreateOutputDevice(settings deviceCommon.Settings) (Device, error) {
	if details, ok := Map[settings.Name]; ok && details.create != nil {
		dev, err := details.create(settings)
		if err != nil {
			return nil, err
		}
		return dev, nil
	}

	return nil, errors.Wrap(ErrDeviceNotSupported, settings.Name)
}

type device struct {
	onRowOutput deviceCommon.WrittenCallback
}
