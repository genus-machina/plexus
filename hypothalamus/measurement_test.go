package hypothalamus

import (
	"testing"

	"periph.io/x/periph/conn/physic"
)

func TestEnvironmentalFahrenheit(t *testing.T) {
	measurement := new(Environmental)
	temperature := 98.6
	measurement.Temperature = physic.Temperature(temperature*float64(physic.Fahrenheit) + float64(physic.ZeroFahrenheit))

	if actual := measurement.Fahrenheit(); actual != temperature {
		t.Errorf("expected %f but got %f", temperature, actual)
	}
}

func TestEnvironmentalMmHg(t *testing.T) {
	measurement := new(Environmental)
	pressure := 32.4
	measurement.Pressure = physic.Pressure(pressure * float64(physic.Pascal) * 3386.38816)

	if actual := measurement.MmHg(); actual != pressure {
		t.Errorf("expected %f but got %f", pressure, actual)
	}
}

func TestEnvironmentalRelativeHumidity(t *testing.T) {
	measurement := new(Environmental)
	humidity := 12.34
	measurement.Humidity = physic.RelativeHumidity(humidity * float64(physic.PercentRH))

	if actual := measurement.RelativeHumidity(); actual != humidity {
		t.Errorf("expected %f but got %f", humidity, actual)
	}
}
