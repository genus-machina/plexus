package synapse

import (
	"time"

	"github.com/genus-machina/plexus/medulla"
)

type TestDevice struct{}

func NewTestDevice() *TestDevice {
	return new(TestDevice)
}

func (device *TestDevice) Halt() error {
	return nil
}

func (device *TestDevice) Name() string {
	return "test"
}

func (device *TestDevice) State() medulla.DeviceState {
	return medulla.NewDeviceState(false, false, time.Unix(0, 0))
}
