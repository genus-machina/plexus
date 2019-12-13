package amygdala

import (
	"image"
)

type Row struct {
	vector
}

type ColumnOptions struct {
	Fixed bool
}

func (options *ColumnOptions) IsFixed() bool {
	return options.Fixed
}

func (options *ColumnOptions) IsValid() bool {
	return options != nil
}

func boundRowItem(previous image.Point, bounds image.Rectangle, width int) image.Rectangle {
	return image.Rect(
		previous.X,
		previous.Y,
		previous.X+width,
		bounds.Max.Y,
	)
}

func measureRowItem(bounds image.Rectangle) int {
	return bounds.Dx()
}

func translateRowItem(width int) image.Point {
	return image.Pt(width, 0)
}

func NewRow() *Row {
	row := new(Row)
	return row
}

func (row *Row) AppendColumn(column Widget, options *ColumnOptions) {
	row.addItem(column, options, measureRowItem, translateRowItem)

	maxY := maxInt(row.Bounds().Max.Y, column.Bounds().Max.Y)
	for _, column := range row.items {
		bounds := column.Bounds()
		bounds.Max.Y = maxY
		column.SetBounds(bounds)
	}

	row.updateBounds()
}

func (row *Row) SetBounds(bounds image.Rectangle) {
	row.bounds = bounds
	row.scale(boundRowItem, measureRowItem, translateRowItem)
}
