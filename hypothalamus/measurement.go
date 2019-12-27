package hypothalamus

import (
	"periph.io/x/periph/conn/physic"
)

type Environmental physic.Env

func (measurement *Environmental) Fahrenheit() float64 {
	return measurement.Temperature.Fahrenheit()
}

func (measurement *Environmental) MmHg() float64 {
	return float64(measurement.Pressure) / float64(physic.Pascal) / 3386.38816
}

func (measurement *Environmental) RelativeHumidity() float64 {
	return float64(measurement.Humidity) / float64(physic.PercentRH)
}
