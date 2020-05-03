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

type ScaledEnvironmental struct {
	source Environmental

	humidityCoefficient    float64
	humidityIntercept      float64
	pressureCoefficient    float64
	pressureIntercept      float64
	temperatureCoefficient float64
	temperatureIntercept   float64
}

func NewScaledEnvironmental(source Environmental) *ScaledEnvironmental {
	scaled := new(ScaledEnvironmental)
	scaled.source = source

	scaled.humidityCoefficient = 1
	scaled.pressureCoefficient = 1
	scaled.temperatureCoefficient = 1
	return scaled
}

func (scaled *ScaledEnvironmental) Fahrenheit() float64 {
	return scaled.source.Fahrenheit()*scaled.temperatureCoefficient + scaled.temperatureIntercept
}

func (scaled *ScaledEnvironmental) InHg() float64 {
	return scaled.source.InHg()*scaled.pressureCoefficient + scaled.pressureIntercept
}

func (scaled *ScaledEnvironmental) RelativeHumidity() float64 {
	return scaled.source.RelativeHumidity()*scaled.humidityCoefficient + scaled.humidityIntercept
}

func (scaled *ScaledEnvironmental) ScaleHumidity(coefficient, intercept float64) *ScaledEnvironmental {
	scaled.humidityCoefficient = coefficient
	scaled.humidityIntercept = intercept
	return scaled
}

func (scaled *ScaledEnvironmental) ScalePressure(coefficient, intercept float64) *ScaledEnvironmental {
	scaled.pressureCoefficient = coefficient
	scaled.pressureIntercept = intercept
	return scaled
}

func (scaled *ScaledEnvironmental) ScaleTemperature(coefficient, intercept float64) *ScaledEnvironmental {
	scaled.temperatureCoefficient = coefficient
	scaled.temperatureIntercept = intercept
	return scaled
}

func (scaled *ScaledEnvironmental) Time() time.Time {
	return scaled.source.Time()
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
