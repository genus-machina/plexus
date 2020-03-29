package medulla

import (
	"fmt"
	"time"
)

type Actuator interface {
	Activate() error
	Deactivate() error
	Device
}

type Device interface {
	Halt() error
	Name() string
	State() DeviceState
}

type DeviceState interface {
	IsActive() bool
	IsHalted() bool
	Time() time.Time
}

type Trigger interface {
	Device
	Subscribe() (<-chan DeviceState, error)
}

func CheckDeviceState(device Device) error {
	if device.State().IsHalted() {
		return fmt.Errorf("Device '%s' has halted.", device.Name())
	}
	return nil
}
