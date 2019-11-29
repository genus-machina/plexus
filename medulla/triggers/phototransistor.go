package triggers

import (
	"time"

	"periph.io/x/periph/conn/gpio"
	"github.com/genus-machina/plexus/medulla"
)

type Phototransistor struct {
	gpioTrigger
}

func NewPhototransistor(name string, pin gpio.PinIn) (*Phototransistor, error) {
	device := new(Phototransistor)
	device.debouncePeriod = 5*time.Minute
	device.denoisePeriod = 1*time.Minute
	device.name = name
	device.pin = pin
	device.subscriptions = make([]chan medulla.DeviceState, 0)

	if err := pin.In(gpio.PullDown, gpio.BothEdges); err != nil {
		return nil, err
	}

	go device.watchEdges()
	return device, nil
}
