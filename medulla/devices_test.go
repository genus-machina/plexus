package medulla

import (
	"strings"
	"testing"
)

func TestCheckDeviceState(t *testing.T) {
	device := NewTestDevice()
	err := CheckDeviceState(device)
	if err != nil {
		t.Errorf("device failed check: %s", err.Error())
	}

	device.Halt()
	err = CheckDeviceState(device)
	if message := err.Error(); !strings.Contains(message, "halted") {
		t.Errorf("unexpected error: %s", message)
	}
}
