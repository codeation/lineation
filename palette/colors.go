package palette

import (
	"image/color"
)

type ColorType color.Color

var (
	black  = color.RGBA{0, 0, 0, 255}
	white  = color.RGBA{255, 255, 255, 255}
	gray   = color.RGBA{208, 208, 208, 255}
	silver = color.RGBA{224, 224, 224, 255}
	lite   = color.RGBA{232, 232, 232, 255}
	red    = color.RGBA{255, 0, 0, 255}
	cursor = color.RGBA{255, 0, 0, 127}
)

var (
	DefaultBackground    = ColorType(white)
	DefaultBoxBackground = ColorType(white)
	DefaultEdge          = ColorType(gray)
	ActiveBoxBackground  = ColorType(silver)
	ActiveEdge           = ColorType(red)
	DefaultText          = ColorType(black)
	CursorBlock          = ColorType(cursor)
	DefaultLine          = ColorType(lite)
)

func (p *Palette) Color(c ColorType) color.Color {
	return color.Color(c)
}
