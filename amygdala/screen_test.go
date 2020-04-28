package amygdala

import (
	"image"
	"testing"

	"periph.io/x/periph/conn/display/displaytest"
)

func createTestScreen() (*displaytest.Drawer, *Screen) {
	display := new(displaytest.Drawer)
	display.Img = image.NewNRGBA(image.Rect(0, 0, 128, 64))
	screen := NewScreen(display)
	return display, screen
}

func TestScreenBounds(t *testing.T) {
	display, screen := createTestScreen()
	assertRectangle(t, display.Bounds(), screen.Bounds())
}

func TestScreenRenderTextBox(t *testing.T) {
	display, screen := createTestScreen()
	face := NewFontFace(8)
	text := NewTextBox(face, "test string")
	assertScreenRender(t, screen, text)
	assertRectangle(t, display.Bounds(), text.Bounds())
	saveImage(t, "screen", display.Img)
}

func TestScreenRenderCell(t *testing.T) {
	display, screen := createTestScreen()
	face := NewFontFace(8)
	text := NewTextBox(face, "hello")
	cell := NewCell(text)
	cell.Align(AlignMiddle)
	cell.Justify(JustifyCenter)
	assertScreenRender(t, screen, cell)
	assertRectangle(t, display.Bounds(), cell.Bounds())
	saveImage(t, "screen", display.Img)
}

func TestScreenRenderColumn(t *testing.T) {
	display, screen := createTestScreen()
	face := NewFontFace(8)
	column := NewColumn()

	text := NewTextBox(face, "one")
	cell := NewCell(text)
	cell.Align(AlignTop)
	cell.Justify(JustifyLeft)
	column.AppendRow(cell, nil)

	text = NewTextBox(face, "two")
	cell = NewCell(text)
	cell.Align(AlignMiddle)
	cell.Justify(JustifyCenter)
	column.AppendRow(cell, nil)

	text = NewTextBox(face, "three")
	cell = NewCell(text)
	cell.Align(AlignBottom)
	cell.Justify(JustifyRight)
	column.AppendRow(cell, nil)

	assertScreenRender(t, screen, column)
	assertRectangle(t, display.Bounds(), column.Bounds())
	saveImage(t, "screen", display.Img)
}

func TestScreenRenderColumnRotated(t *testing.T) {
	display, screen := createTestScreen()
	face := NewFontFace(8)
	column := NewColumn()

	text := NewTextBox(face, "one")
	cell := NewCell(text)
	cell.Align(AlignTop)
	cell.Justify(JustifyLeft)
	column.AppendRow(cell, nil)

	text = NewTextBox(face, "two")
	cell = NewCell(text)
	cell.Align(AlignMiddle)
	cell.Justify(JustifyCenter)
	column.AppendRow(cell, nil)

	text = NewTextBox(face, "three")
	cell = NewCell(text)
	cell.Align(AlignBottom)
	cell.Justify(JustifyRight)
	column.AppendRow(cell, nil)

	screen.Rotate(IMAGE_ROTATE_90)
	assertScreenRender(t, screen, column)
	assertRectangle(t, image.Rect(0, 0, 64, 128), column.Bounds())
	saveImage(t, "screen", display.Img)
}

func TestScreenLayout(t *testing.T) {
	display, screen := createTestScreen()
	column := NewColumn()
	row := NewRow()

	face := NewFontFace(7)
	text := NewTextBox(face, "left")
	cell := NewCell(text)
	cell.Align(AlignTop)
	cell.Justify(JustifyLeft)
	row.AppendColumn(cell, &ColumnOptions{Fixed: true})

	png, _ := NewPNG("dialog-warning-symbolic.png")
	cell = NewCell(png)
	cell.Align(AlignTop)
	cell.Justify(JustifyLeft)
	cell.Pad(PadAll, 2)
	row.AppendColumn(cell, nil)

	text = NewTextBox(face, "right")
	cell = NewCell(text)
	cell.Align(AlignTop)
	cell.Justify(JustifyRight)
	row.AppendColumn(cell, nil)

	row.SetBounds(image.Rect(0, 0, display.Bounds().Dx(), 12))
	column.AppendRow(row, &RowOptions{Fixed: true})

	face = NewFontFace(20)
	text = NewTextBox(face, "body")
	cell = NewCell(text)
	cell.Align(AlignMiddle)
	cell.Justify(JustifyCenter)
	column.AppendRow(cell, nil)

	assertScreenRender(t, screen, column)
	assertRectangle(t, display.Bounds(), column.Bounds())
	saveImage(t, "screen", display.Img)
}
