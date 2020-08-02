package amygdala

import (
	"image"
	"testing"
)

func TestTextAreaFit(t *testing.T) {
	face := NewFontFace(10)
	text := "Hello"
	widget := NewTextArea(face, text)
	canvas := image.NewNRGBA(image.Rect(0, 0, 128, 64))

	widget.SetBounds(canvas.Bounds())
	widget.Render(canvas)
	saveImage(t, "wrap", canvas)
}

func TestTextAreaLong(t *testing.T) {
	face := NewFontFace(10)
	text := "This is a really long thing to print"
	widget := NewTextArea(face, text)
	canvas := image.NewNRGBA(image.Rect(0, 0, 128, 64))

	widget.SetBounds(canvas.Bounds())
	widget.Render(canvas)
	saveImage(t, "wrap", canvas)
}

func TestTextAreaOverflow(t *testing.T) {
	face := NewFontFace(10)
	text := "Supercalifragilisticexpialidocious boom"
	widget := NewTextArea(face, text)
	canvas := image.NewNRGBA(image.Rect(0, 0, 128, 64))

	widget.SetBounds(canvas.Bounds())
	widget.Render(canvas)
	saveImage(t, "wrap", canvas)
}

func TestTextAreaWrap(t *testing.T) {
	face := NewFontFace(10)
	text := "Tropical Cyclone Statement"
	widget := NewTextArea(face, text)
	canvas := image.NewNRGBA(image.Rect(0, 0, 128, 64))

	widget.SetBounds(canvas.Bounds())
	widget.Render(canvas)
	saveImage(t, "wrap", canvas)
}
