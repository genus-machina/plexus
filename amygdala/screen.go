package amygdala

import (
	"image"
	"image/color"

	"periph.io/x/periph/conn/display"
)

type Screen struct {
	display  display.Drawer
	rotation int
}

func NewScreen(display display.Drawer) *Screen {
	screen := new(Screen)
	screen.display = display
	return screen
}

func (screen *Screen) Bounds() image.Rectangle {
	return screen.display.Bounds()
}

func (screen *Screen) Clear() error {
	buffer := image.NewUniform(color.Black)
	return screen.display.Draw(screen.display.Bounds(), buffer, image.Pt(0, 0))
}

func (screen *Screen) Render(content Widget) error {
	buffer := image.NewNRGBA(screen.display.Bounds())
	canvas := NewRotatedImage(buffer, screen.rotation)
	content.SetBounds(canvas.Bounds())
	content.Render(canvas)
	return screen.display.Draw(screen.display.Bounds(), buffer, image.Pt(0, 0))
}

func (screen *Screen) Rotate(rotation int) {
	screen.rotation = rotation
}
