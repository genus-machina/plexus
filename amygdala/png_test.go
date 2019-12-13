package amygdala

import (
	"image"
	"image/draw"
	"testing"
)

func createTestPNG() (draw.Image, *PNG) {
	canvas := image.NewNRGBA(image.Rect(0, 0, 128, 64))
	png, _ := NewPNG("dialog-warning-symbolic.png")
	return canvas, png
}

func TestPNGRender(t *testing.T) {
	canvas, png := createTestPNG()
	png.SetBounds(canvas.Bounds())
	png.Render(canvas)
	saveImage(t, "png", canvas)
}

func TestPNGSetBounds(t *testing.T) {
	canvas, png := createTestPNG()
	bounds := image.Rect(0, 0, 12, 12)
	png.SetBounds(bounds)
	assertRectangle(t, bounds, png.Bounds())
	png.Render(canvas)
	saveImage(t, "png", canvas)
}
