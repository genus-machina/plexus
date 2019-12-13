package amygdala

import (
	"image"
	"image/draw"
)

type Widget interface {
	Bounds() image.Rectangle
	Render(canvas draw.Image)
	SetBounds(bounds image.Rectangle)
}
