package amygdala

import (
	"image"
	"testing"
)

func assertRectangle(t *testing.T, expected image.Rectangle, actual image.Rectangle) {
	if !actual.Eq(expected) {
		t.Errorf("expected rectangle %s but got %s", expected, actual)
	}
}

func assertScreenRender(t *testing.T, screen *Screen) {
	if err := screen.Render(); err != nil {
		t.Errorf("failed to render: %s", err.Error())
	}
}
