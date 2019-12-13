package amygdala

import (
	"image"
	"image/draw"
	"testing"
)

func createTestCell() (draw.Image, Widget, *Cell) {
	canvas := image.NewNRGBA(image.Rect(0, 0, 128, 64))
	content := NewTestWidget(8, 8)
	cell := NewCell(content)
	return canvas, content, cell
}

func TestCellSetBounds(t *testing.T) {
	_, content, cell := createTestCell()
	assertRectangle(t, content.Bounds(), cell.Bounds())

	newBounds := image.Rect(0, 0, 5, 6)
	cell.SetBounds(newBounds)
	assertRectangle(t, newBounds, cell.Bounds())
}

func TestCellRenderDefault(t *testing.T) {
	canvas, content, cell := createTestCell()
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(0, 56, 8, 64), content.Bounds())
	cell.Render(canvas)
	saveImage(t, "cell", canvas)
}

func TestCellRenderAlignBottom(t *testing.T) {
	canvas, content, cell := createTestCell()
	cell.Align(AlignBottom)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(0, 56, 8, 64), content.Bounds())
	cell.Render(canvas)
	saveImage(t, "cell", canvas)
}

func TestCellRenderAlignMiddle(t *testing.T) {
	canvas, content, cell := createTestCell()
	cell.Align(AlignMiddle)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(0, 28, 8, 36), content.Bounds())
	cell.Render(canvas)
	saveImage(t, "cell", canvas)
}

func TestCellRenderAlignTop(t *testing.T) {
	canvas, content, cell := createTestCell()
	cell.Align(AlignTop)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(0, 0, 8, 8), content.Bounds())
	cell.Render(canvas)
	saveImage(t, "cell", canvas)
}

func TestCellRenderJustifyCenter(t *testing.T) {
	canvas, content, cell := createTestCell()
	cell.Justify(JustifyCenter)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(60, 56, 68, 64), content.Bounds())
	cell.Render(canvas)
	saveImage(t, "cell", canvas)
}

func TestCellRenderJustifyLeft(t *testing.T) {
	canvas, content, cell := createTestCell()
	cell.Justify(JustifyLeft)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(0, 56, 8, 64), content.Bounds())
	cell.Render(canvas)
	saveImage(t, "cell", canvas)
}

func TestCellRenderJustifyRight(t *testing.T) {
	canvas, content, cell := createTestCell()
	cell.Justify(JustifyRight)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(120, 56, 128, 64), content.Bounds())
	cell.Render(canvas)
	saveImage(t, "cell", canvas)
}

func TestCellRenderCenter(t *testing.T) {
	canvas, content, cell := createTestCell()
	cell.Align(AlignMiddle)
	cell.Justify(JustifyCenter)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(60, 28, 68, 36), content.Bounds())
	cell.Render(canvas)
	saveImage(t, "cell", canvas)
}

func TestCellRenderTopRight(t *testing.T) {
	canvas, content, cell := createTestCell()
	cell.Align(AlignTop)
	cell.Justify(JustifyRight)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(120, 0, 128, 8), content.Bounds())
	cell.Render(canvas)
	saveImage(t, "cell", canvas)
}

func TestCellRenderOversized(t *testing.T) {
	canvas, _, cell := createTestCell()
	bounds := image.Rect(0, 0, 5, 5)
	cell.SetBounds(bounds)
	assertRectangle(t, bounds, cell.Bounds())
	cell.Render(canvas)
	saveImage(t, "cell", canvas)
}

func TestCellPadAll(t *testing.T) {
	canvas, content, cell := createTestCell()
	cell.Pad(PadAll, 2)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(2, 54, 10, 62), content.Bounds())
	cell.Render(canvas)

	cell.Align(AlignTop)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(2, 2, 10, 10), content.Bounds())
	cell.Render(canvas)

	cell.Justify(JustifyRight)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(118, 2, 126, 10), content.Bounds())
	cell.Render(canvas)

	cell.Align(AlignBottom)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(118, 54, 126, 62), content.Bounds())
	cell.Render(canvas)

	saveImage(t, "cell", canvas)
}

func TestCellPadBottom(t *testing.T) {
	canvas, content, cell := createTestCell()
	cell.Pad(PadBottom, 2)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(0, 54, 8, 62), content.Bounds())
	cell.Render(canvas)

	cell.Align(AlignTop)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(0, 0, 8, 8), content.Bounds())
	cell.Render(canvas)

	cell.Justify(JustifyRight)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(120, 0, 128, 8), content.Bounds())
	cell.Render(canvas)

	cell.Align(AlignBottom)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(120, 54, 128, 62), content.Bounds())
	cell.Render(canvas)

	saveImage(t, "cell", canvas)
}

func TestCellPadLeft(t *testing.T) {
	canvas, content, cell := createTestCell()
	cell.Pad(PadLeft, 2)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(2, 56, 10, 64), content.Bounds())
	cell.Render(canvas)

	cell.Align(AlignTop)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(2, 0, 10, 8), content.Bounds())
	cell.Render(canvas)

	cell.Justify(JustifyRight)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(120, 0, 128, 8), content.Bounds())
	cell.Render(canvas)

	cell.Align(AlignBottom)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(120, 56, 128, 64), content.Bounds())
	cell.Render(canvas)

	saveImage(t, "cell", canvas)
}

func TestCellPadRight(t *testing.T) {
	canvas, content, cell := createTestCell()
	cell.Pad(PadRight, 2)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(0, 56, 8, 64), content.Bounds())
	cell.Render(canvas)

	cell.Align(AlignTop)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(0, 0, 8, 8), content.Bounds())
	cell.Render(canvas)

	cell.Justify(JustifyRight)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(118, 0, 126, 8), content.Bounds())
	cell.Render(canvas)

	cell.Align(AlignBottom)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(118, 56, 126, 64), content.Bounds())
	cell.Render(canvas)

	saveImage(t, "cell", canvas)
}

func TestCellPadTop(t *testing.T) {
	canvas, content, cell := createTestCell()
	cell.Pad(PadTop, 2)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(0, 56, 8, 64), content.Bounds())
	cell.Render(canvas)

	cell.Align(AlignTop)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(0, 2, 8, 10), content.Bounds())
	cell.Render(canvas)

	cell.Justify(JustifyRight)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(120, 2, 128, 10), content.Bounds())
	cell.Render(canvas)

	cell.Align(AlignBottom)
	cell.SetBounds(canvas.Bounds())
	assertRectangle(t, image.Rect(120, 56, 128, 64), content.Bounds())
	cell.Render(canvas)

	saveImage(t, "cell", canvas)
}

func TestCellPadOversized(t *testing.T) {
	canvas, content, cell := createTestCell()
	bounds := image.Rect(0, 0, 5, 5)
	cell.Pad(PadAll, 2)
	cell.SetBounds(bounds)
	assertRectangle(t, bounds, cell.Bounds())
	assertRectangle(t, image.Rect(2, 2, 3, 3), content.Bounds())
	cell.Render(canvas)
	saveImage(t, "cell", canvas)
}

func TestCellPaddingOverflow(t *testing.T) {
	canvas, content, cell := createTestCell()
	bounds := image.Rect(0, 0, 5, 5)
	cell.Pad(PadAll, 10)
	cell.SetBounds(bounds)
	assertRectangle(t, bounds, cell.Bounds())
	assertRectangle(t, image.Rect(0, 0, 0, 0), content.Bounds())
	cell.Render(canvas)
	saveImage(t, "cell", canvas)
}

func TestCellResize(t *testing.T) {
	_, content, cell := createTestCell()
	bounds := image.Rect(0, 0, 10, 10)
	cell.SetBounds(bounds)
	assertRectangle(t, bounds, cell.Bounds())
	assertRectangle(t, image.Rect(0, 2, 8, 10), content.Bounds())

	bounds = image.Rect(0, 0, 5, 5)
	cell.SetBounds(bounds)
	assertRectangle(t, bounds, cell.Bounds())
	assertRectangle(t, image.Rect(0, 0, 5, 5), content.Bounds())

	bounds = image.Rect(0, 0, 20, 20)
	cell.SetBounds(bounds)
	assertRectangle(t, bounds, cell.Bounds())
	assertRectangle(t, image.Rect(0, 12, 8, 20), content.Bounds())
}
