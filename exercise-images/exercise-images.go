package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, height int
	clr           uint8
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width-1, img.height-1)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{img.clr + uint8(x), img.clr + uint8(y), 234, 205}
}

func main() {
	m := Image{100, 100, 128}
	pic.ShowImage(m)
}
