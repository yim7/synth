package ui

import (
	"log"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/yim7/synth/ui/events"
)

type Application struct {
	window *Window
	m      sync.Mutex
	quit   bool
}

var (
	App           *Application
	createAppOnce sync.Once
)

func init() {
	if err := sdl.Init(sdl.INIT_AUDIO | sdl.INIT_VIDEO | sdl.INIT_TIMER); err != nil {
		log.Fatal(err)
	}
}

func createSharedApp() {
	App = &Application{}
}

func NewApplication() *Application {
	createAppOnce.Do(createSharedApp)
	return App
}

func (app *Application) Run() {
loop:
	for !app.quit {
		window := app.window
		window.Renderer()

		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e := e.(type) {
			case *sdl.QuitEvent:
				app.quit = true
				break loop
			case *sdl.KeyboardEvent:
				event := events.SDLEventToKeyboardEvent(e)
				if event.Type == events.KeyDown {
					log.Println("key down:", event)
					window.KeyDown(event)
				} else {
					log.Println("key up:", event)
					window.KeyUp(event)
				}
			case *sdl.MouseButtonEvent, *sdl.MouseMotionEvent:
				event := events.SDLEventToMouseEvent(e)
				// log.Println("mouse event:", event)
				switch event.Type {
				case events.MouseDown:
					window.MouseDown(event)
				case events.MouseUp:
					window.MouseUp(event)
				case events.MouseMove:
					window.MouseMove(event)
				}
			default:
				// log.Println("todo handle event:", e)
			}
		}
	}
}

func (app *Application) SetWindow(w *Window) {
	app.m.Lock()
	defer app.m.Unlock()
	app.window = w
}

// func (app *Application) RunWithWindow(w *Window) {
// 	app.SetWindow(w)
// 	app.Run()
// }

func (app *Application) Destory() {
	app.window.Destory()
}
