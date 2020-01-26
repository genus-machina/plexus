package hypothalamus

import (
	"periph.io/x/periph/conn/physic"
)

type Environmental interface {
	Fahrenheit() float64
	MmHg() float64
	RelativeHumidity() float64
}

type PhysicEnv physic.Env

func (measurement *PhysicEnv) Fahrenheit() float64 {
	return measurement.Temperature.Fahrenheit()
}

func (measurement *PhysicEnv) MmHg() float64 {
	return float64(measurement.Pressure) / float64(physic.Pascal) / 3386.38816
}

func (measurement *PhysicEnv) RelativeHumidity() float64 {
	return float64(measurement.Humidity) / float64(physic.PercentRH)
}
