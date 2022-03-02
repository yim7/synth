package ui

import (
	"image"
	"image/color"
)

type TextLabel struct {
	frame      image.Rectangle
	text       string
	color      color.RGBA
	background color.RGBA
	font       *Font
}

func NewTextLabel(text string, frame image.Rectangle) *TextLabel {
	return &TextLabel{
		frame:      frame,
		text:       text,
		color:      BlackColor,
		background: TransparentColor,
		font:       DefaultFont,
	}
}

func (view *TextLabel) SetText(text string) {
	view.text = text
}

func (view *TextLabel) SetColor(c color.RGBA) {
	view.color = c
}

func (view *TextLabel) SetBackground(c color.RGBA) {
	view.background = c
}

func (view *TextLabel) SetFontSize(size int) {
	font, _ := LoadFont(view.font.name, size)
	view.font = font
}

func (view *TextLabel) FontSize() int {
	return view.font.size
}

func (view *TextLabel) Frame() image.Rectangle {
	return view.frame
}

func (view *TextLabel) Draw(offset image.Point) {
	if len(view.text) == 0 {
		return
	}

	w := App.window
	rect := view.frame.Add(offset)

	w.DrawRect(rect, view.background)
	w.DrawText(view.text, view.font, rect, view.color)
}

func (view *TextLabel) Update() {

}
