package events

import (
	"image"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type MouseEventType int
type MouseButton int

const (
	MouseUp   MouseEventType = sdl.MOUSEBUTTONUP
	MouseDown MouseEventType = sdl.MOUSEBUTTONDOWN
	MouseMove MouseEventType = sdl.MOUSEMOTION

	LeftButton   MouseButton = sdl.BUTTON_LEFT
	RightButton  MouseButton = sdl.BUTTON_RIGHT
	MiddleButton MouseButton = sdl.BUTTON_MIDDLE
)

type MouseEvent struct {
	Button MouseButton
	Type   MouseEventType
	Clicks int
	Time   time.Time
	Point  image.Point
}

func SDLEventToMouseEvent(e sdl.Event) *MouseEvent {
	switch e := e.(type) {
	case *sdl.MouseButtonEvent:
		return &MouseEvent{
			Button: MouseButton(e.Button),
			Type:   MouseEventType(e.Type),
			Clicks: int(e.Clicks),
			Time:   time.UnixMilli(int64(e.Timestamp)),
			Point:  image.Pt(int(e.X), int(e.Y)),
		}
	case *sdl.MouseMotionEvent:
		return &MouseEvent{
			Type:  MouseMove,
			Time:  time.UnixMilli(int64(e.Timestamp)),
			Point: image.Pt(int(e.X), int(e.Y)),
		}
	default:
		return nil
	}
}

type MouseEventResponder interface {
	MouseDown(*MouseEvent)

	MouseUp(*MouseEvent)

	MouseMove(*MouseEvent)
}
