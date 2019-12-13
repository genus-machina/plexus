package amygdala

import (
	"image"
	"image/draw"
)

const (
	AlignBottom Alignment = iota
	AlignMiddle
	AlignTop
)

const (
	JustifyLeft Justification = iota
	JustifyRight
	JustifyCenter
)

const (
	PadBottom = 1 << iota
	PadLeft
	PadRight
	PadTop

	PadAll = PadBottom + PadLeft + PadRight + PadTop
)

type Alignment int
type Justification int
type Padding int

type sides struct {
	bottom, left, right, top int
}

type Cell struct {
	alignment     Alignment
	bounds        image.Rectangle
	content       Widget
	contentBounds image.Rectangle
	justification Justification
	padding       sides
}

func NewCell(content Widget) *Cell {
	cell := new(Cell)
	cell.bounds = content.Bounds()
	cell.content = content
	cell.contentBounds = content.Bounds()
	return cell
}

func (widget *Cell) Align(alignment Alignment) {
	widget.alignment = alignment
}

func (widget *Cell) Bounds() image.Rectangle {
	return widget.bounds
}

func (widget *Cell) Justify(justification Justification) {
	widget.justification = justification
}

func (widget *Cell) Pad(sides int, padding int) {
	if sides&PadBottom != 0 {
		widget.padding.bottom = padding
	}
	if sides&PadLeft != 0 {
		widget.padding.left = padding
	}
	if sides&PadRight != 0 {
		widget.padding.right = padding
	}
	if sides&PadTop != 0 {
		widget.padding.top = padding
	}
}

func (widget *Cell) paddedBounds() image.Rectangle {
	padded := image.Rect(
		widget.bounds.Min.X+widget.padding.left,
		widget.bounds.Min.Y+widget.padding.top,
		widget.bounds.Max.X-widget.padding.right,
		widget.bounds.Max.Y-widget.padding.bottom,
	)

	if padded.In(widget.bounds) {
		return padded
	}

	return image.Rect(0, 0, 0, 0)
}

func (widget *Cell) Render(canvas draw.Image) {
	widget.content.Render(canvas)
}

func (widget *Cell) SetBounds(bounds image.Rectangle) {
	widget.bounds = bounds
	paddedBounds := widget.paddedBounds()

	if paddedBounds.Dx() < widget.contentBounds.Dx() || paddedBounds.Dy() < widget.contentBounds.Dy() {
		widget.content.SetBounds(paddedBounds)
		return
	}

	minX := paddedBounds.Min.X
	minY := paddedBounds.Max.Y - widget.contentBounds.Dy()
	maxX := paddedBounds.Min.X + widget.contentBounds.Dx()
	maxY := paddedBounds.Max.Y

	switch widget.alignment {
	case AlignBottom:
	case AlignMiddle:
		minY = paddedBounds.Min.Y + (paddedBounds.Dy()-widget.contentBounds.Dy())/2
		maxY = minY + widget.contentBounds.Dy()
	case AlignTop:
		minY = paddedBounds.Min.Y
		maxY = minY + widget.contentBounds.Dy()
	}

	switch widget.justification {
	case JustifyCenter:
		minX = paddedBounds.Min.X + (paddedBounds.Dx()-widget.contentBounds.Dx())/2
		maxX = minX + widget.contentBounds.Dx()
	case JustifyLeft:
	case JustifyRight:
		maxX = paddedBounds.Max.X
		minX = maxX - widget.contentBounds.Dx()
	}

	widget.content.SetBounds(image.Rect(minX, minY, maxX, maxY))
}
