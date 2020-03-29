package actuators

import (
	"time"

	"github.com/genus-machina/plexus/medulla"
)

type Simulator struct {
	active bool
	halted bool
	name   string
}

func NewSimulator(name string) *Simulator {
	simulator := new(Simulator)
	simulator.name = name
	return simulator
}

func (device *Simulator) Activate() error {
	if err := medulla.CheckDeviceState(device); err != nil {
		return err
	}
	device.active = true
	return nil
}

func (device *Simulator) Deactivate() error {
	if err := medulla.CheckDeviceState(device); err != nil {
		return err
	}
	device.active = false
	return nil
}

func (device *Simulator) Halt() error {
	if err := medulla.CheckDeviceState(device); err != nil {
		return err
	}
	device.halted = true
	return nil
}

func (device *Simulator) Name() string {
	return device.name
}

func (device *Simulator) State() medulla.DeviceState {
	return medulla.NewDeviceState(device.active, device.halted, time.Now())
}
