package medulla

import (
	"testing"
	"time"
)

func TestDeviceStateIsActive(t *testing.T) {
	state := NewDeviceState(false, false, time.Now())
	assertIsInactive(t, state)

	state = NewDeviceState(true, false, time.Now())
	assertIsActive(t, state)
}

func TestDeviceStateIsHalted(t *testing.T) {
	state := NewDeviceState(false, false, time.Now())
	assertIsNotHalted(t, state)

	state = NewDeviceState(false, true, time.Now())
	assertIsHalted(t, state)
}

func TestDeviceStateTime(t *testing.T) {
	now := time.Now()
	state := NewDeviceState(false, false, now)
	if stateTime := state.Time(); stateTime != now {
		t.Errorf("Expected %s but got %s", now.Format(time.RFC3339), stateTime.Format(time.RFC3339))
	}
}
