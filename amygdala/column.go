package amygdala

import (
	"image"
)

type Column struct {
	vector
}

type RowOptions struct {
	Fixed bool
}

func (options *RowOptions) IsFixed() bool {
	return options.Fixed
}

func (options *RowOptions) IsValid() bool {
	return options != nil
}

func boundColumnItem(previous image.Point, bounds image.Rectangle, height int) image.Rectangle {
	return image.Rect(
		previous.X,
		previous.Y,
		bounds.Max.X,
		previous.Y+height,
	)
}

func measureColumnItem(bounds image.Rectangle) int {
	return bounds.Dy()
}

func translateColumnItem(height int) image.Point {
	return image.Pt(0, height)
}

func NewColumn() *Column {
	column := new(Column)
	return column
}

func (column *Column) AppendRow(row Widget, options *RowOptions) {
	column.addItem(row, options, measureColumnItem, translateColumnItem)

	maxX := maxInt(column.Bounds().Max.X, row.Bounds().Max.X)
	for _, row := range column.items {
		bounds := row.Bounds()
		bounds.Max.X = maxX
		row.SetBounds(bounds)
	}

	column.updateBounds()
}

func (column *Column) SetBounds(bounds image.Rectangle) {
	column.bounds = bounds
	column.scale(boundColumnItem, measureColumnItem, translateColumnItem)
}
