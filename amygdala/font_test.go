package amygdala

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

func NewFontFace(size float64) font.Face {
	ttf, _ := truetype.Parse(goregular.TTF)
	options := new(truetype.Options)
	options.DPI = 133
	options.Hinting = font.HintingFull
	options.Size = size
	return truetype.NewFace(ttf, options)
}
