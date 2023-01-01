package text

import (
	"image"
	"image/color"

	"github.com/codeation/impress"
	"github.com/codeation/lineation/draw/drawutil"
	"github.com/codeation/lineation/wrap"
)

type textStyler interface {
	Font() *impress.Font
	LineHeight() int
	Margin() image.Point
	Size() image.Point
	Foreground() color.Color
	Background() color.Color
	Border() color.Color
}

type cursorStyler interface {
	Enable() bool
	Size() image.Point
	Foreground() color.Color
}

type option struct {
	Text   textStyler
	Cursor cursorStyler
}

type Text struct {
	content    *wrap.Wrap
	option     *option
	window     *impress.Window
	windowRect image.Rectangle
	textSize   image.Point
	row, col   int
	enable     bool
	lineSize   image.Point
}

func NewText(app *impress.Application, from image.Point, text string,
	textStyler textStyler, cursorStyler cursorStyler,
) *Text {
	option := &option{
		Text:   textStyler,
		Cursor: cursorStyler,
	}
	content := wrap.NewWrap(text, option.Text.Font(), option.Text.Size().X-option.Text.Margin().X*2)
	content.End()
	textSize := image.Pt(option.Text.Size().X, option.Text.LineHeight()*1+option.Text.Margin().Y*2)
	windowRect := image.Rectangle{Min: from, Max: from.Add(textSize)}
	window := app.NewWindow(windowRect, option.Text.Background())
	return &Text{
		content:    content,
		option:     option,
		window:     window,
		windowRect: windowRect,
		textSize:   textSize,
	}
}

func (t *Text) Drop() {
	t.window.Drop()
}

func (t *Text) Show() {
	row, col := t.content.Cursor()
	if t.row == row && t.col == col && t.enable == t.option.Cursor.Enable() {
		return
	}
	if (t.row != row || t.col != col || (t.lineSize.X == 0 && col != 0)) && t.option.Cursor.Enable() {
		t.lineSize = t.option.Text.Font().Size(t.content.Line(row)[:col])
	}
	t.row = row
	t.col = col
	t.enable = t.option.Cursor.Enable()
	t.window.Clear()
	drawutil.DrawRectEdge(
		t.window,
		image.Rectangle{Min: image.Pt(0, 0), Max: t.option.Text.Size().Add(image.Pt(-1, -1))},
		t.option.Text.Border())
	t.window.Fill(
		image.Rectangle{
			Min: image.Pt(1, 1),
			Max: t.option.Text.Size().Add(image.Pt(-1, -1)),
		},
		t.option.Text.Background())

	if t.option.Cursor.Enable() {
		pt := t.option.Text.Margin().Add(image.Pt(t.lineSize.X, (row+1)*t.option.Text.LineHeight()-t.option.Text.Font().Height))
		t.window.Fill(image.Rectangle{Min: pt, Max: pt.Add(t.option.Cursor.Size())}, t.option.Cursor.Foreground())
	}
	for i := 0; i < t.content.Lines(); i++ {
		t.window.Text(
			t.content.Line(i),
			t.option.Text.Font(),
			t.option.Text.Margin().Add(image.Pt(0, i*t.option.Text.LineHeight())),
			t.option.Text.Foreground())
	}
	t.window.Show()
}

func (t *Text) Move(to image.Point) {
	windowRect := image.Rectangle{Min: to, Max: to.Add(t.option.Text.Size())}
	if windowRect == t.windowRect {
		return
	}
	t.windowRect = windowRect
	t.window.Size(t.windowRect)
}
