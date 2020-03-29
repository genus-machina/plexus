package medulla

import (
	"time"
)

type simpleDeviceState struct {
	active bool
	halted bool
	time   time.Time
}

func NewDeviceState(active bool, halted bool, time time.Time) *simpleDeviceState {
	state := new(simpleDeviceState)
	state.active = active
	state.halted = halted
	state.time = time
	return state
}

func (state *simpleDeviceState) IsActive() bool {
	return state.active
}

func (state *simpleDeviceState) IsHalted() bool {
	return state.halted
}

func (state *simpleDeviceState) Time() time.Time {
	return state.time
}
