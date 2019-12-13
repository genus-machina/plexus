package amygdala

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func saveImage(t *testing.T, name string, image image.Image) {
	file, err := os.Create(t.Name() + "_" + name + ".png")
	if err != nil {
		t.Errorf("Failed to create image '%s': %s", name, err.Error())
	}
	defer file.Close()

	if err := png.Encode(file, image); err != nil {
		t.Errorf("Failed to save image '%s': %s", name, err.Error())
	}
}
