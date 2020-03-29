package zygote

type deviceConfig struct {
	Name         string `json:"name"`
	Pin          string `json:"pin,omitempty"`
	Type         string `json:"type"`
	Driver       string `json:"driver,omitempty"`
	CommandTopic string `json:"commandTopic,omitempty"`
	StatusTopic  string `json:"statusTopic,omitempty"`
}

type environmentalConfig struct {
	Period      int    `json:"period"`
	StatusTopic string `json:"statusTopic,omitempty"`
}

type synapseConfig struct {
	Broker   string `json:"broker,omitempty"`
	CA       string `json:"ca,omitempty"`
	Cert     string `json:"cert,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	Key      string `json:"key,omitempty"`
	Type     string `json:"type"`
}

type systemConfig struct {
	EnvironmentalSensor *environmentalConfig `json:"environmentalSensor"`
	Devices             []*deviceConfig      `json:"devices"`
	Screen              bool                 `json:"screen"`
	Synapse             *synapseConfig       `json:"synapse"`
}
