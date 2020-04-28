package amygdala

import (
	"image"
	"image/color"
	"image/draw"
)

const (
	IMAGE_ROTATE_0 = iota
	IMAGE_ROTATE_90
	IMAGE_ROTATE_180
	IMAGE_ROTATE_270
)

type RotatedImage struct {
	image    draw.Image
	rotation int
}

func NewRotatedImage(image draw.Image, rotation int) *RotatedImage {
	img := new(RotatedImage)
	img.image = image
	img.rotation = rotation
	return img
}

func (img *RotatedImage) At(x, y int) color.Color {
	tx, ty := img.translate(x, y)
	return img.image.At(tx, ty)
}

func (img *RotatedImage) Bounds() image.Rectangle {
	bounds := img.image.Bounds()

	if img.rotation == IMAGE_ROTATE_90 || img.rotation == IMAGE_ROTATE_270 {
		return image.Rect(
			bounds.Min.X,
			bounds.Min.Y,
			bounds.Min.X+bounds.Dy(),
			bounds.Min.Y+bounds.Dx(),
		)
	}

	return bounds
}

func (img *RotatedImage) ColorModel() color.Model {
	return img.image.ColorModel()
}

func (img *RotatedImage) Set(x, y int, c color.Color) {
	tx, ty := img.translate(x, y)
	img.image.Set(tx, ty, c)
}

func (img *RotatedImage) translate(x, y int) (int, int) {
	var tx, ty int
	bounds := img.image.Bounds()

	switch img.rotation {
	default:
		fallthrough
	case IMAGE_ROTATE_0:
		tx, ty = x, y
	case IMAGE_ROTATE_90:
		tx, ty = y, bounds.Max.Y-x-1
	case IMAGE_ROTATE_180:
		tx, ty = bounds.Max.X-x-1, bounds.Max.Y-y-1
	case IMAGE_ROTATE_270:
		tx, ty = bounds.Max.X-y-1, x
	}

	return tx, ty
}
