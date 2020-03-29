package synapse

import (
	"time"

	"github.com/genus-machina/plexus/hypothalamus"
	"github.com/genus-machina/plexus/medulla"
)

const (
	jsonTimeFormat = time.RFC3339
)

type jsonDeviceState struct {
	Active       bool   `json:"active"`
	Halted       bool   `json:"halted"`
	TimeRecorded string `json:"time"`
}

func JsonDeviceState(state medulla.DeviceState) *jsonDeviceState {
	json := new(jsonDeviceState)
	json.Active = state.IsActive()
	json.Halted = state.IsHalted()
	json.TimeRecorded = state.Time().Format(jsonTimeFormat)
	return json
}

func (json *jsonDeviceState) IsActive() bool {
	return json.Active
}

func (json *jsonDeviceState) IsHalted() bool {
	return json.Halted
}

func (json *jsonDeviceState) Time() time.Time {
	jsonTime, _ := time.Parse(jsonTimeFormat, json.TimeRecorded)
	return jsonTime
}

type jsonEnvironmental struct {
	Humidity     float64 `json:"humidity"`
	Pressure     float64 `json:"pressure"`
	Temperature  float64 `json:"temperature"`
	TimeRecorded string  `json:"time"`
}

func JsonEnvironmental(environmental hypothalamus.Environmental) *jsonEnvironmental {
	json := new(jsonEnvironmental)
	json.Humidity = environmental.RelativeHumidity()
	json.Pressure = environmental.MmHg()
	json.Temperature = environmental.Fahrenheit()
	json.TimeRecorded = environmental.Time().Format(jsonTimeFormat)
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

func (json *jsonEnvironmental) Time() time.Time {
	jsonTime, _ := time.Parse(jsonTimeFormat, json.TimeRecorded)
	return jsonTime
}

type jsonStatus struct {
	Status string `json:"status"`
}
