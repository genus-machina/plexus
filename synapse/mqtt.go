package synapse

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/genus-machina/plexus/hypothalamus"
	"github.com/genus-machina/plexus/medulla"
)

type MQTT struct {
	client        paho.Client
	logger        *log.Logger
	mutex         sync.Mutex
	subscriptions map[string][]chan Message
}

type MQTTOptions struct {
	Broker   string
	CaFile   string
	CertFile string
	ClientId string
	KeyFile  string
}

const (
	MQTT_QOS_AT_MOST_ONCE = iota
	MQTT_QOS_AT_LEAST_ONCE
	MQTT_QOS_EXACTLY_ONCE
)

const (
	defaultQos = MQTT_QOS_AT_MOST_ONCE
)

func getStatusTopic(clientId string) string {
	return fmt.Sprintf("/devices/%s/status", clientId)
}

func publishBirthMessage(client paho.Client) {
	reader := client.OptionsReader()
	options := &reader
	statusTopic := getStatusTopic(options.ClientID())
	client.Publish(statusTopic, MQTT_QOS_AT_LEAST_ONCE, true, "{\"status\": \"online\"}")
}

func connectionHandler(mqtt *MQTT) paho.OnConnectHandler {
	return func(client paho.Client) {
		mqtt.logger.Println("Connected to MQTT broker.")
		publishBirthMessage(client)
		mqtt.resubscribe()
	}
}

func connectionLostHandler(mqtt *MQTT) paho.ConnectionLostHandler {
	return func(client paho.Client, err error) {
		mqtt.logger.Printf("Lost connection to MQTT broker. %s.\n", err.Error())
	}
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

	synapse := new(MQTT)
	synapse.logger = logger
	synapse.subscriptions = make(map[string][]chan Message)

	statusTopic := getStatusTopic(options.ClientId)

	tlsConfig := new(tls.Config)
	tlsConfig.Certificates = []tls.Certificate{keyPair}
	tlsConfig.RootCAs = caPool

	clientOptions := paho.NewClientOptions()
	clientOptions.AddBroker(options.Broker)
	clientOptions.SetAutoReconnect(true)
	clientOptions.SetCleanSession(true)
	clientOptions.SetClientID(options.ClientId)
	clientOptions.SetConnectionLostHandler(connectionLostHandler(synapse))
	clientOptions.SetKeepAlive(time.Minute)
	clientOptions.SetOnConnectHandler(connectionHandler(synapse))
	clientOptions.SetTLSConfig(tlsConfig)
	clientOptions.SetWill(statusTopic, "{\"status\": \"offline\"}", MQTT_QOS_AT_LEAST_ONCE, true)

	synapse.client = paho.NewClient(clientOptions)
	synapse.client.Connect()
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

	for _, subscription := range mqtt.subscriptions {
		for _, channel := range subscription {
			close(channel)
		}
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
	mqtt.client.Publish(topic, defaultQos, true, []byte(message))
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

func (mqtt *MQTT) createSubscription(topic string) <-chan Message {
	mqtt.mutex.Lock()
	defer mqtt.mutex.Unlock()

	messages := make(chan Message)
	subscriptions := mqtt.getSubscriptions(topic)

	if len(subscriptions) == 0 {
		mqtt.subscribeToTopic(topic)
	}

	subscriptions = append(subscriptions, messages)
	mqtt.setSubscriptions(topic, subscriptions)
	return messages
}

func (mqtt *MQTT) getSubscriptions(topic string) []chan Message {
	return mqtt.subscriptions[topic]
}

func (mqtt *MQTT) setSubscriptions(topic string, subscriptions []chan Message) {
	mqtt.subscriptions[topic] = subscriptions
}

func (mqtt *MQTT) resubscribe() {
	mqtt.mutex.Lock()
	defer mqtt.mutex.Unlock()

	if len(mqtt.subscriptions) > 0 {
		mqtt.logger.Println("Resubscribing to all previously subscribed topics...")
		for topic, _ := range mqtt.subscriptions {
			mqtt.subscribeToTopic(topic)
		}
	} else {
		mqtt.logger.Println("No previously subscribed topics.")
	}
}

func (mqtt *MQTT) subscribeToTopic(topic string) {
	mqtt.client.Subscribe(topic, defaultQos, func(client paho.Client, message paho.Message) {
		mqtt.mutex.Lock()
		defer mqtt.mutex.Unlock()

		subscriptions := mqtt.getSubscriptions(topic)
		for _, channel := range subscriptions {
			channel <- Message(message.Payload())
		}
		message.Ack()
	})
}

func (mqtt *MQTT) Subscribe(topic string) (<-chan Message, error) {
	mqtt.logger.Printf("Subscribing to topic '%s'...", topic)
	messages := mqtt.createSubscription(topic)
	return messages, nil
}
