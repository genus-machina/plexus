package amygdala

import (
	"image"
	"image/color"
	"image/draw"
)

type TestWidget struct {
	bounds image.Rectangle
	image  *image.Uniform
}

func NewTestWidget(width, height int) Widget {
	widget := new(TestWidget)
	widget.bounds = image.Rect(0, 0, width, height)
	widget.image = image.NewUniform(color.Black)
	return widget
}

func (widget *TestWidget) Bounds() image.Rectangle {
	return widget.bounds
}

func (widget *TestWidget) Render(canvas draw.Image) {
	draw.Draw(canvas, widget.bounds, widget.image, image.Pt(0, 0), draw.Src)
}

func (widget *TestWidget) SetBounds(bounds image.Rectangle) {
	widget.bounds = bounds
}
