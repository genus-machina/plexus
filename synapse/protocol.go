package synapse

import (
	"github.com/genus-machina/plexus/hypothalamus"
	"github.com/genus-machina/plexus/medulla"
)

type Message []byte

type Protocol interface {
	Apply(message Message, device medulla.Actuator) error
	Close() error
	ParseEnvironmental(message Message) (*hypothalamus.Environmental, error)
	ParseState(message Message) (medulla.DeviceState, error)
	Publish(message Message, topic string) error
	PublishEnvironmental(environmental *hypothalamus.Environmental, topic string) error
	PublishState(state medulla.DeviceState, topic string) error
	Subscribe(topic string) (<-chan Message, error)
}
