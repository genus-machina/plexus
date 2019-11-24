package actuators

import (
	"strings"
	"testing"
)

func assertActivate(t *testing.T, simulator *Simulator) {
	if err := simulator.Activate(); err != nil {
		t.Errorf("activation error: %s", err.Error())
	}
}

func assertActive(t *testing.T, simulator *Simulator) {
	if !simulator.State().IsActive() {
		t.Error("expected device to be active")
	}
}

func assertDeactivate(t *testing.T, simulator *Simulator) {
	if err := simulator.Deactivate(); err != nil {
		t.Errorf("deactivation error: %s", err.Error())
	}
}

func assertError(t *testing.T, err error, expected string) {
	if message := err.Error(); !strings.Contains(message, expected) {
		t.Errorf("unexpected error: %s", message)
	}
}

func assertInactive(t *testing.T, simulator *Simulator) {
	if simulator.State().IsActive() {
		t.Error("expected device to be inactive")
	}
}

func assertName(t *testing.T, simulator *Simulator, expected string) {
	if name := simulator.Name(); name != expected {
		t.Errorf("expected '%s' but got '%s'", expected, name)
	}
}
