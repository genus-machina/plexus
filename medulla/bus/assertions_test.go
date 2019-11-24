package bus

import (
	"strings"
	"testing"

	"github.com/genus-machina/plexus/medulla"
)

func assertActivate(t *testing.T, bus *DeviceBus, name string) {
	if err := bus.Activate(name); err != nil {
		t.Errorf("activation failed: %s", err.Error())
	}
}

func assertActive(t *testing.T, device medulla.Device) {
	if !device.State().IsActive() {
		t.Errorf("expected device '%s' to be active", device.Name())
	}
}

func assertDeactivate(t *testing.T, bus *DeviceBus, name string) {
	if err := bus.Deactivate(name); err != nil {
		t.Errorf("deactivation error: %s", err.Error())
	}
}

func assertError(t *testing.T, err error, expected string) {
	if message := err.Error(); !strings.Contains(message, expected) {
		t.Errorf("unexpected error: %s", message)
	}
}

func assertHalt(t *testing.T, bus *DeviceBus) {
	if err := bus.Halt(); err != nil {
		t.Errorf("halting error: %s", err.Error())
	}
}

func assertHaltDevice(t *testing.T, device medulla.Device) {
	if err := device.Halt(); err != nil {
		t.Errorf("device '%s' failed to halt: %s", device.Name(), err.Error())
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

func assertRegister(t *testing.T, bus *DeviceBus, device medulla.Device) {
	if err := bus.RegisterDevice(device); err != nil {
		t.Errorf("registration error: %s", err.Error())
	}
}

func assertSubscribe(t *testing.T, bus *DeviceBus, name string) <-chan medulla.DeviceState {
	channel, err := bus.Subscribe(name)
	if err != nil {
		t.Errorf("subscription error: %s", err.Error())
	}
	return channel
}

func assertStates(t *testing.T, expected []medulla.DeviceState, actual []medulla.DeviceState) {
	actualCount := len(actual)
	expectedCount := len(expected)

	if actualCount != expectedCount {
		t.Errorf("expected %d states but got %d", expectedCount, actualCount)
	}

	for i := range expected {
		if actual[i].IsActive() != expected[i].IsActive() {
			t.Errorf("expected value %d IsActive to be %t", i, actual[i].IsActive())
		}
		if actual[i].IsHalted() != expected[i].IsHalted() {
			t.Errorf("expected value %d IsHalted to be %t", i, actual[i].IsHalted())
		}
	}
}
