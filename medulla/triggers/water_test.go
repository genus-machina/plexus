package triggers

import (
	"testing"
	"time"

	"github.com/genus-machina/plexus/medulla"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpiotest"
)

func TestWaterHalt(t *testing.T) {
	name := "test"
	pin := new(gpiotest.Pin)
	pin.EdgesChan = make(chan gpio.Level)
	device := assertWater(t, name, pin)
	assertIsNotHalted(t, device)
	assertHalt(t, device)
	assertIsHalted(t, device)
}

func TestWaterName(t *testing.T) {
	name := "test"
	pin := new(gpiotest.Pin)
	pin.EdgesChan = make(chan gpio.Level)
	device := assertWater(t, name, pin)
	assertName(t, device, name)
}

func TestWaterState(t *testing.T) {
	name := "test"
	pin := new(gpiotest.Pin)
	pin.EdgesChan = make(chan gpio.Level)
	device := assertWater(t, name, pin)
	// override denoise to speed up tests...
	device.denoisePeriod = 100 * time.Millisecond
	states := assertSubscribe(t, device)
	results := make([]medulla.DeviceState, 0)

	assertPullDown(t, pin)
	assertIsInactive(t, device)

	go func() {
		pin.EdgesChan <- gpio.High
	}()

	results = append(results, <-states)
	assertIsActive(t, device)

	go func() {
		pin.EdgesChan <- gpio.Low
	}()

	results = append(results, <-states)
	assertIsInactive(t, device)

	expected := []medulla.DeviceState{
		medulla.NewDeviceState(true, false, time.Unix(0, 0)),
		medulla.NewDeviceState(false, false, time.Unix(0, 0)),
	}
	assertStates(t, expected, results)
}

func TestWaterSubscribeError(t *testing.T) {
	name := "test"
	pin := new(gpiotest.Pin)
	pin.EdgesChan = make(chan gpio.Level)
	device := assertWater(t, name, pin)
	assertHalt(t, device)

	_, err := device.Subscribe()
	assertError(t, err, "halted")
}
