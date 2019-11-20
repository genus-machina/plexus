package zygote

type deviceConfig struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Driver string `json:"driver,omitempty"`
}

type systemConfig struct {
	Devices []*deviceConfig `json:"devices"`
}
