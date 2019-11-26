package actuators

import (
	"github.com/genus-machina/plexus/medulla"
	"periph.io/x/periph/conn/gpio"
)

type LED struct {
	active bool
	halted bool
	name   string
	pin    gpio.PinOut
}

func NewLED(name string, pin gpio.PinOut) *LED {
	device := new(LED)
	device.name = name
	device.pin = pin

	device.pin.Out(gpio.Low)
	device.active = false
	device.halted = false
	return device
}

func (device *LED) Activate() error {
	var err error
	if err = device.pin.Out(gpio.High); err == nil {
		device.active = true
	}
	return err
}

func (device *LED) Deactivate() error {
	var err error
	if err = device.pin.Out(gpio.Low); err == nil {
		device.active = false
	}
	return err
}

func (device *LED) Halt() error {
	var err error
	if err = device.pin.Halt(); err == nil {
		device.halted = true
	}
	return err
}

func (device *LED) Name() string {
	return device.name
}

func (device *LED) State() medulla.DeviceState {
	return medulla.NewDeviceState(device.active, device.halted)
}
