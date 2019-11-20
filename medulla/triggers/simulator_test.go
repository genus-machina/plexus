package triggers

import (
	"strings"
	"sync"
	"testing"

	"github.com/genus-machina/plexus/medulla"
)

func TestSimulatorName(t *testing.T) {
	trigger := NewSimulator("test")
	name := trigger.Name()

	if name != "test" {
		t.Errorf("expected name %s but got %s", "test", name)
	}
}

func TestSimulatorHalt(t *testing.T) {
	trigger := NewSimulator("test")
	err := trigger.Halt()

	if err != nil {
		t.Errorf("halting error: %s", err.Error())
	}

	if !trigger.State().IsHalted() {
		t.Error("trigger failed to halt")
	}
}

func TestSimulatorHaltError(t *testing.T) {
	trigger := NewSimulator("test")
	trigger.Halt()
	err := trigger.Halt()

	if message := err.Error(); !strings.Contains(message, "halted") {
		t.Errorf("unexpected error: %s", message)
	}
}

func TestSimulatorActivate(t *testing.T) {
	trigger := NewSimulator("test")
	err := trigger.Activate()

	if err != nil {
		t.Errorf("activation failure: %s", err.Error())
	}

	if !trigger.State().IsActive() {
		t.Error("trigger failed to activate")
	}
}

func TestSimulatorActivationError(t *testing.T) {
	trigger := NewSimulator("test")
	trigger.Halt()
	err := trigger.Activate()

	if message := err.Error(); !strings.Contains(message, "halted") {
		t.Errorf("unexpected error: %s", message)
	}
}

func TestSimulatorDeactivate(t *testing.T) {
	trigger := NewSimulator("test")
	trigger.Activate()
	err := trigger.Deactivate()

	if err != nil {
		t.Errorf("deactivation failure: %s", err.Error())
	}

	if trigger.State().IsActive() {
		t.Error("trigger failed to deactivate")
	}
}

func TestSimulatorDeactivationError(t *testing.T) {
	trigger := NewSimulator("test")
	trigger.Halt()
	err := trigger.Deactivate()

	if message := err.Error(); !strings.Contains(message, "halted") {
		t.Errorf("unexpected error: %s", message)
	}
}

func TestSimulatorSubscribe(t *testing.T) {
	trigger := NewSimulator("test")
	values := make([]medulla.DeviceState, 0)

	changes, err := trigger.Subscribe()
	if err != nil {
		t.Errorf("failed to subscribe: %s", err.Error())
	}

	go func() {
		trigger.Activate()
		trigger.Activate()
		trigger.Deactivate()
		trigger.Deactivate()
		trigger.Activate()
		trigger.Halt()
	}()

	for state := range changes {
		values = append(values, state)
	}

	if count := len(values); count != 3 {
		t.Errorf("expected %d values but found %d", 3, count)
	}

	if !values[0].IsActive() {
		t.Error("expected value 1 to be active")
	}

	if values[1].IsActive() {
		t.Error("expected value 2 to be inactive")
	}

	if !values[2].IsActive() {
		t.Error("expected value 3 to be active")
	}
}

func TestSimulatorMultipleSubscriptions(t *testing.T) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	trigger := NewSimulator("test")
	changes1, _ := trigger.Subscribe()
	changes2, _ := trigger.Subscribe()
	values1 := make([]medulla.DeviceState, 0)
	values2 := make([]medulla.DeviceState, 0)

	go func() {
		trigger.Activate()
		trigger.Deactivate()
		trigger.Halt()
	}()

	go func() {
		for state := range changes1 {
			values1 = append(values1, state)
		}
		waitGroup.Done()
	}()

	go func() {
		for state := range changes2 {
			values2 = append(values2, state)
		}
		waitGroup.Done()
	}()

	waitGroup.Wait()

	if count := len(values1); count != 2 {
		t.Errorf("expected %d values on channel 1 but found %d", 2, count)
	}

	if !values1[0].IsActive() {
		t.Error("expected value 1 on channel 1 to be active")
	}

	if values1[1].IsActive() {
		t.Error("expected value 2 on channel 1 to be inactive")
	}

	if count := len(values2); count != 2 {
		t.Errorf("expected %d values on channel 2 but found %d", 2, count)
	}

	if !values2[0].IsActive() {
		t.Error("expected value 1 on channel 2 to be active")
	}

	if values2[1].IsActive() {
		t.Error("expected value 2 on channel 2 to be inactive")
	}
}

func TestSimulatorSubscribeFailure(t *testing.T) {
	trigger := NewSimulator("test")
	trigger.Halt()
	channel, err := trigger.Subscribe()

	if message := err.Error(); !strings.Contains(message, "halted") {
		t.Errorf("unexpected error: %s", message)
	}

	if channel != nil {
		t.Error("expected channel to be nil")
	}
}
