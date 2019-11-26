package actuators

import (
	"strings"
	"testing"

	"github.com/genus-machina/plexus/medulla"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpiotest"
)

func assertActivate(t *testing.T, device medulla.Actuator) {
	if err := device.Activate(); err != nil {
		t.Errorf("activation error: %s", err.Error())
	}
}

func assertActive(t *testing.T, device medulla.Actuator) {
	if !device.State().IsActive() {
		t.Errorf("expected device '%s' to be active", device.Name())
	}
}

func assertDeactivate(t *testing.T, device medulla.Actuator) {
	if err := device.Deactivate(); err != nil {
		t.Errorf("deactivation error: %s", err.Error())
	}
}

func assertError(t *testing.T, err error, expected string) {
	if message := err.Error(); !strings.Contains(message, expected) {
		t.Errorf("unexpected error: %s", message)
	}
}

func assertGpioLevel(t *testing.T, pin *gpiotest.Pin, expected gpio.Level) {
	pin.Lock()
	defer pin.Unlock()

	if pin.L != expected {
		t.Errorf("expected pin to have level: %v", expected)
	}
}

func assertHalt(t *testing.T, device medulla.Device) {
	if err := device.Halt(); err != nil {
		t.Errorf("halt error: %s", err.Error())
	}
}

func assertHalted(t *testing.T, device medulla.Device) {
	if !device.State().IsHalted() {
		t.Errorf("expected device '%s' to be halted", device.Name())
	}
}

func assertInactive(t *testing.T, device medulla.Device) {
	if device.State().IsActive() {
		t.Errorf("expected device '%s' to be inactive", device.Name())
	}
}

func assertName(t *testing.T, device medulla.Device, expected string) {
	if name := device.Name(); name != expected {
		t.Errorf("expected '%s' but got '%s'", expected, name)
	}
}

func assertNotHalted(t *testing.T, device medulla.Device) {
	if device.State().IsHalted() {
		t.Errorf("expected device '%s' to not be halted", device.Name())
	}
}
