package triggers

import (
	"time"

	"github.com/genus-machina/plexus/medulla"
	"periph.io/x/periph/conn/gpio"
)

type Water struct {
	gpioTrigger
}

func NewWater(name string, pin gpio.PinIn) (*Water, error) {
	device := new(Water)
	device.debouncePeriod = 5 * time.Minute
	device.denoisePeriod = 1 * time.Minute
	device.name = name
	device.pin = pin
	device.subscriptions = make([]chan medulla.DeviceState, 0)

	if err := pin.In(gpio.PullDown, gpio.BothEdges); err != nil {
		return nil, err
	}

	go device.watchEdges()
	return device, nil
}
