package ui

import "log"

type Control struct {
	KeyUpAction     func(key string)
	KeyDownAction   func(key string)
	MouseDownAction func()
	MouseUpAction   func()
}

func NewControl() *Control {
	return &Control{}
}

func (c *Control) KeyUp(key string) {
	log.Println("key up:", key)
	action := c.KeyUpAction
	if action != nil {
		action(key)
	}
}

func (c *Control) KeyDown(key string) {
	log.Println("key down:", key)
	action := c.KeyDownAction
	if action != nil {
		action(key)
	}
}

func (c *Control) MouseUp() {
	log.Println("mouse up")
	action := c.MouseUpAction
	if action != nil {
		action()
	}
}

func (c *Control) MouseDown() {
	log.Println("mouse down")
	action := c.MouseDownAction
	if action != nil {
		action()
	}
}
