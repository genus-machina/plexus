package amygdala

import (
	"image"
	"testing"
	"time"
)

func assertFalse(t *testing.T, value bool) {
	if value {
		t.Errorf("expected value to be false but was true")
	}
}

func assertFloatEquals(t *testing.T, actual, expected float64) {
	if actual != expected {
		t.Errorf("expected %f but got %f", expected, actual)
	}
}

func assertIntEquals(t *testing.T, actual, expected int) {
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}

func assertRectangle(t *testing.T, expected, actual image.Rectangle) {
	if !actual.Eq(expected) {
		t.Errorf("expected rectangle %s but got %s", expected, actual)
	}
}

func assertScreenRender(t *testing.T, screen *Screen, content Widget) {
	if err := screen.Render(content); err != nil {
		t.Errorf("failed to render: %s", err.Error())
	}
}

func assertSeriesEquals(t *testing.T, actual, expected TimeSeries) {
	actualLen := len(actual)
	expectedLen := len(expected)

	if actualLen != expectedLen {
		t.Errorf("expected series to have length %d but was %d", expectedLen, actualLen)
	}

	for i := 0; i < expectedLen; i++ {
		actualPoint := actual[i]
		expectedPoint := expected[i]

		if !actualPoint.Time.Equal(expectedPoint.Time) {
			t.Errorf(
				"expected time %d at position %d but got %d",
				expectedPoint.Time.UnixNano(),
				i,
				actualPoint.Time.UnixNano(),
			)
		}

		if actualPoint.Value != expectedPoint.Value {
			t.Errorf(
				"expected value %f at position %d but got %f",
				expectedPoint.Value,
				i,
				actualPoint.Value,
			)
		}
	}
}

func assertTimeEquals(t *testing.T, actual, expected time.Time) {
	if !actual.Equal(expected) {
		t.Errorf("expected %d but got %d", expected.Unix(), actual.Unix())
	}
}

func assertTrue(t *testing.T, value bool) {
	if !value {
		t.Errorf("expected value to be true but was false")
	}
}
