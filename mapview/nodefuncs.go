package mapview

import (
	"github.com/codeation/tile/view/fn"
	"github.com/codeation/tile/view/solid"

	"github.com/codeation/lineation/palette"
)

var (
	cursor                 *solid.Solid
	defaultTextColor       fn.Color
	textMargin             fn.Point
	activeForegroundColor  fn.Color
	defaultForegroundColor fn.Color
	activeBackgroundColor  fn.Color
	defaultBackgroundColor fn.Color
)

func setNodeFuncs(pal *palette.Palette) {
	cursor = solid.New(pal.CursorSize(), fn.Const(pal.Color(palette.CursorBlock)))
	defaultTextColor = fn.Const(pal.Color(palette.DefaultText))
	textMargin = pal.TextAlign
	activeForegroundColor = fn.Const(pal.Color(palette.ActiveEdge))
	defaultForegroundColor = fn.Const(pal.Color(palette.DefaultEdge))
	activeBackgroundColor = fn.Const(pal.Color(palette.ActiveBoxBackground))
	defaultBackgroundColor = fn.Const(pal.Color(palette.DefaultBackground))
}
