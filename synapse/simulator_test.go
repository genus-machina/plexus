package synapse

import (
	"log"
	"os"
	"sync"
	"testing"

	"github.com/genus-machina/plexus/hypothalamus"
	"github.com/genus-machina/plexus/medulla/actuators"
)

var (
	logger *log.Logger = log.New(os.Stderr, "simulator test", 0)
)

func TestSimulatorApplyUnknownMessage(t *testing.T) {
	simulator := NewSimulator(logger)
	actuator := actuators.NewSimulator("test")
	message := []byte("unknown")
	assertApply(t, simulator, message, actuator)
	assertInactive(t, actuator)
}

func TestSimulatorApplyActivateMessage(t *testing.T) {
	simulator := NewSimulator(logger)
	actuator := actuators.NewSimulator("test")
	message := SIMULATOR_ACTIVATED
	assertApply(t, simulator, message, actuator)
	assertActive(t, actuator)
}

func TestSimulatorApplyDeactivateMessage(t *testing.T) {
	simulator := NewSimulator(logger)
	actuator := actuators.NewSimulator("test")
	message := SIMULATOR_DEACTIVATED
	assertActivate(t, actuator)
	assertApply(t, simulator, message, actuator)
	assertInactive(t, actuator)
}

func TestSimulatorPublishMissing(t *testing.T) {
	simulator := NewSimulator(logger)
	message := []byte("message")
	topic := "topic"
	assertPublish(t, simulator, message, topic)
}

func TestSimulatorSubscribeMissing(t *testing.T) {
	simulator := NewSimulator(logger)
	topic := "topic"
	channel := assertSubscribe(t, simulator, topic)
	messages := make([]Message, 0)

	go func() {
		assertClose(t, simulator)
	}()

	for message := range channel {
		messages = append(messages, message)
	}

	expected := []Message{}
	assertMessages(t, expected, messages)
}

func TestSimulatorPubSub(t *testing.T) {
	simulator := NewSimulator(logger)
	topic := "topic"

	channel1 := assertSubscribe(t, simulator, topic)
	channel2 := assertSubscribe(t, simulator, topic)
	messages1 := make([]Message, 0)
	messages2 := make([]Message, 0)

	var wg sync.WaitGroup
	wg.Add(2)

	expected := []Message{
		[]byte("one"),
		[]byte("two"),
	}

	go func() {
		for _, message := range expected {
			assertPublish(t, simulator, message, topic)
		}
		assertClose(t, simulator)
	}()

	go func() {
		for message := range channel1 {
			messages1 = append(messages1, message)
		}
		wg.Done()
	}()

	go func() {
		for message := range channel2 {
			messages2 = append(messages2, message)
		}
		wg.Done()
	}()

	wg.Wait()
	assertMessages(t, expected, messages1)
	assertMessages(t, expected, messages2)
}

func TestSimulatorPubSubMismatch(t *testing.T) {
	simulator := NewSimulator(logger)
	topic1 := "topic 1"
	topic2 := "topic 2"

	channel1 := assertSubscribe(t, simulator, topic1)
	channel2 := assertSubscribe(t, simulator, topic2)
	messages1 := make([]Message, 0)
	messages2 := make([]Message, 0)

	var wg sync.WaitGroup
	wg.Add(2)

	expected1 := []Message{
		[]byte("one"),
		[]byte("two"),
	}

	expected2 := []Message{}

	go func() {
		for _, message := range expected1 {
			assertPublish(t, simulator, message, topic1)
		}
		assertClose(t, simulator)
	}()

	go func() {
		for message := range channel1 {
			messages1 = append(messages1, message)
		}
		wg.Done()
	}()

	go func() {
		for message := range channel2 {
			messages2 = append(messages2, message)
		}
		wg.Done()
	}()

	wg.Wait()
	assertMessages(t, expected1, messages1)
	assertMessages(t, expected2, messages2)
}

func TestSimulatorPublishState(t *testing.T) {
	simulator := NewSimulator(logger)
	topic := "state"
	channel := assertSubscribe(t, simulator, topic)
	messages := make([]Message, 0)

	go func() {
		assertPublishState(t, simulator, false, false, topic)
		assertPublishState(t, simulator, true, false, topic)
		assertPublishState(t, simulator, false, true, topic)
		assertClose(t, simulator)
	}()

	for message := range channel {
		messages = append(messages, message)
	}

	expected := []Message{
		SIMULATOR_DEACTIVATED,
		SIMULATOR_ACTIVATED,
		SIMULATOR_DEACTIVATED,
	}
	assertMessages(t, expected, messages)
}

func TestSimulatorEnvironmental(t *testing.T) {
	simulator := NewSimulator(logger)
	topic := "environment"
	channel := assertSubscribe(t, simulator, topic)
	value := new(hypothalamus.PhysicEnv)
	value.Humidity = 1
	value.Pressure = 2
	value.Temperature = 3

	go func() {
		assertPublishEnvironmental(t, simulator, value, topic)
		assertClose(t, simulator)
	}()

	received := assertParseEnvironmental(t, simulator, <-channel)

	if received.RelativeHumidity() != value.RelativeHumidity() {
		t.Errorf("expected humidity %f but got %f", value.RelativeHumidity(), received.RelativeHumidity())
	}
	if received.MmHg() != value.MmHg() {
		t.Errorf("expected pressure %f but got %f", value.MmHg(), received.MmHg())
	}
	if received.Fahrenheit() != value.Fahrenheit() {
		t.Errorf("expected temperature %f but got %f", value.Fahrenheit(), received.Fahrenheit())
	}
}

func TestSimulatorDeviceState(t *testing.T) {
	simulator := NewSimulator(logger)
	topic := "state"
	channel := assertSubscribe(t, simulator, topic)

	go func() {
		assertPublishState(t, simulator, true, false, topic)
		assertClose(t, simulator)
	}()

	received := assertParseState(t, simulator, <-channel)

	if !received.IsActive() {
		t.Error("expected device state to be active")
	}
	if received.IsHalted() {
		t.Error("expected device state not to be halted")
	}
}
