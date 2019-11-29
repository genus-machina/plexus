package actuators

import (
	"testing"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpiotest"
)

func TestLampName(t *testing.T) {
	name := "test"
	pin := new(gpiotest.Pin)
	lamp := NewLamp(name, pin)
	assertName(t, lamp, name)
}

func TestLampActivate(t *testing.T) {
	name := "test"
	pin := new(gpiotest.Pin)
	lamp := NewLamp(name, pin)

	assertInactive(t, lamp)
	assertGpioLevel(t, pin, gpio.High)
	assertActivate(t, lamp)
	assertActive(t, lamp)
	assertGpioLevel(t, pin, gpio.Low)
}

func TestLampDeactivate(t *testing.T) {
	name := "test"
	pin := new(gpiotest.Pin)
	lamp := NewLamp(name, pin)

	assertActivate(t, lamp)
	assertGpioLevel(t, pin, gpio.Low)
	assertDeactivate(t, lamp)
	assertInactive(t, lamp)
	assertGpioLevel(t, pin, gpio.High)
}

func TestLampHalt(t *testing.T) {
	name := "test"
	pin := new(gpiotest.Pin)
	lamp := NewLamp(name, pin)

	assertNotHalted(t, lamp)
	assertHalt(t, lamp)
	assertHalted(t, lamp)
}
