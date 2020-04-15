package hypothalamus

import (
	"time"

	"periph.io/x/periph/conn/physic"
)

type Environmental interface {
	Fahrenheit() float64
	InHg() float64
	RelativeHumidity() float64
	Time() time.Time
}

type physicEnv struct {
	env  physic.Env
	time time.Time
}

func (measurement *physicEnv) Fahrenheit() float64 {
	return measurement.env.Temperature.Fahrenheit()
}

func (measurement *physicEnv) InHg() float64 {
	return float64(measurement.env.Pressure) / float64(physic.Pascal) / 3386.38816
}

func (measurement *physicEnv) RelativeHumidity() float64 {
	return float64(measurement.env.Humidity) / float64(physic.PercentRH)
}

func (measurement *physicEnv) Time() time.Time {
	return measurement.time
}
