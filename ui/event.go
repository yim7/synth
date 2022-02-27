package ui

type KeyEventResponder interface {
	KeyDown(key string)
	KeyUp(key string)
}

type MouseEventResponder interface {
	MouseDown()
	MouseUp()
}
