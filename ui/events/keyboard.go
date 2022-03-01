package events

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type KeyboardEventType int

const (
	KeyDown KeyboardEventType = sdl.KEYDOWN
	KeyUp   KeyboardEventType = sdl.KEYUP
)

type KeyboardEvent struct {
	Key    string
	Type   KeyboardEventType
	Repeat int
	Time   time.Time
}

func SDLEventToKeyboardEvent(e *sdl.KeyboardEvent) *KeyboardEvent {
	return &KeyboardEvent{
		Key:    sdl.GetKeyName(e.Keysym.Sym),
		Type:   KeyboardEventType(e.Type),
		Repeat: int(e.Repeat),
		Time:   time.UnixMilli(int64(e.Timestamp)),
	}
}

type KeyboardEventResponder interface {
	KeyDown(*KeyboardEvent)

	KeyUp(*KeyboardEvent)
}
