package hypothalamus

import (
	"periph.io/x/periph/conn/physic"
)

type Environmental struct {
	env physic.Env
}

func (measurement *Environmental) Fahrenheit() float64 {
	return measurement.env.Temperature.Fahrenheit()
}

func (measurement *Environmental) MmHg() float64 {
	return float64(measurement.env.Pressure) / float64(physic.Pascal) / 3386.38816
}

func (measurement *Environmental) RelativeHumidity() float64 {
	return float64(measurement.env.Humidity) / float64(physic.PercentRH)
}
