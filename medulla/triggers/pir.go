package triggers

import (
	"periph.io/x/periph/conn/gpio"

	"github.com/genus-machina/plexus/medulla"
)

type PIR struct {
	gpioTrigger
}

func NewPIR(name string, pin gpio.PinIn) (*PIR, error) {
	device := new(PIR)
	device.name = name
	device.pin = pin
	device.subscriptions = make([]chan medulla.DeviceState, 0)

	if err := pin.In(gpio.PullDown, gpio.BothEdges); err != nil {
		return nil, err
	}

	go device.watchEdges()
	return device, nil
}
