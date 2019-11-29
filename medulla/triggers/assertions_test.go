package triggers

import (
	"strings"
	"testing"

	"github.com/genus-machina/plexus/medulla"
	"periph.io/x/periph/conn/gpio"
)

func assertActivate(t *testing.T, device medulla.Actuator) {
	if err := device.Activate(); err != nil {
		t.Errorf("activation error: %s", err.Error())
	}
}

func assertButton(t *testing.T, name string, pin gpio.PinIO) *Button {
	if button, err := NewButton(name, pin); err == nil {
		return button
	} else {
		t.Errorf("failed to create button '%s': %s", name, err.Error())
		return nil
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

func assertHalt(t *testing.T, device medulla.Device) {
	if err := device.Halt(); err != nil {
		t.Errorf("halt error: %s", err.Error())
	}
}

func assertIsActive(t *testing.T, device medulla.Device) {
	if !device.State().IsActive() {
		t.Errorf("expected device '%s' to be active", device.Name())
	}
}

func assertIsHalted(t *testing.T, device medulla.Device) {
	if !device.State().IsHalted() {
		t.Errorf("device '%s' failed to halt", device.Name())
	}
}

func assertIsNotHalted(t *testing.T, device medulla.Device) {
	if device.State().IsHalted() {
		t.Errorf("device '%s' is halted", device.Name())
	}
}

func assertIsInactive(t *testing.T, device medulla.Device) {
	if device.State().IsActive() {
		t.Errorf("expected device '%s' to be inactive", device.Name())
	}
}

func assertName(t *testing.T, device medulla.Device, expected string) {
	if name := device.Name(); name != expected {
		t.Errorf("expected name '%s' but got '%s'", expected, name)
	}
}

func assertPhototransistor(t *testing.T, name string, pin gpio.PinIO) *Phototransistor {
	if phototransistor, err := NewPhototransistor(name, pin); err == nil {
		return phototransistor
	} else {
		t.Errorf("failed to create phototransistor '%s': %s", name, err.Error())
		return nil
	}
}

func assertPIR(t *testing.T, name string, pin gpio.PinIO) *PIR {
	if pir, err := NewPIR(name, pin); err == nil {
		return pir
	} else {
		t.Errorf("failed to create PIR '%s': %s", name, err.Error())
		return nil
	}
}

func assertPullDown(t *testing.T, pin gpio.PinIn) {
	if pin.Pull() != gpio.PullDown {
		t.Error("expected pin to be pulled down")
	}
}

func assertPullUp(t *testing.T, pin gpio.PinIn) {
	if pin.Pull() != gpio.PullUp {
		t.Error("expected pin to be pulled up")
	}
}

func assertStates(t *testing.T, expected []medulla.DeviceState, actual []medulla.DeviceState) {
	actualCount := len(actual)
	expectedCount := len(expected)

	if actualCount != expectedCount {
		t.Errorf("expected %d states but got %d", expectedCount, actualCount)
	}

	for i := range expected {
		if actual[i].IsActive() != expected[i].IsActive() {
			t.Errorf("expected state %d IsActive to be %t", i, expected[i].IsActive())
		}
		if actual[i].IsHalted() != expected[i].IsHalted() {
			t.Errorf("expected state %d IsHalted to be %t", i, expected[i].IsHalted())
		}
	}
}

func assertSubscribe(t *testing.T, device medulla.Trigger) <-chan medulla.DeviceState {
	channel, err := device.Subscribe()
	if err != nil {
		t.Errorf("subscription error: %s", err.Error())
	}
	return channel
}
