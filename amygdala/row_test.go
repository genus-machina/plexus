package amygdala

import (
	"image"
	"testing"
)

func TestRowSetBounds(t *testing.T) {
	row := NewRow()
	assertRectangle(t, image.Rectangle{}, row.Bounds())
	newBounds := image.Rect(0, 0, 1, 2)
	row.SetBounds(newBounds)
	assertRectangle(t, newBounds, row.Bounds())
}

func TestAddColumnBounds(t *testing.T) {
	row := NewRow()
	one := NewTestWidget(2, 2)
	two := NewTestWidget(3, 2)
	three := NewTestWidget(4, 3)
	assertRectangle(t, image.Rect(0, 0, 0, 0), row.Bounds())
	row.AppendColumn(one, nil)
	assertRectangle(t, image.Rect(0, 0, 2, 2), row.Bounds())
	row.AppendColumn(two, nil)
	assertRectangle(t, image.Rect(0, 0, 5, 2), row.Bounds())
	row.AppendColumn(three, nil)
	assertRectangle(t, image.Rect(0, 0, 9, 3), row.Bounds())
}

func TestRowColumnBounds(t *testing.T) {
	row := NewRow()
	one := NewTestWidget(5, 5)
	two := NewTestWidget(6, 6)
	three := NewTestWidget(7, 7)
	row.AppendColumn(one, nil)
	row.AppendColumn(two, nil)
	row.AppendColumn(three, nil)
	row.SetBounds(image.Rect(0, 0, 128, 64))
	assertRectangle(t, image.Rect(0, 0, 43, 64), one.Bounds())
	assertRectangle(t, image.Rect(43, 0, 86, 64), two.Bounds())
	assertRectangle(t, image.Rect(86, 0, 128, 64), three.Bounds())
}

func TestRowRender(t *testing.T) {
	canvas := image.NewNRGBA(image.Rect(0, 0, 128, 64))
	row := NewRow()
	_, text1 := createTestTextBox()
	_, text2 := createTestTextBox()
	_, text3 := createTestTextBox()
	row.AppendColumn(text1, nil)
	row.AppendColumn(text2, nil)
	row.AppendColumn(text3, nil)
	row.SetBounds(canvas.Bounds())
	row.Render(canvas)
	saveImage(t, "row", canvas)
}

func TestRowCellRender(t *testing.T) {
	canvas := image.NewNRGBA(image.Rect(0, 0, 128, 64))
	row := NewRow()
	one := NewTestWidget(8, 8)
	two := NewTestWidget(8, 8)
	three := NewTestWidget(8, 8)

	cell := NewCell(one)
	cell.Align(AlignBottom)
	cell.Justify(JustifyLeft)
	row.AppendColumn(cell, nil)

	cell = NewCell(two)
	cell.Align(AlignMiddle)
	cell.Justify(JustifyCenter)
	row.AppendColumn(cell, nil)

	cell = NewCell(three)
	cell.Align(AlignTop)
	cell.Justify(JustifyRight)
	row.AppendColumn(cell, nil)

	row.SetBounds(canvas.Bounds())
	row.Render(canvas)
	saveImage(t, "row", canvas)
}

func TestRowColumnsRender(t *testing.T) {
	canvas := image.NewNRGBA(image.Rect(0, 0, 128, 64))
	row := NewRow()

	column := NewColumn()
	widget := NewTestWidget(8, 8)
	cell := NewCell(widget)
	cell.Align(AlignMiddle)
	cell.Justify(JustifyCenter)
	column.AppendRow(cell, nil)

	widget = NewTestWidget(8, 8)
	cell = NewCell(widget)
	cell.Align(AlignMiddle)
	cell.Justify(JustifyCenter)
	column.AppendRow(cell, nil)

	row.AppendColumn(column, nil)

	column = NewColumn()
	widget = NewTestWidget(8, 8)
	cell = NewCell(widget)
	cell.Align(AlignMiddle)
	cell.Justify(JustifyCenter)
	column.AppendRow(cell, nil)

	row.AppendColumn(column, nil)

	row.SetBounds(canvas.Bounds())
	row.Render(canvas)
	saveImage(t, "row", canvas)
}

func TestRowFixedColumns(t *testing.T) {
	row := NewRow()
	one := NewTestWidget(8, 8)
	two := NewTestWidget(5, 10)
	three := NewTestWidget(3, 12)
	row.AppendColumn(one, &ColumnOptions{Fixed: true})
	row.AppendColumn(two, nil)
	row.AppendColumn(three, nil)
	row.SetBounds(image.Rect(0, 0, 30, 30))
	assertRectangle(t, image.Rect(0, 0, 30, 30), row.Bounds())
	assertRectangle(t, image.Rect(0, 0, 8, 30), one.Bounds())
	assertRectangle(t, image.Rect(8, 0, 19, 30), two.Bounds())
	assertRectangle(t, image.Rect(19, 0, 30, 30), three.Bounds())
}

func TestRowFixedColumn(t *testing.T) {
	row := NewRow()
	one := NewTestWidget(8, 8)
	row.AppendColumn(one, &ColumnOptions{Fixed: true})
	row.SetBounds(image.Rect(0, 0, 30, 30))
	assertRectangle(t, image.Rect(0, 0, 30, 30), row.Bounds())
	assertRectangle(t, image.Rect(0, 0, 8, 30), one.Bounds())
}
