package actuators

import (
	"strings"
	"testing"
)

func TestSimulatorInitialization(t *testing.T) {
	simulator := NewSimulator("test")

	if simulator.State().IsActive() {
		t.Error("indicator is active")
	}
}

func TestSimulatorActivate(t *testing.T) {
	simulator := NewSimulator("test")
	err := simulator.Activate()

	if err != nil {
		t.Errorf("activation error: %s", err.Error())
	}

	if !simulator.State().IsActive() {
		t.Error("indicator is not active")
	}
}

func TestSimulatorDeactivate(t *testing.T) {
	simulator := NewSimulator("test")
	simulator.Activate()
	err := simulator.Deactivate()

	if err != nil {
		t.Errorf("deactivation error: %s", err.Error())
	}

	if simulator.State().IsActive() {
		t.Error("indicator is active")
	}
}

func TestSimulatorName(t *testing.T) {
	simulator := NewSimulator("test")
	name := simulator.Name()

	if name != "test" {
		t.Errorf("expected %s but got %s", "test", name)
	}
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

	err = simulator.Activate()

	if message := err.Error(); !strings.Contains(message, "halted") {
		t.Errorf("unexpected error: %s", message)
	}

	err = simulator.Deactivate()

	if message := err.Error(); !strings.Contains(message, "halted") {
		t.Errorf("unexpected error: %s", message)
	}
}

func TestSimulatorHaltError(t *testing.T) {
	simulator := NewSimulator("test")
	simulator.Halt()
	err := simulator.Halt()

	if message := err.Error(); !strings.Contains(message, "halted") {
		t.Errorf("unexpected error: %s", message)
	}
}
