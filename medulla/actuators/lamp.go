package actuators

import (
	"periph.io/x/periph/conn/gpio"
)

type Lamp struct {
	gpioActuator
}

func NewLamp(name string, pin gpio.PinOut) *Lamp {
	device := new(Lamp)
	device.inverted = true
	device.name = name
	device.pin = pin

	device.pin.Out(gpio.High)
	device.active = false
	device.halted = false
	return device
}
