package ui

import (
	"image/color"
)

var (
	TransparentColor = RGBA(0, 0, 0, 0)
	BlackColor       = RGBA(0, 0, 0, 255)
	NavyColor        = RGBA(0, 0, 128, 255)
	BlueColor        = RGBA(0, 0, 255, 255)
	GreenColor       = RGBA(0, 128, 0, 255)
	TealColor        = RGBA(0, 128, 128, 255)
	LimeColor        = RGBA(0, 255, 0, 255)
	AquaColor        = RGBA(0, 255, 255, 255)
	MaroonColor      = RGBA(128, 0, 0, 255)
	PurpleColor      = RGBA(128, 0, 128, 255)
	OliveColor       = RGBA(128, 128, 0, 255)
	GrayColor        = RGBA(128, 128, 128, 255)
	SilverColor      = RGBA(192, 192, 192, 255)
	RedColor         = RGBA(255, 0, 0, 255)
	FuchsiaColor     = RGBA(255, 0, 255, 255)
	YellowColor      = RGBA(255, 255, 0, 255)
	WhiteColor       = RGBA(255, 255, 255, 255)
)

func RGBA(r, g, b, a byte) color.RGBA {
	return color.RGBA{r, g, b, a}
}
