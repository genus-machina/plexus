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
	update  chan bool

	index    int
	maxIndex int
	mutex    sync.Mutex
	playing  bool
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

	switch widget.content.LoopCount {
	case 0:
		widget.maxIndex = -1
	case -1:
		widget.maxIndex = len(widget.content.Image)
	default:
		widget.maxIndex = len(widget.content.Image) * (widget.content.LoopCount + 1)
	}

	widget.update = make(chan bool, 1)
	return widget, nil
}

func (widget *GIF) advanceFrame() bool {
	var state bool
	widget.mutex.Lock()
	defer widget.mutex.Unlock()

	if widget.playing {
		if widget.maxIndex < 0 || widget.index < widget.maxIndex {
			widget.index++
			state = true
		}

		select {
		case widget.update <- state:
		default:
		}
	}

	return state
}

func (widget *GIF) Bounds() image.Rectangle {
	return widget.bounds
}

func (widget *GIF) getDelay() time.Duration {
	widget.mutex.Lock()
	defer widget.mutex.Unlock()

	length := len(widget.content.Delay)
	delayIndex := widget.index % length
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
	widget.playing = false
	close(widget.update)
}

func (widget *GIF) play() {
	for playing := true; playing; playing = widget.advanceFrame() {
		delay := widget.getDelay()
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

func (widget *GIF) Updates() <-chan bool {
	return widget.update
}
