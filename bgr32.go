package framebuffer

import (
	"image"
	"image/color"
)

// BGRModel is the color model for BGR32 images.
var BGRModel = color.ModelFunc(bgrModel)

func bgrModel(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return color.RGBA{uint8(b >> 8), uint8(g >> 8), uint8(r >> 8), uint8(a >> 8)}
}

// BGR32 represents a 32-bit BGR image.
// Also reffered as BGR0
type BGR32 struct {
	Pix    []byte
	Rect   image.Rectangle
	Stride int
}

// Bounds returns the image's bounds.
func (i *BGR32) Bounds() image.Rectangle { return i.Rect }

// ColorModel returns the image's color model.
func (i *BGR32) ColorModel() color.Model { return BGRModel }

// At returns the color of the pixel at (x, y).
func (i *BGR32) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(i.Rect)) {
		return color.RGBA{}
	}

	n := i.PixOffset(x, y)
	return color.RGBA{i.Pix[n+2], i.Pix[n+1], i.Pix[n], 0xFF}
}

// Set sets the pixel at (x, y) to a color.
func (i *BGR32) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(i.Rect)) {
		return
	}

	n := i.PixOffset(x, y)
	c1 := color.RGBAModel.Convert(c).(color.RGBA)
	i.Pix[n+2] = c1.R
	i.Pix[n+1] = c1.G
	i.Pix[n] = c1.B
}

// PixOffset returns the index of the first element of the pixel at (x, y).
func (i *BGR32) PixOffset(x, y int) int {
	return (y-i.Rect.Min.Y)*i.Stride + (x-i.Rect.Min.X)*4
}

// SetBGR sets the pixel at (x, y) to a BGR color.
func (i *BGR32) SetBGR(x, y int, c RGBColor) {
	if !(image.Point{x, y}.In(i.Rect)) {
		return
	}

	n := i.PixOffset(x, y)
	i.Pix[n+2] = c.B
	i.Pix[n+1] = c.G
	i.Pix[n] = c.R
}
