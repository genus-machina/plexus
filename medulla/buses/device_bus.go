package buses

import (
	"fmt"
	"log"

	"github.com/genus-machina/plexus/medulla"
)

type deviceMap map[string]medulla.Device

type simpleDeviceBus struct {
	devices deviceMap
	log     *log.Logger
}

func NewDeviceBus(logger *log.Logger) *simpleDeviceBus {
	bus := new(simpleDeviceBus)
	bus.devices = make(deviceMap)
	bus.log = logger
	return bus
}

func (bus *simpleDeviceBus) Activate(name string) error {
	actuator, err := bus.getActuator(name)
	if err != nil {
		return err
	}

	err = actuator.Activate()
	if err == nil {
		bus.logActivation(name)
	}
	return err
}

func (bus *simpleDeviceBus) Deactivate(name string) error {
	actuator, err := bus.getActuator(name)
	if err != nil {
		return err
	}

	err = actuator.Deactivate()
	if err == nil {
		bus.logDeactivation(name)
	}
	return err
}

func (bus *simpleDeviceBus) getActuator(name string) (medulla.Actuator, error) {
	device := bus.devices[name]
	actuator, _ := device.(medulla.Actuator)

	if actuator == nil {
		return nil, fmt.Errorf("An actuator named '%s' has not been registered.", name)
	}

	return actuator, nil
}

func (bus *simpleDeviceBus) getTrigger(name string) (medulla.Trigger, error) {
	device := bus.devices[name]
	trigger, _ := device.(medulla.Trigger)

	if trigger == nil {
		return nil, fmt.Errorf("A trigger named '%s' has not been registered.", name)
	}

	return trigger, nil
}

func (bus *simpleDeviceBus) Halt() error {
	bus.log.Printf("Halting all devices...")
	for name, device := range bus.devices {
		if err := device.Halt(); err != nil {
			return fmt.Errorf("Device %s failed to halt: %s", name, err.Error())
		}
		bus.log.Printf("Device '%s' has halted.", name)
	}
	return nil
}

func (bus *simpleDeviceBus) logActivation(name string) {
	bus.log.Printf("Device '%s' is active.", name)
}

func (bus *simpleDeviceBus) logDeactivation(name string) {
	bus.log.Printf("Device '%s' is inactive.", name)
}

func (bus *simpleDeviceBus) RegisterDevice(device medulla.Device) error {
	name := device.Name()
	if existing := bus.devices[name]; existing != nil {
		return fmt.Errorf("A device named '%s' has already been registered.", name)
	}

	if trigger, ok := device.(medulla.Trigger); ok {
		go bus.watchTrigger(trigger)
	}

	bus.devices[name] = device
	bus.log.Printf("Registered new device '%s'.", name)
	return nil
}

func (bus *simpleDeviceBus) Subscribe(name string) (<-chan medulla.DeviceState, error) {
	trigger, err := bus.getTrigger(name)
	if err != nil {
		return nil, err
	}
	bus.log.Printf("Subscribing to device '%s'...", name)
	return trigger.Subscribe()
}

func (bus *simpleDeviceBus) watchTrigger(trigger medulla.Trigger) {
	name := trigger.Name()
	states, err := trigger.Subscribe()

	if err != nil {
		bus.log.Printf("Failed to watch trigger '%s': %s", name, err.Error())
		return
	}

	for state := range states {
		if state.IsActive() {
			bus.logActivation(name)
		} else {
			bus.logDeactivation(name)
		}
	}
}
