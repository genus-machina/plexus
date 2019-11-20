package zygote

import (
	"testing"

	"github.com/genus-machina/plexus/hippocampus"
)

func TestLoadSimulators(t *testing.T) {
	logger := hippocampus.NewLogger("loader test")
	system, err := LoadJSON("./simulators.json", logger)
	if err != nil {
		t.Errorf("failed to load system: %s", err.Error())
	}

	if err := system.DeviceBus.Activate("actuator"); err != nil {
		t.Errorf("failed to activate actuator: %s", err.Error())
	}

	if err := system.DeviceBus.Deactivate("actuator"); err != nil {
		t.Errorf("failed to deactivate actuator: %s", err.Error())
	}

	if err := system.DeviceBus.Halt(); err != nil {
		t.Errorf("failed to halt device bus: %s", err.Error())
	}
}
