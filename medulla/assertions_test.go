package medulla

import (
	"strings"
	"testing"
)

func assertCheckDeviceState(t *testing.T, device Device) {
	if err := CheckDeviceState(device); err != nil {
		t.Errorf("device '%s' failed check: %s", device.Name(), err.Error())
	}
}

func assertError(t *testing.T, err error, expected string) {
	if message := err.Error(); !strings.Contains(message, expected) {
		t.Errorf("unexpected error: %s", message)
	}
}

func assertHalt(t *testing.T, device Device) {
	if err := device.Halt(); err != nil {
		t.Errorf("halt error: %s", err.Error())
	}
}

func assertIsActive(t *testing.T, state DeviceState) {
	if !state.IsActive() {
		t.Error("expected state to be active")
	}
}

func assertIsHalted(t *testing.T, state DeviceState) {
	if !state.IsHalted() {
		t.Error("expected state to be halted")
	}
}

func assertIsInactive(t *testing.T, state DeviceState) {
	if state.IsActive() {
		t.Error("expected state to be inactive")
	}
}

func assertIsNotHalted(t *testing.T, state DeviceState) {
	if state.IsHalted() {
		t.Error("expected state not to be halted")
	}
}
