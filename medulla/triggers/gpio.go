package triggers

import (
	"sync"
	"time"

	"github.com/genus-machina/plexus/medulla"
	"periph.io/x/periph/conn/gpio"
)

type gpioTrigger struct {
	active         bool
	debouncePeriod time.Duration
	denoisePeriod  time.Duration
	halted         bool
	inverted       bool
	lastTime       time.Time
	level          gpio.Level
	mutex          sync.Mutex
	name           string
	pin            gpio.PinIn
	subscriptions  []chan medulla.DeviceState
	timer          *time.Timer
}

func (device *gpioTrigger) Halt() error {
	var err error
	if err = device.pin.Halt(); err == nil {
		device.halted = true
	}
	return err
}

func (device *gpioTrigger) Name() string {
	return device.name
}

func (device *gpioTrigger) State() medulla.DeviceState {
	return medulla.NewDeviceState(device.active, device.halted)
}

func (device *gpioTrigger) Subscribe() (<-chan medulla.DeviceState, error) {
	if err := medulla.CheckDeviceState(device); err != nil {
		return nil, err
	}

	device.mutex.Lock()
	defer device.mutex.Unlock()

	subscription := make(chan medulla.DeviceState, 0)
	device.subscriptions = append(device.subscriptions, subscription)
	return subscription, nil
}

func (device *gpioTrigger) broadcast() {
	device.mutex.Lock()
	defer device.mutex.Unlock()

	for _, subscription := range device.subscriptions {
		subscription <- medulla.NewDeviceState(device.active, device.halted)
	}
}

func (device *gpioTrigger) debounce() {
	value := bool(device.level)

	if device.inverted {
		value = !value
	}

	if now := time.Now(); value != device.active || now.After(device.lastTime.Add(device.debouncePeriod)) {
		device.active = value
		go device.broadcast()
	}
}

func (device *gpioTrigger) watchEdges() {
	for {
		if device.pin.WaitForEdge(-1) {
			device.level = device.pin.Read()

			if device.timer != nil {
				device.timer.Stop()
			}

			if device.denoisePeriod > 0 {
				device.timer = time.AfterFunc(device.denoisePeriod, device.debounce)
			} else {
				device.debounce()
			}
		}
	}
}
