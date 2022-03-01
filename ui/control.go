package ui

import (
	"log"

	"github.com/yim7/synth/ui/events"
)

type Control struct {
	KeyUpAction     func(*events.KeyboardEvent)
	KeyDownAction   func(*events.KeyboardEvent)
	MouseDownAction func(*events.MouseEvent)
	MouseUpAction   func(*events.MouseEvent)
	MouseMoveAction func(*events.MouseEvent)
}

func NewControl() *Control {
	return &Control{}
}

func (c *Control) KeyUp(e *events.KeyboardEvent) {
	log.Println("key up:", e.Key)
	action := c.KeyUpAction
	if action != nil {
		action(e)
	}
}

func (c *Control) KeyDown(e *events.KeyboardEvent) {
	log.Println("key down:", e.Key)
	action := c.KeyDownAction
	if action != nil {
		action(e)
	}
}

func (c *Control) MouseUp(e *events.MouseEvent) {
	log.Println("mouse up")
	action := c.MouseUpAction
	if action != nil {
		action(e)
	}
}

func (c *Control) MouseDown(e *events.MouseEvent) {
	log.Println("mouse down")
	action := c.MouseDownAction
	if action != nil {
		action(e)
	}
}

func (c *Control) MouseMove(e *events.MouseEvent) {
	log.Println("mouse move")
	action := c.MouseMoveAction
	if action != nil {
		action(e)
	}
}
