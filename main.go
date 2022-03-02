package main

import (
	"image"
	"log"

	"github.com/yim7/synth/ui"
	"github.com/yim7/synth/ui/events"
)

func main() {
	window := ui.NewWindow("Synthesizer", image.Rect(0, 0, 600, 400))

	osci := NewOscillator(440, 44100, Sine)
	adsr := NewEnvelope(0.01, 0.1, 0.5, 0.1, 44100)
	adsr.SetInput(osci)

	output, _ := OpenAudioPlayer(44100, 1)
	output.Resume()
	output.SetInput(adsr)
	go output.Run()

	isBlack := false
	h := 200
	w := 30
	for i := 0; i < 30; i++ {
		c := ui.BlueColor
		bg := ui.WhiteColor
		if isBlack {
			bg = ui.BlackColor
		}

		button := ui.NewButton("", image.Rect(w*i, h, w*(i+1), 400))
		freq := float64(440 + 30*i)
		button.MouseDownAction = func(e *events.MouseEvent) {
			osci.frequency = freq
			log.Println("key down")
			adsr.NoteOn()
		}
		button.MouseUpAction = func(e *events.MouseEvent) {
			log.Println("key up")
			adsr.NoteOff()
		}

		button.SetColor(c)
		button.SetBackground(bg)
		window.AddView(button)

		isBlack = !isBlack
	}

	title := ui.NewTextLabel("Perfect Piano!", image.Rect(0, 0, 600, 200))
	title.SetColor(ui.GrayColor)
	title.SetFontSize(30)
	window.AddView(title)

	// start ui
	app := ui.NewApplication()
	defer app.Destory()

	app.SetWindow(window)
	app.Run()
}
