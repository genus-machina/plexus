package amygdala

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
)

type TextBox struct {
	bounds     image.Rectangle
	face       font.Face
	text       string
	textBounds image.Rectangle
}

func NewTextBox(face font.Face, text string) *TextBox {
	bounds, _ := font.BoundString(face, text)
	textBounds := rectangleFromFixed(bounds)

	widget := new(TextBox)
	widget.bounds = textBounds
	widget.face = face
	widget.text = text
	widget.textBounds = textBounds
	return widget
}

func (widget *TextBox) Bounds() image.Rectangle {
	return widget.bounds
}

func (widget *TextBox) Render(canvas draw.Image) {
	buffer := image.NewNRGBA(widget.bounds)

	drawer := new(font.Drawer)
	drawer.Dot = pointToFixed(widget.bounds.Min.Sub(widget.textBounds.Min))
	drawer.Dst = buffer
	drawer.Face = widget.face
	drawer.Src = image.NewUniform(color.White)
	drawer.DrawString(widget.text)

	draw.Draw(canvas, widget.bounds, drawer.Dst, widget.bounds.Min, draw.Src)
}

func (widget *TextBox) SetBounds(bounds image.Rectangle) {
	widget.bounds = bounds
}
