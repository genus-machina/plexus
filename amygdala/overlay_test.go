package amygdala

import (
	"image"
	"testing"
)

func TestOverlayBounds(t *testing.T) {
	left := NewTestWidget(10, 10)
	left.SetBounds(image.Rect(0, 0, 10, 10))
	right := NewTestWidget(15, 15)
	right.SetBounds(image.Rect(5, 5, 20, 20))
	overlay := NewOverlay(left, right)
	assertRectangle(t, image.Rect(0, 0, 20, 20), overlay.Bounds())

	overlay.SetBounds(image.Rect(5, 5, 10, 10))
	assertRectangle(t, image.Rect(5, 5, 10, 10), overlay.Bounds())
}

func TestOverlayRender(t *testing.T) {
	bounds := image.Rect(0, 0, 128, 64)
	canvas := image.NewRGBA(bounds)
	font := NewFontFace(20)
	hello := NewTextBox(font, "hello")
	world := NewTextBox(font, "WORLD")

	overlay := NewOverlay(hello, world)
	overlay.SetBounds(canvas.Bounds())
	overlay.Render(canvas)
	saveImage(t, "overlay", canvas)
}
