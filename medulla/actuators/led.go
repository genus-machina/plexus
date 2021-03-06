package actuators

import (
	"periph.io/x/periph/conn/gpio"
)

type LED struct {
	gpioActuator
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
