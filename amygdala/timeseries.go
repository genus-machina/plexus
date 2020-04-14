package amygdala

import (
	"encoding/json"
	"image"
	"image/color"
	"image/draw"
	"math"
	"os"
	"sort"
	"time"
)

type TimeSeriesPoint struct {
	Time  time.Time
	Value float64
}

type timeSeriesPointComparison func(i, j int) bool

type TimeSeries []TimeSeriesPoint

func LoadTimeSeries(path string) (TimeSeries, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var series TimeSeries
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&series)
	return series, err
}

func (series TimeSeries) at(t time.Time) *TimeSeriesPoint {
	length := len(series)

	if len(series) < 1 {
		return nil
	}

	index := sort.Search(length, func(i int) bool { return t.Before(series[i].Time) })
	next := int(math.Min(float64(length-1), float64(index)))
	previous := int(math.Max(float64(0), float64(index-1)))
	since := t.Sub(series[previous].Time)
	until := series[next].Time.Sub(t)

	if since < until {
		return &series[previous]
	}

	return &series[next]
}

func (series TimeSeries) find(comparator timeSeriesPointComparison) *TimeSeriesPoint {
	index := -1

	for i := len(series); i > 0; i-- {
		if index < 0 || comparator(i-1, index) {
			index = i - 1
		}
	}

	if index < 0 {
		return nil
	}

	return &series[index]
}

func (series TimeSeries) First() *TimeSeriesPoint {
	return series.find(series.timeIsLessThan)
}

func (series TimeSeries) inRange(i, j int) bool {
	length := len(series)
	return i >= 0 && i < length && j >= 0 && j < length
}

func (series TimeSeries) Interpolate(size int) TimeSeries {
	length := len(series)
	result := make(TimeSeries, 0, size)

	if length > 0 {
		if size > 0 {
			series.Sort()
			result = append(result, series[0])
		}

		if size > 1 {
			duration := series[length-1].Time.Sub(series[0].Time)
			partitions := size - 1
			interval := duration / time.Duration(partitions)

			for i := 1; i < size; i++ {
				targetTime := series[0].Time.Add(time.Duration(i) * interval)
				nearestNeighbor := series.at(targetTime)
				result = append(result, TimeSeriesPoint{targetTime, nearestNeighbor.Value})
			}
		}
	}

	return result
}

func (series TimeSeries) Last() *TimeSeriesPoint {
	return series.find(series.timeIsGreaterThan)
}

func (series TimeSeries) Max() *TimeSeriesPoint {
	return series.find(series.valueIsGreaterThan)
}

func (series TimeSeries) Min() *TimeSeriesPoint {
	return series.find(series.valueIsLessThan)
}

func (series TimeSeries) Plot() *TimeSeriesPlot {
	widget := new(TimeSeriesPlot)
	widget.series = series
	return widget
}

func (series TimeSeries) Serialize(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(series)
}

func (series TimeSeries) Sort() {
	sort.Slice(series, series.timeIsLessThan)
}

func (series TimeSeries) timeIsGreaterThan(i, j int) bool {
	return series.inRange(i, j) && series[i].Time.After(series[j].Time)
}

func (series TimeSeries) timeIsLessThan(i, j int) bool {
	return series.inRange(i, j) && series[i].Time.Before(series[j].Time)
}

func (series TimeSeries) valueIsGreaterThan(i, j int) bool {
	return series.inRange(i, j) && series[i].Value > series[j].Value
}

func (series TimeSeries) valueIsLessThan(i, j int) bool {
	return series.inRange(i, j) && series[i].Value < series[j].Value
}

func (series TimeSeries) Zoom(start, end time.Time) TimeSeries {
	var beginning, ending time.Time
	result := make(TimeSeries, 0, 0)

	if start.Before(end) {
		beginning = start
		ending = end
	} else {
		beginning = end
		ending = start
	}

	for _, point := range series {
		if (point.Time.Equal(beginning) || point.Time.After(beginning)) && point.Time.Before(ending) {
			result = append(result, point)
		}
	}

	return result
}

type TimeSeriesPlot struct {
	bounds image.Rectangle
	fill   bool
	series TimeSeries
}

func (widget *TimeSeriesPlot) Bounds() image.Rectangle {
	return widget.bounds
}

func (widget *TimeSeriesPlot) drawLine(canvas draw.Image, x, y1, y2 int) {
	end := int(math.Max(float64(y1), float64(y2)))
	start := int(math.Min(float64(y1), float64(y2)))

	for y := start; y <= end; y++ {
		canvas.Set(x, y, color.White)
	}
}

func (widget *TimeSeriesPlot) Render(canvas draw.Image) {
	if len(widget.series) < 1 {
		return
	}

	height := widget.bounds.Dy()
	width := widget.bounds.Dx()

	data := widget.series.Interpolate(width)
	max := data.Max().Value
	min := data.Min().Value

	var scale float64
	valueRange := max - min
	if valueRange > 0 {
		scale = float64(height-1) / valueRange
	} else {
		scale = float64(height - 1)
	}

	previousY := 0
	for index, point := range data {
		value := int(math.Round((point.Value - min) * scale))
		x := widget.bounds.Min.X + index
		y := widget.bounds.Max.Y - value - 1

		if widget.fill {
			widget.drawLine(canvas, x, height-1, y)
		} else {
			if index > 0 {
				widget.drawLine(canvas, x, previousY, y)
			} else {
				canvas.Set(x, y, color.White)
			}
		}

		previousY = y
	}
}

func (widget *TimeSeriesPlot) SetBounds(bounds image.Rectangle) {
	widget.bounds = bounds
}

func (widget *TimeSeriesPlot) SetFill(fill bool) {
	widget.fill = fill
}
