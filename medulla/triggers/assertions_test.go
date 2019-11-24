package triggers

import (
	"strings"
	"testing"

	"github.com/genus-machina/plexus/medulla"
)

func assertActivate(t *testing.T, device medulla.Actuator) {
	if err := device.Activate(); err != nil {
		t.Errorf("activation error: %s", err.Error())
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
