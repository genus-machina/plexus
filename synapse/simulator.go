package synapse

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/genus-machina/plexus/hypothalamus"
	"github.com/genus-machina/plexus/medulla"
)

var (
	SIMULATOR_ACTIVATED   = []byte("activated")
	SIMULATOR_DEACTIVATED = []byte("deactivated")
)

type Simulator struct {
	logger        *log.Logger
	subscriptions map[string][]chan Message
}

func NewSimulator(logger *log.Logger) *Simulator {
	simulator := new(Simulator)
	simulator.logger = logger
	simulator.subscriptions = make(map[string][]chan Message)
	return simulator
}

func (simulator *Simulator) Apply(message Message, device medulla.Actuator) error {
	if bytes.Equal(message, SIMULATOR_ACTIVATED) {
		if err := device.Activate(); err == nil {
			simulator.logger.Printf("Activated device '%s'.", device.Name())
			return nil
		} else {
			return err
		}
	} else if bytes.Equal(message, SIMULATOR_DEACTIVATED) {
		if err := device.Deactivate(); err == nil {
			simulator.logger.Printf("Deactivated device '%s'.", device.Name())
			return nil
		} else {
			return err
		}
	} else {
		return nil
	}
}

func (simulator *Simulator) Close() error {
	simulator.logger.Println("Closing synapse...")
	for _, subscriptions := range simulator.subscriptions {
		for _, subscription := range subscriptions {
			close(subscription)
		}
	}
	return nil
}

func (simulator *Simulator) getSubscriptions(topic string) []chan Message {
	if subscriptions, ok := simulator.subscriptions[topic]; ok {
		return subscriptions
	} else {
		return make([]chan Message, 0)
	}
}

func (simulator *Simulator) ParseEnvironmental(message Message) (*hypothalamus.Environmental, error) {
	buffer := bytes.NewBuffer([]byte(message))
	value := new(hypothalamus.Environmental)

	if err := binary.Read(buffer, binary.BigEndian, value); err != nil {
		return nil, err
	}

	return value, nil
}

func (simulator *Simulator) ParseState(message Message) (medulla.DeviceState, error) {
	state := medulla.NewDeviceState(bytes.Equal(message, SIMULATOR_ACTIVATED), false)
	return state, nil
}

func (simulator *Simulator) Publish(message Message, topic string) error {
	simulator.logger.Printf("Publishing message '%s' to topic '%s'...", message, topic)
	for _, subscription := range simulator.getSubscriptions(topic) {
		subscription <- message
	}
	return nil
}

func (simulator *Simulator) PublishEnvironmental(environmental *hypothalamus.Environmental, topic string) error {
	buffer := new(bytes.Buffer)
	if err := binary.Write(buffer, binary.BigEndian, environmental); err != nil {
		return err
	}
	return simulator.Publish(Message(buffer.Bytes()), topic)
}

func (simulator *Simulator) PublishState(state medulla.DeviceState, topic string) error {
	if state.IsActive() {
		return simulator.Publish(SIMULATOR_ACTIVATED, topic)
	} else {
		return simulator.Publish(SIMULATOR_DEACTIVATED, topic)
	}
}

func (simulator *Simulator) Subscribe(topic string) (<-chan Message, error) {
	simulator.logger.Printf("Subscribing to topic '%s'...", topic)
	subscription := make(chan Message)
	subscriptions := simulator.getSubscriptions(topic)
	simulator.subscriptions[topic] = append(subscriptions, subscription)
	return subscription, nil
}
