package actuators

import (
	"testing"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpiotest"
)

func TestLEDName(t *testing.T) {
	name := "test"
	pin := new(gpiotest.Pin)
	led := NewLED(name, pin)
	assertName(t, led, name)
}

func TestLEDActivate(t *testing.T) {
	name := "test"
	pin := new(gpiotest.Pin)
	led := NewLED(name, pin)

	assertInactive(t, led)
	assertGpioLevel(t, pin, gpio.Low)
	assertActivate(t, led)
	assertActive(t, led)
	assertGpioLevel(t, pin, gpio.High)
}

func TestLEDDeactivate(t *testing.T) {
	name := "test"
	pin := new(gpiotest.Pin)
	led := NewLED(name, pin)

	assertActivate(t, led)
	assertGpioLevel(t, pin, gpio.High)
	assertDeactivate(t, led)
	assertInactive(t, led)
	assertGpioLevel(t, pin, gpio.Low)
}

func TestLEDHalt(t *testing.T) {
	name := "test"
	pin := new(gpiotest.Pin)
	led := NewLED(name, pin)

	assertNotHalted(t, led)
	assertHalt(t, led)
	assertHalted(t, led)
}
