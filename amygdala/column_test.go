package amygdala

import (
	"image"
	"testing"
)

func TestColumnSetBounds(t *testing.T) {
	column := NewColumn()
	assertRectangle(t, image.Rectangle{}, column.Bounds())
	newBounds := image.Rect(0, 0, 1, 2)
	column.SetBounds(newBounds)
	assertRectangle(t, newBounds, column.Bounds())
}

func TestAddRowBounds(t *testing.T) {
	column := NewColumn()
	one := NewTestWidget(2, 3)
	two := NewTestWidget(2, 2)
	three := NewTestWidget(3, 4)
	assertRectangle(t, image.Rect(0, 0, 0, 0), column.Bounds())
	column.AppendRow(one, nil)
	assertRectangle(t, image.Rect(0, 0, 2, 3), column.Bounds())
	column.AppendRow(two, nil)
	assertRectangle(t, image.Rect(0, 0, 2, 5), column.Bounds())
	column.AppendRow(three, nil)
	assertRectangle(t, image.Rect(0, 0, 3, 9), column.Bounds())
}

func TestColumnRowBounds(t *testing.T) {
	column := NewColumn()
	one := NewTestWidget(5, 5)
	two := NewTestWidget(6, 6)
	three := NewTestWidget(7, 7)
	column.AppendRow(one, nil)
	column.AppendRow(two, nil)
	column.AppendRow(three, nil)
	column.SetBounds(image.Rect(0, 0, 128, 64))
	assertRectangle(t, image.Rect(0, 0, 128, 22), one.Bounds())
	assertRectangle(t, image.Rect(0, 22, 128, 43), two.Bounds())
	assertRectangle(t, image.Rect(0, 43, 128, 64), three.Bounds())
}

func TestColumnRender(t *testing.T) {
	canvas := image.NewNRGBA(image.Rect(0, 0, 128, 64))
	column := NewColumn()
	_, text1 := createTestTextBox()
	_, text2 := createTestTextBox()
	_, text3 := createTestTextBox()
	column.AppendRow(text1, nil)
	column.AppendRow(text2, nil)
	column.AppendRow(text3, nil)
	column.SetBounds(canvas.Bounds())
	column.Render(canvas)
	saveImage(t, "column", canvas)
}

func TestColumnFixedRows(t *testing.T) {
	column := NewColumn()
	one := NewTestWidget(8, 8)
	two := NewTestWidget(5, 10)
	three := NewTestWidget(3, 12)
	column.AppendRow(one, &RowOptions{Fixed: true})
	column.AppendRow(two, nil)
	column.AppendRow(three, nil)
	column.SetBounds(image.Rect(0, 0, 30, 30))
	assertRectangle(t, image.Rect(0, 0, 30, 30), column.Bounds())
	assertRectangle(t, image.Rect(0, 0, 30, 8), one.Bounds())
	assertRectangle(t, image.Rect(0, 8, 30, 19), two.Bounds())
	assertRectangle(t, image.Rect(0, 19, 30, 30), three.Bounds())
}
