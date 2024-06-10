package modified

import (
	"image"

	"github.com/codeation/impress"
	"github.com/codeation/lineation/palette"
	"github.com/codeation/tile/eventlink"
	"github.com/codeation/tile/view"
	"github.com/codeation/tile/view/fn"
	"github.com/codeation/tile/view/solid"
)

type View struct {
	*State
	w      *impress.Window
	viewer view.Viewer
}

func NewView(app eventlink.AppFramer, pal *palette.Palette) *View {
	state := NewState()
	size := image.Pt(pal.CursorSize().Y, pal.CursorSize().X)
	colorFn := fn.If(state.Get, fn.Const(pal.Color(palette.ActiveEdge)), fn.Const(pal.Color(palette.DefaultBackground)))
	return &View{
		State:  state,
		w:      app.NewWindow(image.Rectangle{Max: size}, pal.Color(palette.DefaultBackground)),
		viewer: solid.New(size, colorFn),
	}
}

func (v *View) Destroy() {
	v.w.Drop()
}

func (v *View) Draw() {
	v.w.Clear()
	v.viewer.Draw(v.w, image.Rectangle{Max: v.viewer.Size(image.Point{})})
	v.w.Show()
}
