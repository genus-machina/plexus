package bus

import (
	"log"
	"strings"
	"testing"

	"github.com/genus-machina/plexus/hippocampus"
	"github.com/genus-machina/plexus/medulla"
	"github.com/genus-machina/plexus/medulla/actuators"
	"github.com/genus-machina/plexus/medulla/triggers"
)

var (
	logger *log.Logger = hippocampus.NewLogger("device bus test")
)

func TestDeviceBusActivateExisting(t *testing.T) {
	bus := New(logger)
	indicator := actuators.NewSimulator("test indicator")
	bus.RegisterDevice(indicator)
	err := bus.Activate("test indicator")

	if err != nil {
		t.Errorf("activation failed: %s", err.Error())
	}

	if !indicator.State().IsActive() {
		t.Error("indicator failed to activate")
	}
}

func TestDeviceBusActivateMissing(t *testing.T) {
	bus := New(logger)
	err := bus.Activate("test indicator")

	if message := err.Error(); !strings.Contains(message, "registered") {
		t.Errorf("unexpected error: %s", message)
	}
}

func TestDeviceBusDeactivateExisting(t *testing.T) {
	bus := New(logger)
	indicator := actuators.NewSimulator("test indicator")
	bus.RegisterDevice(indicator)
	bus.Activate("test indicator")
	err := bus.Deactivate("test indicator")

	if err != nil {
		t.Errorf("deactivation failed: %s", err.Error())
	}

	if indicator.State().IsActive() {
		t.Error("indicator failed to deactivate")
	}
}

func TestDeviceBusDeactivateMissing(t *testing.T) {
	bus := New(logger)
	err := bus.Deactivate("test indicator")

	if message := err.Error(); !strings.Contains(message, "registered") {
		t.Errorf("unexpected error: %s", message)
	}
}

func TestDeviceBusActivationError(t *testing.T) {
	bus := New(logger)
	indicator := actuators.NewSimulator("test indicator")
	bus.RegisterDevice(indicator)
	indicator.Halt()
	err := bus.Activate("test indicator")

	if message := err.Error(); !strings.Contains(message, "halted") {
		t.Errorf("unexpected error: %s", message)
	}
}

func TestDeviceBusDeactivationError(t *testing.T) {
	bus := New(logger)
	indicator := actuators.NewSimulator("test indicator")
	bus.RegisterDevice(indicator)
	indicator.Halt()
	err := bus.Deactivate("test indicator")

	if message := err.Error(); !strings.Contains(message, "halted") {
		t.Errorf("unexpected error: %s", message)
	}
}

func TestDeviceBusHalt(t *testing.T) {
	bus := New(logger)
	indicator1 := actuators.NewSimulator("one")
	indicator2 := actuators.NewSimulator("two")
	bus.RegisterDevice(indicator1)
	bus.RegisterDevice(indicator2)
	err := bus.Halt()

	if err != nil {
		t.Errorf("halting failed: %s", err.Error())
	}

	if !indicator1.State().IsHalted() {
		t.Error("indicator1 failed to halt")
	}

	if !indicator2.State().IsHalted() {
		t.Error("indicator2 failed to halt")
	}
}

func TestDeviceBusHaltError(t *testing.T) {
	bus := New(logger)
	indicator1 := actuators.NewSimulator("one")
	indicator2 := actuators.NewSimulator("two")
	bus.RegisterDevice(indicator1)
	bus.RegisterDevice(indicator2)
	indicator1.Halt()
	err := bus.Halt()

	if message := err.Error(); !strings.Contains(message, "one failed to halt") {
		t.Errorf("unexpected error: %s", message)
	}
}

func TestDeviceBusSubscribeMissing(t *testing.T) {
	bus := New(logger)
	channel, err := bus.Subscribe("test")

	if channel != nil {
		t.Error("expected channel to be nil")
	}

	if message := err.Error(); !strings.Contains(message, "registered") {
		t.Errorf("unexpected error: %s", message)
	}
}

func TestDeviceBusSubscribeExisting(t *testing.T) {
	bus := New(logger)
	trigger := triggers.NewSimulator("trigger")
	bus.RegisterDevice(trigger)
	changes, err := bus.Subscribe("trigger")

	if err != nil {
		t.Errorf("subscription error: %s", err.Error())
	}

	go func() {
		trigger.Activate()
		trigger.Deactivate()
		trigger.Halt()
	}()

	values := make([]medulla.DeviceState, 0)
	for value := range changes {
		values = append(values, value)
	}

	if count := len(values); count != 2 {
		t.Errorf("expected 2 values but got %d", count)
	}

	if !values[0].IsActive() {
		t.Error("expected value 1 to be active")
	}

	if values[1].IsActive() {
		t.Error("expected value 2 to be inactive")
	}
}

func TestDeviceBusRegisterDuplicate(t *testing.T) {
	bus := New(logger)
	actuator := actuators.NewSimulator("test")
	trigger := actuators.NewSimulator("test")
	bus.RegisterDevice(actuator)

	err := bus.RegisterDevice(actuator)
	if message := err.Error(); !strings.Contains(message, "registered") {
		t.Errorf("unexpected error: %s", message)
	}

	err = bus.RegisterDevice(trigger)
	if message := err.Error(); !strings.Contains(message, "registered") {
		t.Errorf("unexpected error: %s", message)
	}
}
