package main

import (
	"image"
	"image/color"
	// "golang.org/x/tour/pic"
)

type Image struct {
	W, H  int
	color uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.W, i.H)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.color + uint8(x), i.color + uint8(y), 255, 255}
}

func main() {
	// m := Image{100, 100, 128}
	// pic.ShowImage(m)
}
