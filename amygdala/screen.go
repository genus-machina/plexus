package amygdala

import (
	"image"
	"image/color"

	"periph.io/x/periph/conn/display"
)

type Screen struct {
	display display.Drawer
	widget  Widget
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

func (screen *Screen) SetContent(widget Widget) {
	widget.SetBounds(screen.display.Bounds())
	screen.widget = widget
}

func (screen *Screen) Render() error {
	buffer := image.NewNRGBA(screen.display.Bounds())
	screen.widget.Render(buffer)
	return screen.display.Draw(screen.display.Bounds(), buffer, image.Pt(0, 0))
}
