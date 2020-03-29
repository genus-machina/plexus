package triggers

import (
	"testing"
	"time"

	"github.com/genus-machina/plexus/medulla"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpiotest"
)

func TestButtonHalt(t *testing.T) {
	name := "test"
	pin := new(gpiotest.Pin)
	pin.EdgesChan = make(chan gpio.Level)
	button := assertButton(t, name, pin)

	assertIsNotHalted(t, button)
	assertHalt(t, button)
	assertIsHalted(t, button)
}

func TestButtonName(t *testing.T) {
	name := "test"
	pin := new(gpiotest.Pin)
	pin.EdgesChan = make(chan gpio.Level)
	button := assertButton(t, name, pin)
	assertName(t, button, name)
}

func TestButtonState(t *testing.T) {
	name := "test"
	pin := new(gpiotest.Pin)
	pin.EdgesChan = make(chan gpio.Level)
	button := assertButton(t, name, pin)
	states := assertSubscribe(t, button)
	results := make([]medulla.DeviceState, 0)

	assertPullUp(t, pin)
	assertIsInactive(t, button)

	go func() {
		pin.EdgesChan <- gpio.Low
	}()

	results = append(results, <-states)
	assertIsActive(t, button)
	<-time.After(100 * time.Millisecond)

	go func() {
		pin.EdgesChan <- gpio.High
	}()

	results = append(results, <-states)
	assertIsInactive(t, button)

	expected := []medulla.DeviceState{
		medulla.NewDeviceState(true, false, time.Unix(0, 0)),
		medulla.NewDeviceState(false, false, time.Unix(0, 0)),
	}
	assertStates(t, expected, results)
}

func TestButtonSubscribeError(t *testing.T) {
	name := "test"
	pin := new(gpiotest.Pin)
	pin.EdgesChan = make(chan gpio.Level)
	button := assertButton(t, name, pin)
	assertHalt(t, button)

	_, err := button.Subscribe()
	assertError(t, err, "halted")
}
