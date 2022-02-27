package ui

import (
	"image"
	"image/color"
	"log"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	frame image.Rectangle
	views []View
	w     *sdl.Window
	r     *sdl.Renderer
	sync.Mutex
}

func NewWindow(title string, frame image.Rectangle) *Window {
	w, err := sdl.CreateWindow(title, int32(frame.Min.X), int32(frame.Min.Y), int32(frame.Dx()), int32(frame.Dy()), sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}
	r, err := sdl.CreateRenderer(w, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatal(err)
	}

	if err := r.SetDrawBlendMode(sdl.BLENDMODE_BLEND); err != nil {
		log.Fatal(err)
	}

	return &Window{
		frame: frame,
		w:     w,
		r:     r,
	}
}

func (w *Window) Destory() {
	w.r.Destroy()
	w.w.Destroy()
}

func (w *Window) AddView(views ...View) {
	w.Lock()
	defer w.Unlock()
	w.views = append(w.views, views...)
}

func (w *Window) Clear() {
	w.r.SetDrawColor(0, 0, 0, 0)
	w.r.Clear()
}

func (w *Window) Renderer() {
	w.Clear()

	offset := image.Point{0, 0}
	for _, view := range w.views {
		view.Draw(offset)
	}

	w.r.Present()
}

func (w *Window) DrawRect(r image.Rectangle, c color.RGBA) {
	rect := sdl.Rect{
		X: int32(r.Min.X),
		Y: int32(r.Min.Y),
		W: int32(r.Dx()),
		H: int32(r.Dy()),
	}

	w.r.SetDrawColor(c.R, c.G, c.B, c.A)
	w.r.FillRect(&rect)
}

func (w *Window) DrawText(text string, font *Font, r image.Rectangle, c color.RGBA) {
	surface, err := font.RenderUTF8Blended(text, sdl.Color(c))
	if err != nil {
		log.Fatal(err)
	}
	defer surface.Free()
	texture, err := w.r.CreateTextureFromSurface(surface)
	if err != nil {
		log.Fatal(err)
	}
	defer texture.Destroy()

	rect := sdl.Rect{
		X: int32(r.Min.X+r.Dx()/2) - surface.W/2,
		Y: int32(r.Min.Y+r.Dy()/2) - surface.H/2,
		W: surface.W,
		H: surface.H,
	}

	w.r.Copy(texture, nil, &rect)
}
