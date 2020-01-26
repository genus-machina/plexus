package synapse

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/genus-machina/plexus/hypothalamus"
	"github.com/genus-machina/plexus/medulla"
)

type MQTT struct {
	channels []chan Message
	client   paho.Client
	logger   *log.Logger
}

type MQTTOptions struct {
	Broker   string
	CaFile   string
	CertFile string
	ClientId string
	KeyFile  string
}

func getStatusTopic(clientId string) string {
	return fmt.Sprintf("/devices/%s/status", clientId)
}

func publishBirthMessage(client paho.Client) {
	reader := client.OptionsReader()
	options := &reader
	statusTopic := getStatusTopic(options.ClientID())
	client.Publish(statusTopic, 1, true, "{\"status\": \"online\"}")
}

func NewMQTT(logger *log.Logger, options *MQTTOptions) (*MQTT, error) {
	caContents, err := ioutil.ReadFile(options.CaFile)
	if err != nil {
		return nil, err
	}

	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caContents)

	keyPair, err := tls.LoadX509KeyPair(options.CertFile, options.KeyFile)
	if err != nil {
		return nil, err
	}

	statusTopic := getStatusTopic(options.ClientId)

	tlsConfig := new(tls.Config)
	tlsConfig.Certificates = []tls.Certificate{keyPair}
	tlsConfig.RootCAs = caPool

	clientOptions := paho.NewClientOptions()
	clientOptions.AddBroker(options.Broker)
	clientOptions.SetAutoReconnect(true)
	clientOptions.SetCleanSession(true)
	clientOptions.SetClientID(options.ClientId)
	clientOptions.SetKeepAlive(time.Minute)
	clientOptions.SetOnConnectHandler(publishBirthMessage)
	clientOptions.SetTLSConfig(tlsConfig)
	clientOptions.SetWill(statusTopic, "{\"status\": \"offline\"}", 1, true)

	synapse := new(MQTT)
	synapse.client = paho.NewClient(clientOptions)

	token := synapse.client.Connect()
	token.Wait()

	if err := token.Error(); err != nil {
		return nil, err
	}

	synapse.logger = logger
	return synapse, nil
}

func (mqtt *MQTT) Apply(message Message, device medulla.Actuator) error {
	state, err := mqtt.ParseState(message)
	if err != nil {
		return err
	}

	if state.IsActive() {
		if err := device.Activate(); err == nil {
			mqtt.logger.Printf("Activated device '%s'.", device.Name())
			return nil
		} else {
			return err
		}
	} else {
		if err := device.Deactivate(); err == nil {
			mqtt.logger.Printf("Deactivated device '%s'.", device.Name())
			return nil
		} else {
			return err
		}
	}
}

func (mqtt *MQTT) Close() error {
	mqtt.logger.Printf("Closing synapse...")
	mqtt.client.Disconnect(250)

	for _, channel := range mqtt.channels {
		close(channel)
	}

	return nil
}

func (mqtt *MQTT) ParseEnvironmental(message Message) (hypothalamus.Environmental, error) {
	environmental := new(jsonEnvironmental)
	err := json.Unmarshal([]byte(message), environmental)
	return environmental, err
}

func (mqtt *MQTT) ParseState(message Message) (medulla.DeviceState, error) {
	state := new(jsonDeviceState)
	err := json.Unmarshal([]byte(message), state)
	return state, err
}

func (mqtt *MQTT) Publish(message Message, topic string) error {
	mqtt.logger.Printf("Publishing MQTT message to topic '%s'...", topic)
	mqtt.client.Publish(topic, 1, true, []byte(message))
	return nil
}

func (mqtt *MQTT) PublishEnvironmental(environmental hypothalamus.Environmental, topic string) error {
	message, err := json.Marshal(JsonEnvironmental(environmental))
	if err != nil {
		return err
	}

	return mqtt.Publish(message, topic)
}

func (mqtt *MQTT) PublishState(state medulla.DeviceState, topic string) error {
	message, err := json.Marshal(JsonDeviceState(state))
	if err != nil {
		return err
	}

	return mqtt.Publish(message, topic)
}

func (mqtt *MQTT) Subscribe(topic string) (<-chan Message, error) {
	mqtt.logger.Printf("Subscribing to topic '%s'...", topic)
	messages := make(chan Message)
	mqtt.channels = append(mqtt.channels, messages)

	mqtt.client.Subscribe(topic, 1, func(client paho.Client, message paho.Message) {
		messages <- Message(message.Payload())
		message.Ack()
	})

	return messages, nil
}
