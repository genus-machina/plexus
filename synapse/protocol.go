package synapse

import (
	"github.com/genus-machina/plexus/medulla"
)

type Message []byte

type Protocol interface {
	Apply(message Message, device medulla.Actuator) error
	Close() error
	Publish(message Message, topic string) error
	PublishState(state medulla.DeviceState, topic string) error
	Subscribe(topic string) (<-chan Message, error)
}
