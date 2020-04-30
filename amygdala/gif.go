package amygdala

import (
	"image"
	"image/draw"
	"image/gif"
	"os"
	"sync"
	"time"
)

type GIF struct {
	bounds  image.Rectangle
	content *gif.GIF

	index   int
	mutex   sync.Mutex
	playing bool
}

func NewGIF(path string) (*GIF, error) {
	widget := new(GIF)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if widget.content, err = gif.DecodeAll(file); err != nil {
		return nil, err
	}

	widget.bounds = widget.content.Image[0].Bounds()
	return widget, nil

}

func (widget *GIF) advanceFrame() time.Duration {
	widget.mutex.Lock()
	defer widget.mutex.Unlock()

	if widget.playing {
		delay := widget.getDelay(widget.index)
		widget.index++
		return delay
	}

	return -1
}

func (widget *GIF) Bounds() image.Rectangle {
	return widget.bounds
}

func (widget *GIF) getDelay(index int) time.Duration {
	length := len(widget.content.Delay)
	delayIndex := index % length
	return time.Duration(widget.content.Delay[delayIndex]) * 10 * time.Millisecond
}

func (widget *GIF) getFrame(index int) image.Image {
	length := len(widget.content.Image)
	frameIndex := index % length
	return widget.content.Image[frameIndex]
}

func (widget *GIF) Halt() {
	widget.mutex.Lock()
	defer widget.mutex.Unlock()

	if widget.playing {
		widget.playing = false
	}
}

func (widget *GIF) play() {
	var delay time.Duration
	for ; delay >= 0; delay = widget.advanceFrame() {
		time.Sleep(delay)
	}
}

func (widget *GIF) Render(canvas draw.Image) {
	widget.mutex.Lock()
	defer widget.mutex.Unlock()
	frame := widget.getFrame(widget.index)
	scaleImage(canvas, widget.bounds, frame, frame.Bounds())

	if !widget.playing {
		widget.playing = true
		go widget.play()
	}
}

func (widget *GIF) SetBounds(bounds image.Rectangle) {
	widget.mutex.Lock()
	defer widget.mutex.Unlock()
	widget.bounds = computeImageBounds(widget.getFrame(0), bounds)
}
