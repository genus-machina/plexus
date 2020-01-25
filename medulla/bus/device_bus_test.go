package bus

import (
	"log"
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
	assertRegister(t, bus, indicator)
	assertActivate(t, bus, "test indicator")
	assertActive(t, indicator)
}

func TestDeviceBusActivateMissing(t *testing.T) {
	bus := New(logger)
	assertError(t, bus.Activate("test indicator"), "registered")
}

func TestDeviceBusDeactivateExisting(t *testing.T) {
	bus := New(logger)
	indicator := actuators.NewSimulator("test indicator")
	assertRegister(t, bus, indicator)
	bus.Activate("test indicator")
	assertDeactivate(t, bus, "test indicator")
	assertInactive(t, indicator)
}

func TestDeviceBusDeactivateMissing(t *testing.T) {
	bus := New(logger)
	assertError(t, bus.Deactivate("test indicator"), "registered")
}

func TestDeviceBusActivationError(t *testing.T) {
	bus := New(logger)
	indicator := actuators.NewSimulator("test indicator")
	assertRegister(t, bus, indicator)
	assertHaltDevice(t, indicator)
	assertError(t, bus.Activate("test indicator"), "halted")
}

func TestDeviceBusDeactivationError(t *testing.T) {
	bus := New(logger)
	indicator := actuators.NewSimulator("test indicator")
	assertRegister(t, bus, indicator)
	bus.Activate("test indicator")
	assertHaltDevice(t, indicator)
	assertError(t, bus.Deactivate("test indicator"), "halted")
}

func TestDeviceBusHalt(t *testing.T) {
	bus := New(logger)
	indicator1 := actuators.NewSimulator("one")
	indicator2 := actuators.NewSimulator("two")
	assertRegister(t, bus, indicator1)
	assertRegister(t, bus, indicator2)
	assertHalt(t, bus)
	assertHalted(t, indicator1)
	assertHalted(t, indicator2)
}

func TestDeviceBusHaltError(t *testing.T) {
	bus := New(logger)
	indicator1 := actuators.NewSimulator("one")
	indicator2 := actuators.NewSimulator("two")
	assertRegister(t, bus, indicator1)
	assertRegister(t, bus, indicator2)
	assertHaltDevice(t, indicator1)
	assertError(t, bus.Halt(), "one failed to halt")
}

func TestDeviceBusSubscribeMissing(t *testing.T) {
	bus := New(logger)
	channel, err := bus.Subscribe("test")

	if channel != nil {
		t.Errorf("expcted channel to be nil")
	}

	assertError(t, err, "registered")
}

func TestDeviceBusSubscribeExisting(t *testing.T) {
	bus := New(logger)
	trigger := triggers.NewSimulator("trigger")
	assertRegister(t, bus, trigger)
	changes := assertSubscribe(t, bus, "trigger")

	go func() {
		trigger.Activate()
		trigger.Deactivate()
		trigger.Halt()
	}()

	values := make([]medulla.DeviceState, 0)
	for value := range changes {
		values = append(values, value)
	}

	expected := []medulla.DeviceState{
		medulla.NewDeviceState(true, false),
		medulla.NewDeviceState(false, false),
	}

	assertStates(t, expected, values)
}

func TestDeviceBusRegisterDuplicate(t *testing.T) {
	bus := New(logger)
	actuator := actuators.NewSimulator("test")
	trigger := actuators.NewSimulator("test")
	assertRegister(t, bus, actuator)
	assertError(t, bus.RegisterDevice(actuator), "registered")
	assertError(t, bus.RegisterDevice(trigger), "registered")
}
