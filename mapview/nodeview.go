package mapview

import (
	"image"

	"github.com/codeation/impress"
	"github.com/codeation/tile/eventlink"
	"github.com/codeation/tile/view"
	"github.com/codeation/tile/view/box"
	"github.com/codeation/tile/view/fieldview"
	"github.com/codeation/tile/view/fn"
	"github.com/codeation/tile/view/solid"

	"github.com/codeation/lineation/mapmodel"
	"github.com/codeation/lineation/palette"
)

type nodeView struct {
	w         *impress.Window
	toDestroy bool
	viewer    view.Viewer
	size      image.Point
}

func (v *View) newNode(app eventlink.App, node *mapmodel.Node) *nodeView {
	return &nodeView{
		w:      app.NewWindow(image.Rectangle{}, v.pal.Color(palette.DefaultBackground)),
		viewer: v.nodeViewer(node),
	}
}

func (v *View) nodeViewer(node *mapmodel.Node) view.Viewer {
	isSelectedFn := func() bool {
		return node == v.m.Selected
	}
	isAccentedFn := func() bool {
		return node == v.m.Selected || node == v.droppingNode
	}
	isChangeableFn := func() bool {
		return node == v.m.Selected && node != v.draggingNode
	}
	nodeWidthFn := func() int {
		return v.pal.BoxWidth(node.Level()) - v.pal.BoxPadding().X*2
	}
	var viewer view.Viewer
	viewer = fieldview.New(fieldview.WithFocused(node.Value, isChangeableFn),
		v.pal.DefaultFont(), fn.Const(v.pal.Color(palette.DefaultText))).
		Width(nodeWidthFn).
		LineHeight(v.pal.TextLineHeight()).
		WithCursor(solid.New(v.pal.CursorSize(), fn.Const(v.pal.Color(palette.CursorBlock)))).
		MaxRows(v.pal.BoxRowCount())
	viewer = box.New(viewer).
		Padding(v.pal.BoxPadding()).
		Foreground(fn.If(isSelectedFn, v.pal.Color(palette.ActiveEdge), v.pal.Color(palette.DefaultEdge))).
		Background(fn.If(isAccentedFn, v.pal.Color(palette.ActiveBoxBackground), v.pal.Color(palette.DefaultBackground)))
	return viewer
}

func (n *nodeView) Destroy() {
	n.w.Drop()
}

func (n *nodeView) ResetSize(defaultSize image.Point) {
	n.size = n.viewer.Size()
}

func (n *nodeView) Size(defaultSize image.Point) image.Point {
	return n.size
}

func (n *nodeView) Draw(rect image.Rectangle) {
	n.w.Size(image.Rectangle{Min: rect.Min, Max: rect.Min.Add(n.size)})
	n.w.Clear()
	n.viewer.Draw(n.w, image.Pt(0, 0))
	n.w.Show()
}

func (n *nodeView) Raise() {
	n.w.Raise()
}
