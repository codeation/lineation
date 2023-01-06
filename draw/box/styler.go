package box

import (
	"image"
	"image/color"

	"github.com/codeation/lineation/palette"
	"github.com/codeation/lineation/text"
)

type BoxTextOption struct {
	text.TextStyler
	b *Box
}

func newTextOption(b *Box) *BoxTextOption {
	return &BoxTextOption{
		TextStyler: text.NewSimpleTextOption(&text.TextOption{
			Font:       b.pal.DefaultFont(),
			LineHeight: b.pal.DefaultFont().Height + b.pal.TextLineOffset(),
			Margin:     image.Pt(b.pal.HorizontalTextAlign(), b.pal.VerticalTextAlign()),
			Foreground: b.pal.Color(palette.DefaultText),
		}),
		b: b,
	}
}

func (o *BoxTextOption) Edge() int {
	return o.b.pal.BoxWidth(o.b.level()) - o.b.pal.HorizontalTextAlign()*2
}

func (o *BoxTextOption) Size() image.Point {
	if o.b.textBox != nil {
		return image.Pt(o.b.width(), o.b.height())
	}
	return image.Pt(o.b.pal.BoxWidth(o.b.level()), o.b.pal.DefaultFont().Height)
}

func (o *BoxTextOption) Background() color.Color {
	if o.b.emphasized || o.b.isActive {
		return o.b.pal.Color(palette.ActiveBoxBackground)
	}
	return o.b.pal.Color(palette.DefaultBoxBackground)
}

func (o *BoxTextOption) Border() color.Color {
	if o.b.isActive {
		return o.b.pal.Color(palette.ActiveEdge)
	}
	return o.b.pal.Color(palette.DefaultEdge)
}

type BoxCursorOption struct {
	text.CursorStyler
	b *Box
}

func newCursorOption(b *Box) *BoxCursorOption {
	return &BoxCursorOption{
		CursorStyler: text.NewSimpleCursorOption(&text.CursorOption{
			Foreground: b.pal.Color(palette.CursorBlock),
			Size:       b.pal.CursorSize(),
		}),
		b: b,
	}
}

func (o *BoxCursorOption) Enable() bool {
	return o.b.isActive
}
