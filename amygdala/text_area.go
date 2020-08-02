package amygdala

import (
	"image"
	"image/draw"
	"strings"

	"golang.org/x/image/font"
)

type TextArea struct {
	face   font.Face
	text   string
	widget Widget
}

func NewTextArea(face font.Face, text string) *TextArea {
	widget := new(TextArea)
	widget.face = face
	widget.text = text
	widget.widget = NewTextBox(face, text)
	return widget
}

func (widget *TextArea) Bounds() image.Rectangle {
	return widget.widget.Bounds()
}

func (widget *TextArea) buildCell(text string) *Cell {
	textbox := NewTextBox(widget.face, text)
	cell := NewCell(textbox)
	cell.Align(AlignMiddle)
	cell.Justify(JustifyCenter)
	return cell
}

func (widget *TextArea) buildCells(lines []string) []*Cell {
	cells := make([]*Cell, 0)

	for _, line := range lines {
		cells = append(cells, widget.buildCell(line))
	}

	return cells
}

func (widget *TextArea) buildTable(lines []*Cell) *Column {
	column := NewColumn()

	for _, cell := range lines {
		column.AppendRow(cell, nil)
	}

	return column
}

func (widget *TextArea) findMaxIndex(lines []*Cell) int {
	maxIndex := 0
	maxWidth := 0

	for index := 0; index < len(lines); index++ {
		if width := lines[index].Bounds().Dx(); width > maxWidth {
			maxIndex = index
			maxWidth = width
		}
	}

	return maxIndex
}

func (widget *TextArea) joinTokens(tokens [][]string) []string {
	lines := make([]string, 0)

	for _, line := range tokens {
		lines = append(lines, strings.Join(line, " "))
	}

	return lines
}

func (widget *TextArea) Render(canvas draw.Image) {
	widget.widget.Render(canvas)
}

func (widget *TextArea) SetBounds(bounds image.Rectangle) {
	found := false
	lines := []string{widget.text}
	rows := widget.buildCells(lines)

	for !found {
		index := widget.findMaxIndex(rows)

		if rows[index].Bounds().Dx() > bounds.Dx() {
			var ok bool
			lines, ok = widget.wrapLine(lines, index)
			rows = widget.buildCells(lines)
			found = !ok
		} else {
			found = true
		}
	}

	widget.widget = widget.buildTable(rows)
	widget.widget.SetBounds(bounds)
}

func (widget *TextArea) splitLines(lines []string) [][]string {
	tokens := make([][]string, 0)

	for _, line := range lines {
		tokens = append(tokens, strings.Split(line, " "))
	}

	return tokens
}

func (widget *TextArea) wrapLine(lines []string, index int) ([]string, bool) {
	tokens := widget.splitLines(lines)

	if index < len(tokens) && len(tokens[index]) > 1 {
		target := tokens[index]
		word, line := target[len(target)-1], target[:len(target)-1]
		tokens[index] = line

		if index < len(tokens)-1 {
			tokens[index+1] = append(
				[]string{word},
				tokens[index+1]...,
			)
		} else {
			tokens = append(tokens, []string{word})
		}

		return widget.joinTokens(tokens), true
	}

	return lines, false
}
