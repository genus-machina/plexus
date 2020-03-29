package triggers

import (
	"sync"
	"testing"
	"time"

	"github.com/genus-machina/plexus/medulla"
)

func TestSimulatorName(t *testing.T) {
	trigger := NewSimulator("test")
	assertName(t, trigger, "test")
}

func TestSimulatorHalt(t *testing.T) {
	trigger := NewSimulator("test")
	assertHalt(t, trigger)
	assertIsHalted(t, trigger)
}

func TestSimulatorHaltError(t *testing.T) {
	trigger := NewSimulator("test")
	assertHalt(t, trigger)
	assertError(t, trigger.Halt(), "halted")
}

func TestSimulatorActivate(t *testing.T) {
	trigger := NewSimulator("test")
	assertActivate(t, trigger)
	assertIsActive(t, trigger)
}

func TestSimulatorActivationError(t *testing.T) {
	trigger := NewSimulator("test")
	assertHalt(t, trigger)
	assertError(t, trigger.Activate(), "halted")
}

func TestSimulatorDeactivate(t *testing.T) {
	trigger := NewSimulator("test")
	assertActivate(t, trigger)
	assertDeactivate(t, trigger)
	assertIsInactive(t, trigger)
}

func TestSimulatorDeactivationError(t *testing.T) {
	trigger := NewSimulator("test")
	assertHalt(t, trigger)
	assertError(t, trigger.Deactivate(), "halted")
}

func TestSimulatorSubscribe(t *testing.T) {
	trigger := NewSimulator("test")
	values := make([]medulla.DeviceState, 0)
	changes := assertSubscribe(t, trigger)

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

	expected := []medulla.DeviceState{
		medulla.NewDeviceState(true, false, time.Unix(0, 0)),
		medulla.NewDeviceState(false, false, time.Unix(0, 0)),
		medulla.NewDeviceState(true, false, time.Unix(0, 0)),
	}

	assertStates(t, expected, values)
}

func TestSimulatorMultipleSubscriptions(t *testing.T) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	trigger := NewSimulator("test")
	changes1 := assertSubscribe(t, trigger)
	changes2 := assertSubscribe(t, trigger)
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

	expected := []medulla.DeviceState{
		medulla.NewDeviceState(true, false, time.Unix(0, 0)),
		medulla.NewDeviceState(false, false, time.Unix(0, 0)),
	}

	assertStates(t, expected, values1)
	assertStates(t, expected, values2)
}

func TestSimulatorSubscribeFailure(t *testing.T) {
	trigger := NewSimulator("test")
	assertHalt(t, trigger)
	channel, err := trigger.Subscribe()

	assertError(t, err, "halted")

	if channel != nil {
		t.Error("expected channel to be nil")
	}
}
