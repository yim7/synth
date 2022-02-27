package ui

import (
	"image"
	"image/color"
)

type Button struct {
	frame      image.Rectangle
	background color.RGBA
	*TextLabel
	*Control
}

func NewButton(text string, frame image.Rectangle) *Button {
	margin := 5
	labelFrame := image.Rect(margin, margin, frame.Dx()-margin, frame.Dy()-margin)
	label := NewTextLabel(text, labelFrame)
	label.SetColor(BlueColor)

	return &Button{
		frame:      frame,
		background: SilverColor,
		TextLabel:  label,
		Control:    NewControl(),
	}
}

func (view *Button) SetBackground(c color.RGBA) {
	view.background = c
}

func (view *Button) Frame() image.Rectangle {
	return view.frame
}

func (view *Button) Draw(offset image.Point) {
	w := App.window
	rect := view.frame.Add(offset)
	w.DrawRect(rect, view.background)
	view.TextLabel.Draw(rect.Min)
	// log.Println("button draw frame:", rect, view.background)
}

func (view *Button) Update() {

}
