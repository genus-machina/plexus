package hypothalamus

import (
	"time"

	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/devices/bmxx80"
)

type BME280 struct {
	device *bmxx80.Dev
}

func NewBME280(device *bmxx80.Dev) *BME280 {
	bme := new(BME280)
	bme.device = device
	return bme
}

func (device *BME280) convert(values <-chan physic.Env, measurements chan<- *Environmental) {
	defer close(measurements)

	for value := range values {
		measurement := Environmental(value)
		measurements <- &measurement
	}
}

func (device *BME280) Halt() error {
	return device.device.Halt()
}

func (device *BME280) SenseContinuous(interval time.Duration) (<-chan *Environmental, error) {
	if values, err := device.device.SenseContinuous(interval); err == nil {
		measurements := make(chan *Environmental, 0)
		go device.convert(values, measurements)
		return measurements, nil
	} else {
		return nil, err
	}
}
