package medulla

import (
	"testing"
)

func TestDeviceStateIsActive(t *testing.T) {
	state := NewDeviceState(false, false)
	if state.IsActive() {
		t.Error("state is active")
	}

	state = NewDeviceState(true, false)
	if !state.IsActive() {
		t.Error("state is not active")
	}
}

func TestDeviceStateIsHalted(t *testing.T) {
	state := NewDeviceState(false, false)
	if state.IsHalted() {
		t.Error("state is halted")
	}

	state = NewDeviceState(false, true)
	if !state.IsHalted() {
		t.Error("state is not halted")
	}
}
