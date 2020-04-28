package amygdala

import (
	"image"
	"image/color"
	"image/draw"
	"testing"
)

func TestRotate0(t *testing.T) {
	black := image.NewUniform(color.Black)
	bounds := image.Rect(0, 0, 128, 64)
	canvas := image.NewNRGBA(bounds)
	rotated := NewRotatedImage(canvas, IMAGE_ROTATE_0)
	assertRectangle(t, bounds, rotated.Bounds())
	draw.Draw(rotated, image.Rect(0, 0, 8, 8), black, image.Pt(0, 0), draw.Src)
	saveImage(t, "canvas", canvas)
	saveImage(t, "rotated", rotated)
}

func TestRotate90(t *testing.T) {
	black := image.NewUniform(color.Black)
	bounds := image.Rect(0, 0, 128, 64)
	canvas := image.NewNRGBA(bounds)
	rotated := NewRotatedImage(canvas, IMAGE_ROTATE_90)
	assertRectangle(t, image.Rect(0, 0, 64, 128), rotated.Bounds())
	draw.Draw(rotated, image.Rect(0, 0, 8, 8), black, image.Pt(0, 0), draw.Src)
	saveImage(t, "canvas", canvas)
	saveImage(t, "rotated", rotated)
}

func TestRotate180(t *testing.T) {
	black := image.NewUniform(color.Black)
	bounds := image.Rect(0, 0, 128, 64)
	canvas := image.NewNRGBA(bounds)
	rotated := NewRotatedImage(canvas, IMAGE_ROTATE_180)
	assertRectangle(t, image.Rect(0, 0, 128, 64), rotated.Bounds())
	draw.Draw(rotated, image.Rect(60, 0, 68, 8), black, image.Pt(0, 0), draw.Src)
	saveImage(t, "canvas", canvas)
	saveImage(t, "rotated", rotated)
}

func TestRotate270(t *testing.T) {
	black := image.NewUniform(color.Black)
	bounds := image.Rect(0, 0, 128, 64)
	canvas := image.NewNRGBA(bounds)
	rotated := NewRotatedImage(canvas, IMAGE_ROTATE_270)
	assertRectangle(t, image.Rect(0, 0, 64, 128), rotated.Bounds())
	draw.Draw(rotated, image.Rect(0, 0, 8, 8), black, image.Pt(0, 0), draw.Src)
	saveImage(t, "canvas", canvas)
	saveImage(t, "rotated", rotated)
}
