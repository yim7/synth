package main

import (
	"image"
	"image/color"
	"log"

	"github.com/yim7/synth/ui"
	"github.com/yim7/synth/ui/events"
)

func main() {
	w := ui.NewWindow("Synth", image.Rect(0, 0, 800, 600))

	moving := false
	var prevPos image.Point

	button := ui.NewButton("Enter", image.Rect(400, 300, 500, 350))
	// keyevent action: change font size
	button.KeyDownAction = func(e *events.KeyboardEvent) {
		key := e.Key
		switch key {
		case "A":
			button.SetFontSize(button.FontSize() + 1)
		case "D":
			button.SetFontSize(button.FontSize() - 1)
		}
	}
	// mouse event action: drag button
	button.MouseDownAction = func(me *events.MouseEvent) {
		moving = true
		prevPos = me.Point
	}
	button.MouseUpAction = func(me *events.MouseEvent) {
		moving = false
	}
	button.MouseMoveAction = func(me *events.MouseEvent) {
		if moving {
			frame := button.Frame()
			offset := me.Point.Sub(prevPos)
			prevPos = me.Point
			newFrame := frame.Add(offset)
			button.SetFrame(newFrame)
		}
	}

	label := ui.NewTextLabel("Hello, World!", image.Rect(0, 100, 800, 250))
	label.SetFontSize(40)
	label.SetColor(color.RGBA{255, 0, 128, 255})

	// add subview to window
	w.AddView(button, label)

	app := ui.NewApplication()
	defer app.Destory()
	app.SetWindow(w)
	app.Run()
	log.Printf("app quit!")
}
