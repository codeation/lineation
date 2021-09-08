package palette

import (
	"log"
	"sync"

	"github.com/codeation/impress"
)

type Palette struct {
}

type ColorType impress.Color

var (
	black  = impress.NewColor(0, 0, 0)
	white  = impress.NewColor(255, 255, 255)
	gray   = impress.NewColor(192, 192, 192)
	silver = impress.NewColor(239, 239, 239)
	lite   = impress.NewColor(224, 224, 224)
	red    = impress.NewColor(255, 0, 0)
)

var (
	DefaultBackground    = ColorType(white)
	DefaultBoxBackground = ColorType(white)
	DefaultEdge          = ColorType(gray)
	ActiveBoxBackground  = ColorType(silver)
	ActiveEdge           = ColorType(red)
	DefaultText          = ColorType(black)
	CursorBlock          = ColorType(red)
	DefaultLine          = ColorType(lite)
)

func NewPalette() *Palette {
	return &Palette{}
}

func (p *Palette) DefaultAppRect() impress.Rect {
	return impress.NewRect(0, 0, 1280, 768)
}

func (p *Palette) Color(c ColorType) impress.Color {
	return impress.Color(c)
}

func (p *Palette) HorizontalBoxAlign() int {
	return 20
}

func (p *Palette) VerticalBoxAlign() int {
	return 16
}

func (p *Palette) VerticalBoxOffset() int {
	return 16
}

func (p *Palette) HorizontalBoxOffset(level int) int {
	switch level {
	case 1:
		return 0
	case 2:
		return 50 + (p.BoxWidth(level)+p.BoxWidth(level-1))/2
	case 3:
		return 50 + (p.BoxWidth(level)+p.BoxWidth(level-1))/2
	default:
		return 50
	}
}

func (p *Palette) BoxWidth(level int) int {
	switch level {
	case 1:
		return 100
	case 2:
		return 140
	default:
		return 300
	}
}

func (p *Palette) BoxHeight(level int, linecount int) int {
	offset := 0
	if linecount == 0 {
		linecount = 1
	}
	if linecount > 1 {
		offset = (linecount - 1) * p.TextLineOffset()
	}
	return p.VerticalTextAlign()*2 + p.DefaultFont().Height*linecount + offset
}

func (p *Palette) VerticalTextAlign() int {
	return 6
}

func (p *Palette) TextLineOffset() int {
	return 2
}

func (p *Palette) HorizontalTextAlign() int {
	return 10
}

func (p *Palette) TextLinePoint(level int, lineno int) impress.Point {
	offset := 0
	if lineno < 0 {
		lineno = 0
	}
	if lineno > 0 {
		offset = (lineno) * p.TextLineOffset()
	}
	return impress.NewPoint(p.HorizontalTextAlign(), p.VerticalTextAlign()+p.DefaultFont().Height*lineno+offset)
}

func (p *Palette) CursorPoint() impress.Point {
	return impress.NewPoint(0, p.TextLineOffset())
}

func (p *Palette) CursorSize() impress.Size {
	return impress.NewSize(2, p.DefaultFont().Height)
}

var once sync.Once
var font *impress.Font

func (p *Palette) DefaultFont() *impress.Font {
	once.Do(func() {
		var err error
		font, err = impress.NewFont(`{"family":"Verdana"}`, 12)
		if err != nil {
			log.Fatalf("NewFont: %v", err)
		}
	})
	return font
}
