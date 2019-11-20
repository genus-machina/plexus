package zygote

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/genus-machina/plexus/hippocampus"
	"github.com/genus-machina/plexus/medulla"
	"github.com/genus-machina/plexus/medulla/actuators"
	"github.com/genus-machina/plexus/medulla/buses"
	"github.com/genus-machina/plexus/medulla/triggers"
)

type System struct {
	DeviceBus medulla.DeviceBus
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

func buildDevices(config *systemConfig) ([]medulla.Device, error) {
	devices := make([]medulla.Device, 0)

	for _, deviceConfig := range config.Devices {
		switch deviceConfig.Type {
		case "simulator":
			simulator, err := buildSimulator(deviceConfig)
			if err != nil {
				return nil, fmt.Errorf("Failed to build simulator '%s': %s", deviceConfig.Name, err.Error())
			}
			devices = append(devices, simulator)
		default:
			return nil, fmt.Errorf("Invalid device type '%s'.", deviceConfig.Type)
		}
	}

	return devices, nil
}

func buildSimulator(config *deviceConfig) (medulla.Device, error) {
	switch config.Driver {
	case "actuator":
		return actuators.NewSimulator(config.Name), nil
	case "trigger":
		return triggers.NewSimulator(config.Name), nil
	default:
		return nil, fmt.Errorf("Invalid simulator Driver '%s'.", config.Driver)
	}
}

func buildSystem(config *systemConfig, logger *log.Logger) (*System, error) {
	system := new(System)
	system.DeviceBus = buses.NewDeviceBus(hippocampus.ChildLogger(logger, "device bus"))

	devices, err := buildDevices(config)
	if err != nil {
		return system, err
	}

	for _, device := range devices {
		if err := system.DeviceBus.RegisterDevice(device); err != nil {
			return system, fmt.Errorf("Failed to register device '%s': %s", device.Name(), err.Error())
		}
	}

	return system, nil
}

func LoadJSON(path string, logger *log.Logger) (*System, error) {
	config, err := parseJSON(path)
	if err != nil {
		return nil, err
	}
	return buildSystem(config, logger)
}
