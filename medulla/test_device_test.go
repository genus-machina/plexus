package medulla

type TestDevice struct {
	Active     bool
	DeviceName string
	Halted     bool
}

func NewTestDevice() *TestDevice {
	return new(TestDevice)
}

func (device *TestDevice) Halt() error {
	device.Halted = true
	return nil
}

func (device *TestDevice) Name() string {
	return device.DeviceName
}

func (device *TestDevice) State() DeviceState {
	return NewDeviceState(device.Active, device.Halted)
}
