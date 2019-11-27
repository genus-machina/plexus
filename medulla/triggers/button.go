package triggers

import (
	"sync"
	"time"

	"github.com/genus-machina/plexus/medulla"
	"periph.io/x/periph/conn/gpio"
)

const (
	debounce = 200*time.Millisecond
	denoise = 10*time.Millisecond
)

type Button struct {
	active bool
	halted bool
	lastTime time.Time
	level gpio.Level
	mutex sync.Mutex
	name string
	pin gpio.PinIn
	subscriptions []chan medulla.DeviceState
	timer *time.Timer
}

func NewButton(name string, pin gpio.PinIO) (*Button, error) {
	device := new(Button)
	device.name = name
	device.pin = pin
	device.subscriptions = make([]chan medulla.DeviceState, 0)

	if err := device.pin.In(gpio.PullUp, gpio.BothEdges); err != nil {
		return nil, err
	}

	go device.watchEdges()
	return device, nil
}

func (device *Button) Halt() error {
	var err error
	if err = device.pin.Halt(); err == nil {
		device.halted = true
	}
	return err
}

func (device *Button) Name() string {
	return device.name
}

func (device *Button) State() medulla.DeviceState {
	return medulla.NewDeviceState(device.active, device.halted)
}

func (device *Button) Subscribe() (<-chan medulla.DeviceState, error) {
	if err := medulla.CheckDeviceState(device); err != nil {
		return nil, err
	}

	device.mutex.Lock()
	defer device.mutex.Unlock()

	subscription := make(chan medulla.DeviceState, 0)
	device.subscriptions = append(device.subscriptions, subscription)
	return subscription, nil
}

func (device *Button) broadcast() {
	device.mutex.Lock()
	defer device.mutex.Unlock()

	for _, subscription := range device.subscriptions {
		subscription <- medulla.NewDeviceState(device.active, device.halted)
	}
}

func (device *Button) debounce() {
	value := !bool(device.level)

	if now := time.Now(); value != device.active || now.After(device.lastTime.Add(debounce)) {
		device.active = value
		go device.broadcast()
	}
}

func (device *Button) watchEdges() {
	for {
		if device.pin.WaitForEdge(-1) {
			device.level = device.pin.Read()

			if device.timer != nil {
				device.timer.Stop()
			}

			device.timer = time.AfterFunc(denoise, device.debounce)
		}
	}
}
