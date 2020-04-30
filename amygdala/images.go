package amygdala

import (
	"image"
	"image/draw"

	xdraw "golang.org/x/image/draw"
)

func computeImageBounds(content image.Image, bounds image.Rectangle) image.Rectangle {
	contentAspectRatio := float64(content.Bounds().Dx()) / float64(content.Bounds().Dy())
	widgetAspectRatio := float64(bounds.Dx()) / float64(bounds.Dy())

	var scale float64
	if contentAspectRatio > widgetAspectRatio {
		scale = float64(bounds.Dx()) / float64(content.Bounds().Dx())
	} else {
		scale = float64(bounds.Dy()) / float64(content.Bounds().Dy())
	}

	return image.Rect(
		bounds.Min.X,
		bounds.Min.Y,
		bounds.Min.X+int(float64(content.Bounds().Dx())*scale),
		bounds.Min.Y+int(float64(content.Bounds().Dy())*scale),
	)
}

func scaleImage(dst draw.Image, dr image.Rectangle, src image.Image, sr image.Rectangle) {
	xdraw.CatmullRom.Scale(dst, dr, src, sr, draw.Src, nil)
}
