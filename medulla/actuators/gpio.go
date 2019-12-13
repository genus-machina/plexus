package actuators

import (
	"github.com/genus-machina/plexus/medulla"
	"periph.io/x/periph/conn/gpio"
)

type gpioActuator struct {
	active   bool
	halted   bool
	inverted bool
	name     string
	pin      gpio.PinOut
}

func (device *gpioActuator) Activate() error {
	var err error
	value := gpio.High

	if device.inverted {
		value = gpio.Low
	}

	if err = device.pin.Out(value); err == nil {
		device.active = true
	}
	return err
}

func (device *gpioActuator) Deactivate() error {
	var err error
	value := gpio.Low

	if device.inverted {
		value = gpio.High
	}

	if err = device.pin.Out(value); err == nil {
		device.active = false
	}
	return err
}

func (device *gpioActuator) Halt() error {
	var err error
	if err = device.pin.Halt(); err == nil {
		device.halted = true
	}
	return err
}

func (device *gpioActuator) Name() string {
	return device.name
}

func (device *gpioActuator) State() medulla.DeviceState {
	return medulla.NewDeviceState(device.active, device.halted)
}
