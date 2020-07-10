package triggers

import (
	"time"

	"github.com/genus-machina/plexus/medulla"
	"periph.io/x/periph/conn/gpio"
)

type Button struct {
	gpioTrigger
}

func NewButton(name string, pin gpio.PinIO, inverted bool) (*Button, error) {
	device := new(Button)
	device.debouncePeriod = 200 * time.Millisecond
	device.denoisePeriod = 10 * time.Millisecond
	device.inverted = inverted
	device.name = name
	device.pin = pin
	device.subscriptions = make([]chan medulla.DeviceState, 0)

	pull := gpio.PullDown
	if device.inverted {
		pull = gpio.PullUp
	}

	if err := device.pin.In(pull, gpio.BothEdges); err != nil {
		return nil, err
	}

	go device.watchEdges()
	return device, nil
}
