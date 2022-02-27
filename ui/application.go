package ui

import (
	"log"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

type Application struct {
	window *Window
	sync.Mutex
	quit bool
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
		app.window.Renderer()

		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.QuitEvent:
				app.quit = true
				break loop
			default:
				app.window.HandleEvent(e)
				// log.Println("todo handle event:", e)
			}
		}
	}
}

func (app *Application) SetWindow(w *Window) {
	app.Lock()
	defer app.Unlock()
	app.window = w
}

func (app *Application) Destory() {
	app.window.Destory()
}
