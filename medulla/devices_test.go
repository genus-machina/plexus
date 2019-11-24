package medulla

import (
	"testing"
)

func TestCheckDeviceState(t *testing.T) {
	device := NewTestDevice()
	assertCheckDeviceState(t, device)
	assertHalt(t, device)
	assertError(t, CheckDeviceState(device), "halted")
}
