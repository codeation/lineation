package text

import (
	"image"
	"image/color"

	"github.com/codeation/impress"
)

type TextOption struct {
	Font       *impress.Font
	LineHeight int
	Edge       int
	Margin     image.Point
	Size       image.Point
	Foreground color.Color
	Background color.Color
	Border     color.Color
}

type CursorOption struct {
	Enable     bool
	Size       image.Point
	Foreground color.Color
}

type simpleTextOption struct {
	values *TextOption
}

type simpleCursorOption struct {
	values *CursorOption
}

func NewSimpleTextOption(values *TextOption) *simpleTextOption {
	return &simpleTextOption{values: values}
}

func (o *simpleTextOption) Font() *impress.Font     { return o.values.Font }
func (o *simpleTextOption) LineHeight() int         { return o.values.LineHeight }
func (o *simpleTextOption) Margin() image.Point     { return o.values.Margin }
func (o *simpleTextOption) Size() image.Point       { return o.values.Size }
func (o *simpleTextOption) Foreground() color.Color { return o.values.Foreground }
func (o *simpleTextOption) Background() color.Color { return o.values.Background }
func (o *simpleTextOption) Border() color.Color     { return o.values.Border }
func (o *simpleTextOption) Edge() int               { return o.values.Edge }

func NewSimpleCursorOption(values *CursorOption) *simpleCursorOption {
	return &simpleCursorOption{values: values}
}

func (o *simpleCursorOption) Enable() bool            { return o.values.Enable }
func (o *simpleCursorOption) Size() image.Point       { return o.values.Size }
func (o *simpleCursorOption) Foreground() color.Color { return o.values.Foreground }
