package text

import (
	"image"
	"image/color"

	"github.com/codeation/impress"

	"github.com/codeation/lineation/draw/drawutil"
	"github.com/codeation/lineation/wrap"
)

type windower interface {
	NewWindow(rect image.Rectangle, background color.Color) *impress.Window
}

type TextStyler interface {
	Font() *impress.Font
	LineHeight() int
	Edge() int
	Margin() image.Point
	Size() image.Point
	Foreground() color.Color
	Background() color.Color
	Border() color.Color
}

type CursorStyler interface {
	Enable() bool
	Size() image.Point
	Foreground() color.Color
}

type option struct {
	Text   TextStyler
	Cursor CursorStyler
}

type textState struct {
	row, col       int
	enable         bool
	lastBackground color.Color
	width          int
}

type Text struct {
	content    *wrap.Wrap
	option     *option
	window     *impress.Window
	windowRect image.Rectangle
	textSize   image.Point
	lineSize   image.Point
	state      textState
}

func NewText(app windower, text string, textStyler TextStyler, cursorStyler CursorStyler,
) *Text {
	option := &option{
		Text:   textStyler,
		Cursor: cursorStyler,
	}
	content := wrap.NewWrap(text, option.Text.Font(), option.Text)
	content.End()
	textSize := image.Pt(option.Text.Size().X, option.Text.LineHeight()*1+option.Text.Margin().Y*2)
	windowRect := image.Rectangle{Min: image.Point{}, Max: textSize}
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

func (t *Text) stateOk() bool {
	row, col := t.content.Cursor()
	return row == t.state.row && col == t.state.col &&
		t.option.Cursor.Enable() == t.state.enable &&
		t.option.Text.Background() == t.state.lastBackground &&
		t.option.Text.Size().X == t.state.width
}

func (t *Text) keepState() {
	row, col := t.content.Cursor()
	t.state.row = row
	t.state.col = col
	t.state.enable = t.option.Cursor.Enable()
	t.state.lastBackground = t.option.Text.Background()
	t.state.width = t.option.Text.Size().X
}

func (t *Text) Show() {
	if t.stateOk() {
		return
	}
	row, col := t.content.Cursor()
	if (t.state.row != row || t.state.col != col || (t.lineSize.X == 0 && col != 0)) && t.option.Cursor.Enable() {
		t.lineSize = t.option.Text.Font().Size(t.content.Line(row)[:col])
	}
	t.keepState()
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

func (t *Text) Raise() {
	t.window.Raise()
}
