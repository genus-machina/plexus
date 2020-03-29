package synapse

import (
	"bytes"
	"testing"
	"time"

	"github.com/genus-machina/plexus/hypothalamus"
	"github.com/genus-machina/plexus/medulla"
)

func assertActive(t *testing.T, device medulla.Device) {
	if !device.State().IsActive() {
		t.Error("expected device to be active")
	}
}

func assertActivate(t *testing.T, device medulla.Actuator) {
	if err := device.Activate(); err != nil {
		t.Errorf("activation error: %s", err.Error())
	}
}

func assertClose(t *testing.T, simulator *Simulator) {
	if err := simulator.Close(); err != nil {
		t.Errorf("close error: %s", err.Error())
	}
}

func assertApply(t *testing.T, simulator *Simulator, message Message, device medulla.Actuator) {
	if err := simulator.Apply(message, device); err != nil {
		t.Errorf("apply error: %s", err.Error())
	}
}

func assertInactive(t *testing.T, device medulla.Device) {
	if device.State().IsActive() {
		t.Error("expected device to be inactive")
	}
}

func assertMessages(t *testing.T, expected []Message, actual []Message) {
	actualCount := len(actual)
	expectedCount := len(expected)

	if actualCount != expectedCount {
		t.Errorf("expected %d messages but got %d", expectedCount, actualCount)
	}

	for i := range expected {
		if !bytes.Equal(actual[i], expected[i]) {
			t.Errorf("message %d: expected %s but got %s", i, expected[i], actual[1])
		}
	}
}

func assertParseEnvironmental(t *testing.T, simulator *Simulator, message Message) hypothalamus.Environmental {
	if value, err := simulator.ParseEnvironmental(message); err == nil {
		return value
	} else {
		t.Errorf("error parsing environmental: %s", err.Error())
	}
	return nil
}

func assertParseState(t *testing.T, simulator *Simulator, message Message) medulla.DeviceState {
	if value, err := simulator.ParseState(message); err == nil {
		return value
	} else {
		t.Errorf("error parsing state: %s", err.Error())
	}
	return nil
}

func assertPublish(t *testing.T, simulator *Simulator, message Message, topic string) {
	if err := simulator.Publish(message, topic); err != nil {
		t.Errorf("publish error: %s", err.Error())
	}
}

func assertPublishEnvironmental(t *testing.T, simulator *Simulator, environmental hypothalamus.Environmental, topic string) {
	if err := simulator.PublishEnvironmental(environmental, topic); err != nil {
		t.Errorf("publish environmental error: %s", err.Error())
	}
}

func assertPublishState(t *testing.T, simulator *Simulator, active bool, halted bool, topic string) {
	state := medulla.NewDeviceState(active, halted, time.Unix(0, 0))
	if err := simulator.PublishState(state, topic); err != nil {
		t.Errorf("publish state error: %s", err.Error())
	}
}

func assertSubscribe(t *testing.T, simulator *Simulator, topic string) <-chan Message {
	if channel, err := simulator.Subscribe(topic); err == nil {
		return channel
	} else {
		t.Errorf("subscribe error: %s", err.Error())
	}
	return nil
}
