package mapper

import (
	"image"
	"image/color"
	"math"
)

type circle struct {
	Center     image.Point
	Radius     float64
	AlphaRatio float64
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	intRadius := int(math.Ceil(c.Radius))
	return image.Rect(
		c.Center.X-intRadius,
		c.Center.Y-intRadius,
		c.Center.X+intRadius,
		c.Center.Y+intRadius)
}

func (c *circle) At(x, y int) color.Color {
	dX, dY := float64(x-c.Center.X), float64(y-c.Center.Y)
	dZ := math.Sqrt(math.Pow(dX, 2) + math.Pow(dY, 2))
	if dZ > c.Radius {
		return color.Alpha{0}
	} else {
		ratio := (1 - dZ/c.Radius) * c.AlphaRatio
		return color.Alpha{uint8(ratio * (1 << 8))}
	}
}
