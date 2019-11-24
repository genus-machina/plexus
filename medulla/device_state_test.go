package medulla

import (
	"testing"
)

func TestDeviceStateIsActive(t *testing.T) {
	state := NewDeviceState(false, false)
	assertIsInactive(t, state)

	state = NewDeviceState(true, false)
	assertIsActive(t, state)
}

func TestDeviceStateIsHalted(t *testing.T) {
	state := NewDeviceState(false, false)
	assertIsNotHalted(t, state)

	state = NewDeviceState(false, true)
	assertIsHalted(t, state)
}
