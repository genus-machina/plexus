package amygdala

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

type PNG struct {
	bounds  image.Rectangle
	content image.Image
}

func NewPNG(path string) (*PNG, error) {
	widget := new(PNG)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if widget.content, err = png.Decode(file); err != nil {
		return nil, err
	}

	widget.bounds = widget.content.Bounds()
	return widget, nil
}

func (widget *PNG) Bounds() image.Rectangle {
	return widget.bounds
}

func (widget *PNG) Render(canvas draw.Image) {
	scaleImage(canvas, widget.bounds, widget.content, widget.content.Bounds())
}

func (widget *PNG) SetBounds(bounds image.Rectangle) {
	widget.bounds = computeImageBounds(widget.content, bounds)
}
