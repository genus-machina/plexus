package hypothalamus

import (
	"log"
	"time"

	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/devices/bmxx80"
)

type BME280 struct {
	address uint16
	bus     i2c.Bus
	device  *bmxx80.Dev
	done    chan interface{}
}

func NewBME280(bus i2c.Bus, address uint16) (*BME280, error) {
	bme := new(BME280)
	bme.address = address
	bme.bus = bus
	bme.done = make(chan interface{}, 0)
	return bme, bme.buildDevice()
}

func (device *BME280) buildDevice() error {
	var err error

	options := &bmxx80.Opts{
		Temperature: bmxx80.O16x,
		Pressure:    bmxx80.O16x,
		Humidity:    bmxx80.O16x,
	}

	device.device, err = bmxx80.NewI2C(device.bus, device.address, options)
	return err
}

func (device *BME280) Halt() error {
	close(device.done)
	return device.device.Halt()
}

func (device *BME280) readValues(measurements chan Environmental, interval time.Duration) {
	defer close(measurements)

	for {
		var value physic.Env
		if err := device.device.Sense(&value); err == nil {
			measurement := new(physicEnv)
			measurement.env = value
			measurement.time = time.Now()
			measurements <- measurement
		} else {
			log.Printf("BME280 failed. Attempting to rebuild. %s\n", err.Error())
			device.device.Halt()
			device.reset()
			if err := device.buildDevice(); err != nil {
				log.Fatalf("Failed to rebuild BME280. %s\n", err.Error())
			}
		}

		select {
		case <-time.After(interval):
		case <-device.done:
			return
		}
	}
}

func (device *BME280) reset() error {
	address := uint16(0xE0)
	command := []byte{0xB6}
	return device.bus.Tx(address, command, nil)
}

func (device *BME280) SenseContinuous(interval time.Duration) <-chan Environmental {
	measurements := make(chan Environmental, 0)
	go device.readValues(measurements, interval)
	return measurements
}
