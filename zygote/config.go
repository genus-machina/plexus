package zygote

type deviceConfig struct {
	Name   string `json:"name"`
	Pin    string `json:"pin,omitempty"`
	Type   string `json:"type"`
	Driver string `json:"driver,omitempty"`
	Topic  string `json:"topic,omitempty"`
}

type synapseConfig struct {
	Type string `json:"type"`
}

type systemConfig struct {
	EnvironmentalSensor bool            `json:"environmentalSensor"`
	Devices             []*deviceConfig `json:"devices"`
	Screen              bool            `json:"screen"`
	Synapse             *synapseConfig  `json:"synapse"`
}
