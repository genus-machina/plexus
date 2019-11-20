package medulla

type simpleDeviceState struct {
	active bool
	halted bool
}

func NewDeviceState(active bool, halted bool) *simpleDeviceState {
	state := new(simpleDeviceState)
	state.active = active
	state.halted = halted
	return state
}

func (state *simpleDeviceState) IsActive() bool {
	return state.active
}

func (state *simpleDeviceState) IsHalted() bool {
	return state.halted
}
