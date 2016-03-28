package mapper

import (
	"github.com/manuelmenzella/TweetMap/tools"
	"image"
	"image/color"
	"image/draw"
	"math"
)

type Mapper struct {
	MapImage *image.RGBA
	Region   tools.Region
}

func NewMapper(size image.Point) *Mapper {
	return NewMapperRegion(size, tools.Region{-180, 180, -75, 75})
}

func NewMapperRegion(size image.Point, region tools.Region) *Mapper {
	rect := image.Rect(0, 0, size.X, size.Y)
	return &Mapper{image.NewRGBA(rect), region}
}

func (m *Mapper) Clear(c color.RGBA) {
	draw.Draw(
		m.MapImage,
		m.MapImage.Bounds(),
		&image.Uniform{c},
		image.ZP,
		draw.Src)
}

func (m *Mapper) DrawCircleCoor(lon, lat, r float64, c color.RGBA) {
	imageSize := m.MapImage.Bounds().Size()

	lonRange := m.Region.LonMax - m.Region.LonMin
	lonX := int(tools.Round(
		float64(imageSize.X) / lonRange * (lon - m.Region.LonMin)))

	latScaled := scaleLat(lat)
	latMinScaled := scaleLat(m.Region.LatMin)
	latMaxScaled := scaleLat(m.Region.LatMax)
	latRangeScaled := latMaxScaled - latMinScaled
	latY := int(tools.Round(
		float64(imageSize.Y) / latRangeScaled * (latScaled - latMinScaled)))

	m.drawCircle(lonX, imageSize.Y-latY, r, c)
}

func (m *Mapper) drawCircle(x, y int, r float64, c color.RGBA) {
	intRadius := int(math.Ceil(r))
	dstRect := image.Rect(x-intRadius, y-intRadius, x+intRadius, y+intRadius)
	srcPoint := image.Point{x - intRadius, y - intRadius}

	orR, orG, orB, orA := c.RGBA()
	newC := color.RGBA{uint8(orR >> 8), uint8(orG >> 8), uint8(orB >> 8), 255}
	alphaRatio := float64(orA) / math.Pow(2, 16)
	circleMask := circle{image.Point{x, y}, r, alphaRatio}

	draw.DrawMask(
		m.MapImage,
		dstRect,
		&image.Uniform{newC},
		image.ZP,
		&circleMask,
		srcPoint,
		draw.Over)
}

func scaleLat(lat float64) float64 {
	latRad := lat * math.Pi / 180
	return math.Log(math.Tan(latRad) + 1/math.Cos(latRad))
}
