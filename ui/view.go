package ui

import "image"

type View interface {
	Frame() image.Rectangle

	Draw(offset image.Point)

	Update()
}
