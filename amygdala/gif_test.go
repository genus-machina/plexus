package amygdala

import (
	"image"
	"testing"
	"time"
)

func TestGIFRenderAnimated(t *testing.T) {
	canvas := image.NewNRGBA(image.Rect(0, 0, 128, 64))
	gif, _ := NewGIF("ballerine.gif")
	defer gif.Halt()

	gif.SetBounds(canvas.Bounds())
	gif.Render(canvas)
	saveImage(t, "one", canvas)

	time.Sleep(time.Second)
	gif.Render(canvas)
	saveImage(t, "two", canvas)
}

func TestGIFRenderStatic(t *testing.T) {
	canvas := image.NewNRGBA(image.Rect(0, 0, 128, 64))
	gif, _ := NewGIF("bunny.gif")
	defer gif.Halt()

	gif.SetBounds(canvas.Bounds())
	gif.Render(canvas)
	saveImage(t, "gif", canvas)
}

func TestGIFSetBounds(t *testing.T) {
	canvas := image.NewNRGBA(image.Rect(0, 0, 128, 64))
	gif, _ := NewGIF("ballerine.gif")
	defer gif.Halt()

	bounds := image.Rect(0, 0, 12, 12)
	gif.SetBounds(bounds)
	assertRectangle(t, image.Rect(0, 0, 9, 12), gif.Bounds())
	gif.Render(canvas)
	saveImage(t, "gif", canvas)
}
