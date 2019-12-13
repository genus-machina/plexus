package amygdala

import (
	"image"
	"image/draw"
	"image/png"
	"os"

	xdraw "golang.org/x/image/draw"
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
	xdraw.CatmullRom.Scale(canvas, widget.bounds, widget.content, widget.content.Bounds(), draw.Src, nil)
}

func (widget *PNG) SetBounds(bounds image.Rectangle) {
	contentAspectRatio := float64(widget.content.Bounds().Dx()) / float64(widget.content.Bounds().Dy())
	widgetAspectRatio := float64(bounds.Dx()) / float64(bounds.Dy())

	var scale float64
	if contentAspectRatio > widgetAspectRatio {
		scale = float64(bounds.Dx()) / float64(widget.content.Bounds().Dx())
	} else {
		scale = float64(bounds.Dy()) / float64(widget.content.Bounds().Dy())
	}

	widget.bounds = image.Rect(
		bounds.Min.X,
		bounds.Min.Y,
		bounds.Min.X+int(float64(widget.content.Bounds().Dx())*scale),
		bounds.Min.Y+int(float64(widget.content.Bounds().Dy())*scale),
	)
}
