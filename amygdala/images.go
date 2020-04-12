package amygdala

import (
	"image"
	"image/color"
)

type AlphaXor struct {
	left, right image.Image
}

func NewAlphaXor(left, right image.Image) *AlphaXor {
	xor := new(AlphaXor)
	xor.left = left
	xor.right = right
	return xor
}

func (xor *AlphaXor) At(x, y int) color.Color {
	lr, lg, lb, la := color.NRGBAModel.Convert(xor.left.At(x, y)).RGBA()
	rr, rg, rb, ra := color.NRGBAModel.Convert(xor.right.At(x, y)).RGBA()
	result := new(color.NRGBA)
	result.R = uint8(lr | rr)
	result.G = uint8(lg | rg)
	result.B = uint8(lb | rb)
	result.A = uint8(la ^ ra)
	return result
}

func (xor *AlphaXor) Bounds() image.Rectangle {
	return xor.left.Bounds().Intersect(xor.right.Bounds())
}

func (xor *AlphaXor) ColorModel() color.Model {
	return color.NRGBAModel
}
