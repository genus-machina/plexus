package amygdala

import (
	"image"
	"image/draw"
)

type Overlay struct {
	left, right Widget
}

func NewOverlay(left, right Widget) *Overlay {
	widget := new(Overlay)
	widget.left = left
	widget.right = right
	return widget
}

func (widget *Overlay) Bounds() image.Rectangle {
	return widget.left.Bounds().Union(widget.right.Bounds())
}

func (widget *Overlay) Render(canvas draw.Image) {
	left := image.NewNRGBA(widget.left.Bounds())
	right := image.NewRGBA(widget.right.Bounds())
	widget.left.Render(left)
	widget.right.Render(right)
	merged := NewAlphaXor(left, right)
	draw.Draw(canvas, widget.Bounds(), merged, merged.Bounds().Min, draw.Src)
}

func (widget *Overlay) SetBounds(bounds image.Rectangle) {
	widget.left.SetBounds(bounds)
	widget.right.SetBounds(bounds)
}
