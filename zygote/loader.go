package zygote

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/genus-machina/plexus/amygdala"
	"github.com/genus-machina/plexus/hippocampus"
	"github.com/genus-machina/plexus/hypothalamus"
	"github.com/genus-machina/plexus/medulla"
	"github.com/genus-machina/plexus/medulla/actuators"
	"github.com/genus-machina/plexus/medulla/bus"
	"github.com/genus-machina/plexus/medulla/triggers"
	"github.com/genus-machina/plexus/synapse"

	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/devices/ssd1306"
)

type System struct {
	EnvironmentalSensor hypothalamus.Sensor
	DeviceBus           *bus.DeviceBus
	Screen              *amygdala.Screen
	Synapse             synapse.Protocol
}

func (system *System) Halt() {
	if system == nil {
		return
	}

	if system.DeviceBus != nil {
		system.DeviceBus.Halt()
	}

	if system.EnvironmentalSensor != nil {
		system.EnvironmentalSensor.Halt()
	}

	if system.Synapse != nil {
		system.Synapse.Close()
	}
}

func applyMessages(protocol synapse.Protocol, device medulla.Actuator, messages <-chan synapse.Message, statusTopic string) {
	for message := range messages {
		if err := protocol.Apply(message, device); err == nil && statusTopic != "" {
			protocol.Publish(message, statusTopic)
		}
	}
}

func bindActuator(synapse synapse.Protocol, device medulla.Actuator, commandTopic, statusTopic string) error {
	if messages, err := synapse.Subscribe(commandTopic); err == nil {
		go applyMessages(synapse, device, messages, statusTopic)
		return nil
	} else {
		return err
	}
}

func bindTrigger(synapse synapse.Protocol, device medulla.Trigger, statusTopic string) error {
	if states, err := device.Subscribe(); err == nil {
		go publishStates(synapse, states, statusTopic)
		return nil
	} else {
		return err
	}
}

func buildButton(config *deviceConfig, synapse synapse.Protocol) (medulla.Device, error) {
	if config.Pin == "" {
		return nil, errors.New("A GPIO pin is required.")
	}

	device, err := triggers.NewButton(config.Name, gpioreg.ByName(config.Pin))
	if err != nil {
		return nil, err
	}

	if !(config.StatusTopic == "" || synapse == nil) {
		if err := bindTrigger(synapse, device, config.StatusTopic); err != nil {
			return nil, err
		}
	}
	return device, nil
}

func buildDevices(config *systemConfig, synapse synapse.Protocol) ([]medulla.Device, error) {
	devices := make([]medulla.Device, 0)

	for _, deviceConfig := range config.Devices {
		var device medulla.Device

		switch deviceConfig.Type {
		case "button":
			var err error
			if device, err = buildButton(deviceConfig, synapse); err != nil {
				return nil, fmt.Errorf("Failed to build button '%s': %s", deviceConfig.Name, err.Error())
			}
		case "lamp":
			var err error
			if device, err = buildLamp(deviceConfig, synapse); err != nil {
				return nil, fmt.Errorf("Failed to build lamp '%s': %s", deviceConfig.Name, err.Error())
			}
		case "led":
			var err error
			if device, err = buildLED(deviceConfig, synapse); err != nil {
				return nil, fmt.Errorf("Failed to build LED '%s': %s", deviceConfig.Name, err.Error())
			}
		case "phototransistor":
			var err error
			if device, err = buildPhototransistor(deviceConfig, synapse); err != nil {
				return nil, fmt.Errorf("Failed to build phototransistor '%s': %s", deviceConfig.Name, err.Error())
			}
		case "pir":
			var err error
			if device, err = buildPIR(deviceConfig, synapse); err != nil {
				return nil, fmt.Errorf("Failed to build PIR '%s': %s", deviceConfig.Name, err.Error())
			}
		case "simulator":
			var err error
			if device, err = buildSimulator(deviceConfig, synapse); err != nil {
				return nil, fmt.Errorf("Failed to build simulator '%s': %s", deviceConfig.Name, err.Error())
			}
		case "water":
			var err error
			if device, err = buildWater(deviceConfig, synapse); err != nil {
				return nil, fmt.Errorf("Failed to build water sensor '%s': %s", deviceConfig.Name, err.Error())
			}
		default:
			return nil, fmt.Errorf("Invalid device type '%s'.", deviceConfig.Type)
		}

		devices = append(devices, device)
	}

	return devices, nil
}

func buildEnvironmentalSensor(config *environmentalConfig, synapse synapse.Protocol) (hypothalamus.Sensor, error) {
	i2cBus, err := i2creg.Open("")
	if err != nil {
		return nil, err
	}

	sensor, err := hypothalamus.NewBME280(i2cBus, 0x76)
	if err != nil {
		return nil, err
	}

	if !(config.StatusTopic == "" || synapse == nil) {
		period := 10 * time.Minute
		if config.Period != 0 {
			period = time.Duration(config.Period) * time.Second
		}
		measurements := sensor.SenseContinuous(period)
		go publishMeasurements(synapse, measurements, config)
	}

	return sensor, nil
}

func buildLamp(config *deviceConfig, synapse synapse.Protocol) (medulla.Device, error) {
	if config.Pin == "" {
		return nil, errors.New("A GPIO pin is required.")
	}

	device := actuators.NewLamp(config.Name, gpioreg.ByName(config.Pin))
	if !(config.CommandTopic == "" || synapse == nil) {
		if err := bindActuator(synapse, device, config.CommandTopic, config.StatusTopic); err != nil {
			return nil, err
		}
	}
	return device, nil
}

func buildLED(config *deviceConfig, synapse synapse.Protocol) (medulla.Device, error) {
	if config.Pin == "" {
		return nil, errors.New("A GPIO pin is required.")
	}

	device := actuators.NewLED(config.Name, gpioreg.ByName(config.Pin))
	if !(config.CommandTopic == "" || synapse == nil) {
		if err := bindActuator(synapse, device, config.CommandTopic, config.StatusTopic); err != nil {
			return nil, err
		}
	}
	return device, nil
}

func buildPhototransistor(config *deviceConfig, synapse synapse.Protocol) (medulla.Device, error) {
	if config.Pin == "" {
		return nil, errors.New("A GPIO pin is required.")
	}

	device, err := triggers.NewPhototransistor(config.Name, gpioreg.ByName(config.Pin))
	if err != nil {
		return nil, err
	}

	if !(config.StatusTopic == "" || synapse == nil) {
		if err := bindTrigger(synapse, device, config.StatusTopic); err != nil {
			return nil, err
		}
	}
	return device, nil
}

func buildPIR(config *deviceConfig, synapse synapse.Protocol) (medulla.Device, error) {
	if config.Pin == "" {
		return nil, errors.New("A GPIO pin is required.")
	}

	device, err := triggers.NewPIR(config.Name, gpioreg.ByName(config.Pin))
	if err != nil {
		return nil, err
	}

	if !(config.StatusTopic == "" || synapse == nil) {
		if err := bindTrigger(synapse, device, config.StatusTopic); err != nil {
			return nil, err
		}
	}
	return device, nil
}

func buildScreen() (*amygdala.Screen, error) {
	i2cBus, err := i2creg.Open("")
	if err != nil {
		return nil, err
	}

	if device, err := ssd1306.NewI2C(i2cBus, &ssd1306.DefaultOpts); err == nil {
		return amygdala.NewScreen(device), nil
	} else {
		return nil, err
	}
}

func buildSimulator(config *deviceConfig, synapse synapse.Protocol) (medulla.Device, error) {
	switch config.Driver {
	case "actuator":
		device := actuators.NewSimulator(config.Name)
		if !(config.CommandTopic == "" || synapse == nil) {
			if err := bindActuator(synapse, device, config.CommandTopic, config.StatusTopic); err != nil {
				return nil, err
			}
		}
		return device, nil
	case "trigger":
		device := triggers.NewSimulator(config.Name)
		if !(config.StatusTopic == "" || synapse == nil) {
			if err := bindTrigger(synapse, device, config.StatusTopic); err != nil {
				return nil, err
			}
		}
		return device, nil
	default:
		return nil, fmt.Errorf("Invalid simulator driver '%s'.", config.Driver)
	}
}

func buildSynapse(config *synapseConfig, logger *log.Logger) (synapse.Protocol, error) {
	synapticLogger := hippocampus.ChildLogger(logger, "synapse")

	switch config.Type {
	case "mqtt":
		options := &synapse.MQTTOptions{
			Broker:   config.Broker,
			CaFile:   config.CA,
			CertFile: config.Cert,
			ClientId: config.ClientId,
			KeyFile:  config.Key,
		}
		return synapse.NewMQTT(synapticLogger, options)
	case "simulator":
		return synapse.NewSimulator(synapticLogger), nil
	default:
		return nil, fmt.Errorf("Invalid synapse type '%s'.", config.Type)
	}
}

func buildSystem(config *systemConfig, logger *log.Logger) (*System, error) {
	system := new(System)

	if config.Synapse != nil {
		if synapse, err := buildSynapse(config.Synapse, logger); err == nil {
			system.Synapse = synapse
		} else {
			return nil, err
		}
	}

	system.DeviceBus = bus.New(hippocampus.ChildLogger(logger, "device bus"))

	devices, err := buildDevices(config, system.Synapse)
	if err != nil {
		return system, err
	}

	for _, device := range devices {
		if err := system.DeviceBus.RegisterDevice(device); err != nil {
			return system, fmt.Errorf("Failed to register device '%s': %s", device.Name(), err.Error())
		}
	}

	if config.Screen {
		if screen, err := buildScreen(); err == nil {
			system.Screen = screen
		} else {
			return nil, fmt.Errorf("Failed to initialize screen: %s", err.Error())
		}
	}

	if config.EnvironmentalSensor != nil {
		if sensor, err := buildEnvironmentalSensor(config.EnvironmentalSensor, system.Synapse); err == nil {
			system.EnvironmentalSensor = sensor
		} else {
			return nil, fmt.Errorf("Failed to initialize environmental sensor: %s", err.Error())
		}
	}

	return system, nil
}

func buildWater(config *deviceConfig, synapse synapse.Protocol) (medulla.Device, error) {
	if config.Pin == "" {
		return nil, errors.New("A GPIO pin is required.")
	}

	device, err := triggers.NewWater(config.Name, gpioreg.ByName(config.Pin))
	if err != nil {
		return nil, err
	}

	if !(config.StatusTopic == "" || synapse == nil) {
		if err := bindTrigger(synapse, device, config.StatusTopic); err != nil {
			return nil, err
		}
	}
	return device, nil
}

func LoadJSON(path string, logger *log.Logger) (*System, error) {
	config, err := parseJSON(path)
	if err != nil {
		return nil, err
	}
	return buildSystem(config, logger)
}

func parseJSON(path string) (*systemConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config systemConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func publishMeasurements(protocol synapse.Protocol, measurements <-chan hypothalamus.Environmental, config *environmentalConfig) {
	for measurement := range measurements {
		scaled := hypothalamus.NewScaledEnvironmental(measurement)
		if config.HumidityCoefficient != 0 {
			scaled.ScaleHumidity(config.HumidityCoefficient, config.HumidityIntercept)
		}
		if config.PressureCoefficient != 0 {
			scaled.ScalePressure(config.PressureCoefficient, config.PressureIntercept)
		}
		if config.TemperatureCoefficient != 0 {
			scaled.ScaleTemperature(config.TemperatureCoefficient, config.TemperatureIntercept)
		}
		protocol.PublishEnvironmental(scaled, config.StatusTopic)
	}
}

func publishStates(protocol synapse.Protocol, states <-chan medulla.DeviceState, topic string) {
	for state := range states {
		protocol.PublishState(state, topic)
	}
}
