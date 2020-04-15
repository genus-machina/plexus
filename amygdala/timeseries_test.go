package amygdala

import (
	"image"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestTimeSeriesFirst(t *testing.T) {
	series := make(TimeSeries, 0, 0)
	if series.First() != nil {
		t.Error("expected value to be nil")
	}

	series = append(series, TimeSeriesPoint{time.Unix(10, 0), 10})
	series = append(series, TimeSeriesPoint{time.Unix(0, 0), 0})

	assertTimeEquals(t, series.First().Time, time.Unix(0, 0))
	assertFloatEquals(t, series.First().Value, 0)
}

func TestTimeSeriesInterpolate(t *testing.T) {
	series := make(TimeSeries, 0, 0)
	assertSeriesEquals(t, series.Interpolate(5), TimeSeries{})

	series = append(series, TimeSeriesPoint{time.Unix(12, 0), 10})
	series = append(series, TimeSeriesPoint{time.Unix(6, 0), 20})
	series = append(series, TimeSeriesPoint{time.Unix(18, 0), 0})

	assertSeriesEquals(t, series.Interpolate(0), TimeSeries{})

	assertSeriesEquals(
		t,
		series.Interpolate(1),
		TimeSeries{
			TimeSeriesPoint{time.Unix(6, 0), 20},
		},
	)

	assertSeriesEquals(
		t,
		series.Interpolate(2),
		TimeSeries{
			TimeSeriesPoint{time.Unix(6, 0), 20},
			TimeSeriesPoint{time.Unix(18, 0), 0},
		},
	)

	assertSeriesEquals(
		t,
		series.Interpolate(3),
		TimeSeries{
			TimeSeriesPoint{time.Unix(6, 0), 20},
			TimeSeriesPoint{time.Unix(12, 0), 10},
			TimeSeriesPoint{time.Unix(18, 0), 0},
		},
	)

	assertSeriesEquals(
		t,
		series.Interpolate(7),
		TimeSeries{
			TimeSeriesPoint{time.Unix(6, 0), 20},
			TimeSeriesPoint{time.Unix(8, 0), 20},
			TimeSeriesPoint{time.Unix(10, 0), 10},
			TimeSeriesPoint{time.Unix(12, 0), 10},
			TimeSeriesPoint{time.Unix(14, 0), 10},
			TimeSeriesPoint{time.Unix(16, 0), 0},
			TimeSeriesPoint{time.Unix(18, 0), 0},
		},
	)
}

func TestTimeSeriesLast(t *testing.T) {
	series := make(TimeSeries, 0, 0)
	if series.First() != nil {
		t.Error("expected value to be nil")
	}

	series = append(series, TimeSeriesPoint{time.Unix(10, 0), 10})
	series = append(series, TimeSeriesPoint{time.Unix(0, 0), 0})

	assertTimeEquals(t, series.Last().Time, time.Unix(10, 0))
	assertFloatEquals(t, series.Last().Value, 10)
}

func TestTimeSeriesMax(t *testing.T) {
	series := make(TimeSeries, 0, 0)
	if series.Max() != nil {
		t.Error("expected value to be nil")
	}

	series = append(series, TimeSeriesPoint{time.Unix(1, 0), 0})
	series = append(series, TimeSeriesPoint{time.Unix(0, 0), 1})

	max := series.Max()
	assertTimeEquals(t, max.Time, series[1].Time)
	assertFloatEquals(t, max.Value, series[1].Value)
}

func TestTimeSeriesMean(t *testing.T) {
	series := make(TimeSeries, 0, 0)
	assertFloatEquals(t, series.Mean(), 0)

	series = append(series, TimeSeriesPoint{time.Unix(0, 0), 1})
	series = append(series, TimeSeriesPoint{time.Unix(1, 0), 3})
	assertFloatEquals(t, series.Mean(), 2)
}

func TestTimeSeriesMedian(t *testing.T) {
	series := make(TimeSeries, 0, 0)
	assertFloatEquals(t, series.Median(), 0)

	series = append(series, TimeSeriesPoint{time.Unix(0, 0), 1})
	series = append(series, TimeSeriesPoint{time.Unix(2, 0), 2})
	series = append(series, TimeSeriesPoint{time.Unix(1, 0), 3})
	assertFloatEquals(t, series.Median(), 2)

	series = append(series, TimeSeriesPoint{time.Unix(3, 0), 4})
	assertFloatEquals(t, series.Median(), 2.5)
}

func TestTimeSeriesMin(t *testing.T) {
	series := make(TimeSeries, 0, 0)
	if series.Min() != nil {
		t.Error("expected value to be nil")
	}

	series = append(series, TimeSeriesPoint{time.Unix(1, 0), 0})
	series = append(series, TimeSeriesPoint{time.Unix(0, 0), 1})

	min := series.Min()
	assertTimeEquals(t, min.Time, series[0].Time)
	assertFloatEquals(t, min.Value, series[0].Value)
}

func TestTimeSeriesSerialize(t *testing.T) {
	series := make(TimeSeries, 0, 0)
	series = append(series, TimeSeriesPoint{time.Unix(0, 0), 0})
	series = append(series, TimeSeriesPoint{time.Unix(1, 0), 1})

	file := t.Name() + ".json"

	if err := series.Serialize(file); err != nil {
		t.Errorf("Failed to serialize time series. %s.", err.Error())
	}

	if serialized, err := LoadTimeSeries(file); err != nil {
		t.Errorf("Failed to load time series. %s.", err.Error())
	} else {
		assertSeriesEquals(t, serialized, series)
	}
}

func TestTimeSeriesSort(t *testing.T) {
	series := make(TimeSeries, 0, 0)

	series.Sort()
	assertSeriesEquals(t, series, TimeSeries{})

	series = append(series, TimeSeriesPoint{time.Unix(1, 0), 0})
	series = append(series, TimeSeriesPoint{time.Unix(0, 0), 1})

	series.Sort()
	assertSeriesEquals(
		t,
		series,
		TimeSeries{
			TimeSeriesPoint{time.Unix(0, 0), 1},
			TimeSeriesPoint{time.Unix(1, 0), 0},
		},
	)
}

func TestTimeSeriesZoom(t *testing.T) {
	series := make(TimeSeries, 0, 0)
	assertSeriesEquals(t, series, series.Zoom(time.Unix(0, 0), time.Unix(1, 0)))

	series = append(series, TimeSeriesPoint{time.Unix(0, 0), 0})
	series = append(series, TimeSeriesPoint{time.Unix(1, 0), 1})
	series = append(series, TimeSeriesPoint{time.Unix(2, 0), 2})

	assertSeriesEquals(
		t,
		series.Zoom(time.Unix(0, 0), time.Unix(1, 0)),
		TimeSeries{
			TimeSeriesPoint{time.Unix(0, 0), 0},
		},
	)

	assertSeriesEquals(
		t,
		series.Zoom(time.Unix(1, 0), time.Unix(0, 0)),
		TimeSeries{
			TimeSeriesPoint{time.Unix(0, 0), 0},
		},
	)

	assertSeriesEquals(
		t,
		series.Zoom(time.Unix(1, 0), time.Unix(3, 0)),
		TimeSeries{
			TimeSeriesPoint{time.Unix(1, 0), 1},
			TimeSeriesPoint{time.Unix(2, 0), 2},
		},
	)
}

func TestTimeSeriesPlotBounds(t *testing.T) {
	series := make(TimeSeries, 0, 0)
	widget := series.Plot()
	assertRectangle(t, image.Rect(0, 0, 0, 0), widget.Bounds())

	bounds := image.Rect(1, 2, 3, 4)
	widget.SetBounds(bounds)
	assertRectangle(t, bounds, widget.Bounds())
}

func RenderSeries(t *testing.T, series TimeSeries) {
	canvas := image.NewNRGBA(image.Rect(0, 0, 128, 64))
	plot := series.Plot()
	plot.SetBounds(canvas.Bounds())
	plot.Render(canvas)
	saveImage(t, "plain", canvas)

	canvas = image.NewNRGBA(image.Rect(0, 0, 128, 64))
	plot = series.Plot()
	plot.SetBounds(canvas.Bounds())
	plot.SetFill(true)
	plot.Render(canvas)
	saveImage(t, "filled", canvas)

	canvas = image.NewNRGBA(image.Rect(0, 0, 128, 64))
	plot = series.Plot()
	plot.SetBounds(canvas.Bounds())
	plot.SetCenter(CenterFirst)
	plot.SetFill(true)
	plot.Render(canvas)
	saveImage(t, "first", canvas)

	canvas = image.NewNRGBA(image.Rect(0, 0, 128, 64))
	plot = series.Plot()
	plot.SetBounds(canvas.Bounds())
	plot.SetCenter(CenterLast)
	plot.SetFill(true)
	plot.Render(canvas)
	saveImage(t, "last", canvas)

	canvas = image.NewNRGBA(image.Rect(0, 0, 128, 64))
	plot = series.Plot()
	plot.SetBounds(canvas.Bounds())
	plot.SetCenter(CenterMean)
	plot.SetFill(true)
	plot.Render(canvas)
	saveImage(t, "mean", canvas)

	canvas = image.NewNRGBA(image.Rect(0, 0, 128, 64))
	plot = series.Plot()
	plot.SetBounds(canvas.Bounds())
	plot.SetCenter(CenterMedian)
	plot.SetFill(true)
	plot.Render(canvas)
	saveImage(t, "median", canvas)
}

func TestTimeSeriesPlotRenderLine(t *testing.T) {
	series := make(TimeSeries, 0, 0)

	for i := 0; i < 256; i++ {
		series = append(series, TimeSeriesPoint{time.Unix(int64(i), 0), float64(i)})
	}

	RenderSeries(t, series)
}

func TestTimeSeriesPlotRenderCurve(t *testing.T) {
	series := make(TimeSeries, 0, 0)
	step := 2 * math.Pi / 128
	for i := 0; i < 128; i++ {
		series = append(series, TimeSeriesPoint{time.Unix(int64(i), 0), math.Sin(float64(i) * step)})
	}
	RenderSeries(t, series)
}

func TestTimeSeriesPlotRenderCurveNested(t *testing.T) {
	series := make(TimeSeries, 0, 0)
	step := 2 * math.Pi / 128
	for i := 0; i < 128; i++ {
		series = append(series, TimeSeriesPoint{time.Unix(int64(i), 0), math.Sin(float64(i) * step)})
	}

	column := NewColumn()
	face := NewFontFace(10)
	text := NewTextBox(face, "test")
	column.AppendRow(text, &RowOptions{Fixed: true})

	plot := series.Plot()
	plot.SetFill(true)
	column.AppendRow(plot, nil)

	canvas := image.NewNRGBA(image.Rect(0, 0, 128, 64))
	column.SetBounds(canvas.Bounds())
	column.Render(canvas)
	saveImage(t, "filled", canvas)

}

func TestTimeSeriesPlotRenderRandom(t *testing.T) {
	series := make(TimeSeries, 0, 0)
	for i := 0; i < 100; i++ {
		series = append(series, TimeSeriesPoint{time.Unix(int64(i), 0), rand.Float64()})
	}
	RenderSeries(t, series)
}

func TestTimeSeriesPlotRenderStep(t *testing.T) {
	series := make(TimeSeries, 0, 0)
	series = append(series, TimeSeriesPoint{time.Unix(0, 0), 0})
	series = append(series, TimeSeriesPoint{time.Unix(1, 0), 1})
	RenderSeries(t, series)
}

func TestTimeSeriesPlotRenderSteps(t *testing.T) {
	series := make(TimeSeries, 0, 0)
	for i := 0; i < 10; i++ {
		series = append(series, TimeSeriesPoint{time.Unix(int64(i), 0), float64(10 - i)})
	}
	RenderSeries(t, series)
}

func TestTimeSeriesPlotRenderWeight(t *testing.T) {
	series := make(TimeSeries, 0, 0)
	for i := 0; i < 10; i++ {
		series = append(series, TimeSeriesPoint{time.Unix(int64(i), 0), float64(i % 3)})
	}
	series = append(series, TimeSeriesPoint{time.Unix(11, 0), 20})
	RenderSeries(t, series)
}
