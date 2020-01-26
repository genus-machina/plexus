package synapse

import (
	"github.com/genus-machina/plexus/hypothalamus"
	"github.com/genus-machina/plexus/medulla"
)

type jsonDeviceState struct {
	Active bool `json:"active"`
	Halted bool `json:"halted"`
}

func JsonDeviceState(state medulla.DeviceState) *jsonDeviceState {
	json := new(jsonDeviceState)
	json.Active = state.IsActive()
	json.Halted = state.IsHalted()
	return json
}

func (json *jsonDeviceState) IsActive() bool {
	return json.Active
}

func (json *jsonDeviceState) IsHalted() bool {
	return json.Halted
}

type jsonEnvironmental struct {
	Humidity    float64 `json:"humidity"`
	Pressure    float64 `json:"pressure"`
	Temperature float64 `json:"temperature"`
}

func JsonEnvironmental(environmental hypothalamus.Environmental) *jsonEnvironmental {
	json := new(jsonEnvironmental)
	json.Humidity = environmental.RelativeHumidity()
	json.Pressure = environmental.MmHg()
	json.Temperature = environmental.Fahrenheit()
	return json
}

func (json *jsonEnvironmental) Fahrenheit() float64 {
	return json.Temperature
}

func (json *jsonEnvironmental) MmHg() float64 {
	return json.Pressure
}

func (json *jsonEnvironmental) RelativeHumidity() float64 {
	return json.Humidity
}
