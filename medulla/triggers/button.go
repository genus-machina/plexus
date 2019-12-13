package triggers

import (
	"time"

	"github.com/genus-machina/plexus/medulla"
	"periph.io/x/periph/conn/gpio"
)

type Button struct {
	gpioTrigger
}

func NewButton(name string, pin gpio.PinIO) (*Button, error) {
	device := new(Button)
	device.debouncePeriod = 200 * time.Millisecond
	device.denoisePeriod = 10 * time.Millisecond
	device.inverted = true
	device.name = name
	device.pin = pin
	device.subscriptions = make([]chan medulla.DeviceState, 0)

	if err := device.pin.In(gpio.PullUp, gpio.BothEdges); err != nil {
		return nil, err
	}

	go device.watchEdges()
	return device, nil
}
