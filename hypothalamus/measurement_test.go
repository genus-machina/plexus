package hypothalamus

import (
	"testing"
	"time"

	"periph.io/x/periph/conn/physic"
)

func TestEnvironmentalFahrenheit(t *testing.T) {
	measurement := new(physicEnv)
	temperature := 98.6
	measurement.env.Temperature = physic.Temperature(temperature*float64(physic.Fahrenheit) + float64(physic.ZeroFahrenheit))

	if actual := measurement.Fahrenheit(); actual != temperature {
		t.Errorf("expected %f but got %f", temperature, actual)
	}
}

func TestEnvironmentalMmHg(t *testing.T) {
	measurement := new(physicEnv)
	pressure := 32.4
	measurement.env.Pressure = physic.Pressure(pressure * float64(physic.Pascal) * 3386.38816)

	if actual := measurement.MmHg(); actual != pressure {
		t.Errorf("expected %f but got %f", pressure, actual)
	}
}

func TestEnvironmentalRelativeHumidity(t *testing.T) {
	measurement := new(physicEnv)
	humidity := 12.34
	measurement.env.Humidity = physic.RelativeHumidity(humidity * float64(physic.PercentRH))

	if actual := measurement.RelativeHumidity(); actual != humidity {
		t.Errorf("expected %f but got %f", humidity, actual)
	}
}

func TestEnvironmentalTime(t *testing.T) {
	measurement := new(physicEnv)
	currentTime := time.Now()
	measurement.time = currentTime

	if actual := measurement.Time(); actual != currentTime {
		t.Errorf("expected %s but got %s", currentTime.Format(time.RFC3339), actual.Format(time.RFC3339))
	}
}
