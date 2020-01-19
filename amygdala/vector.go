package amygdala

import (
	"image"
	"image/draw"
)

type bounder func(image.Point, image.Rectangle, int) image.Rectangle
type measurer func(image.Rectangle) int
type translator func(int) image.Point

type vector struct {
	bounds  image.Rectangle
	options []vectorOptions
	items   []Widget
}

type vectorOptions interface {
	IsFixed() bool
	IsValid() bool
}

func (v *vector) addItem(item Widget, options vectorOptions, measure measurer, translate translator) {
	diff := v.bounds.Min.Add(translate(measure(v.bounds))).Sub(item.Bounds().Min)
	item.SetBounds(item.Bounds().Add(diff))
	v.options = append(v.options, options)
	v.items = append(v.items, item)
}

func (v *vector) Bounds() image.Rectangle {
	return v.bounds
}

func (v *vector) Render(canvas draw.Image) {
	for _, item := range v.items {
		item.Render(canvas)
	}
}

func (v *vector) scale(bound bounder, measure measurer, translate translator) {
	items := len(v.items)

	if items < 1 {
		return
	}

	dynamicCount := 0
	fixedValue := 0
	values := make([]int, items)

	for i, item := range v.items {
		if v.options[i].IsValid() && v.options[i].IsFixed() {
			value := measure(item.Bounds())
			fixedValue = fixedValue + value
			values[i] = value
		} else {
			dynamicCount = dynamicCount + 1
		}
	}

	unallocated := measure(v.bounds) - fixedValue
	value := unallocated
	padded := 0

	if dynamicCount > 0 {
		value = unallocated / dynamicCount
		padded = unallocated % dynamicCount
	}

	for i := range values {
		if v.options[i].IsValid() && v.options[i].IsFixed() {
			continue
		}

		if i < padded {
			values[i] = value + 1
		} else {
			values[i] = value
		}
	}

	previous := v.bounds.Min
	for i, item := range v.items {
		item.SetBounds(bound(previous, v.bounds, values[i]))
		previous = previous.Add(translate(values[i]))
	}
}

func (v *vector) updateBounds() {
	last := len(v.items) - 1
	v.bounds = image.Rectangle{
		Min: v.items[0].Bounds().Min,
		Max: v.items[last].Bounds().Max,
	}
}
