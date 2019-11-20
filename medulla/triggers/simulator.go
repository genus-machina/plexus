package triggers

import (
	"github.com/genus-machina/plexus/medulla"
)

type Simulator struct {
	active   bool
	channels []chan medulla.DeviceState
	halted   bool
	name     string
}

func NewSimulator(name string) *Simulator {
	simulator := new(Simulator)
	simulator.channels = make([]chan medulla.DeviceState, 0)
	simulator.name = name
	return simulator
}

func (device *Simulator) Activate() error {
	if err := medulla.CheckDeviceState(device); err != nil {
		return err
	}

	previous := device.active
	device.active = true

	if !previous {
		device.emit()
	}

	return nil
}

func (device *Simulator) Deactivate() error {
	if err := medulla.CheckDeviceState(device); err != nil {
		return err
	}

	previous := device.active
	device.active = false

	if previous {
		device.emit()
	}

	return nil
}

func (device *Simulator) Halt() error {
	if err := medulla.CheckDeviceState(device); err != nil {
		return err
	}

	if len(device.channels) > 0 {
		for _, channel := range device.channels {
			close(channel)
		}
	}

	device.halted = true
	return nil
}

func (device *Simulator) Name() string {
	return device.name
}

func (device *Simulator) emit() {
	if len(device.channels) == 0 {
		return
	}

	for _, channel := range device.channels {
		channel <- device.State()
	}
}

func (device *Simulator) State() medulla.DeviceState {
	return medulla.NewDeviceState(device.active, device.halted)
}

func (device *Simulator) Subscribe() (<-chan medulla.DeviceState, error) {
	if err := medulla.CheckDeviceState(device); err != nil {
		return nil, err
	}

	channel := make(chan medulla.DeviceState)
	device.channels = append(device.channels, channel)
	return channel, nil
}
