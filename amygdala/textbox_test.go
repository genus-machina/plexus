package amygdala

import (
	"image"
	"image/draw"
	"testing"

	"golang.org/x/image/font"
)

func createTestTextBox() (draw.Image, *TextBox) {
	content := "hello world"
	face := NewFontFace(12)
	widget := NewTextBox(face, content)
	canvas := image.NewNRGBA(widget.Bounds())
	return canvas, widget
}

func TestTextBoxBounds(t *testing.T) {
	content := "hello world"
	face := NewFontFace(12)
	widget := NewTextBox(face, content)
	expectedBounds, _ := font.BoundString(face, content)
	assertRectangle(t, rectangleFromFixed(expectedBounds), widget.Bounds())

	newBounds := image.Rect(0, 0, 5, 6)
	widget.SetBounds(newBounds)
	assertRectangle(t, newBounds, widget.Bounds())
}

func TestTextBoxRender(t *testing.T) {
	canvas, widget := createTestTextBox()
	widget.Render(canvas)
	saveImage(t, "content", canvas)
}

func TestTextBoxRenderOversized(t *testing.T) {
	canvas, widget := createTestTextBox()
	widget.SetBounds(image.Rect(0, -20, 20, 0))
	widget.Render(canvas)
	saveImage(t, "content", canvas)
}
