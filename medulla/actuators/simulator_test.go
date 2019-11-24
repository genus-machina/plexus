package actuators

import (
	"testing"
)

func TestSimulatorInitialization(t *testing.T) {
	simulator := NewSimulator("test")
	assertInactive(t, simulator)
}

func TestSimulatorActivate(t *testing.T) {
	simulator := NewSimulator("test")
	assertActivate(t, simulator)
	assertActive(t, simulator)
}

func TestSimulatorDeactivate(t *testing.T) {
	simulator := NewSimulator("test")
	simulator.Activate()
	assertDeactivate(t, simulator)
	assertInactive(t, simulator)
}

func TestSimulatorName(t *testing.T) {
	simulator := NewSimulator("test")
	name := simulator.Name()
	assertName(t, simulator, name)
}

func TestSimulatorHalt(t *testing.T) {
	simulator := NewSimulator("test")
	err := simulator.Halt()

	if err != nil {
		t.Errorf("halting error: %s", err.Error())
	}

	if !simulator.State().IsHalted() {
		t.Error("failed to halt")
	}

	assertError(t, simulator.Activate(), "halted")
	assertError(t, simulator.Deactivate(), "halted")
}

func TestSimulatorHaltError(t *testing.T) {
	simulator := NewSimulator("test")
	simulator.Halt()
	assertError(t, simulator.Halt(), "halted")
}
