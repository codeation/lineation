package mapview

import (
	"image"
	"maps"

	"github.com/codeation/impress"
	"github.com/codeation/tile/eventlink"

	"github.com/codeation/lineation/mapmodel"
	"github.com/codeation/lineation/mapview/geo"
	"github.com/codeation/lineation/palette"
)

type View struct {
	m              *mapmodel.MindMap
	pal            *palette.Palette
	w              *impress.Window
	rect           image.Rectangle
	nodes          map[*mapmodel.Node]*nodeView
	draggingNode   *mapmodel.Node
	droppingNode   *mapmodel.Node
	draggingOffset image.Point
	draggingRect   image.Rectangle
	raisingNode    *mapmodel.Node
}

func New(app eventlink.AppFramer, m *mapmodel.MindMap, pal *palette.Palette) *View {
	setNodeFuncs(pal)
	return &View{
		m:     m,
		pal:   pal,
		w:     app.NewWindow(pal.DefaultAppRect(), pal.Color(palette.DefaultBackground)),
		rect:  pal.DefaultAppRect(),
		nodes: map[*mapmodel.Node]*nodeView{},
	}
}

func (v *View) Destroy() {
	for _, nodeView := range v.nodes {
		nodeView.Destroy()
	}
	v.w.Drop()
}

func (v *View) Configure(size image.Point) {
	v.rect = image.Rectangle{Max: size}
	v.w.Size(v.rect)
}

func (v *View) nodesReset(app eventlink.App) {
	for _, nodeView := range v.nodes {
		nodeView.toDestroy = true
	}
	v.m.Every(func(node *mapmodel.Node) {
		if nodeView, ok := v.nodes[node]; ok {
			nodeView.toDestroy = false
			nodeView.ResetSize(image.Point{})
			return
		}
		noveView := v.newNode(app, node)
		noveView.ResetSize(image.Point{})
		v.nodes[node] = noveView
	})
	maps.DeleteFunc(v.nodes, func(_ *mapmodel.Node, nodeView *nodeView) bool {
		if nodeView.toDestroy {
			nodeView.Destroy()
		}
		return nodeView.toDestroy
	})
}

func (v *View) Draw(app eventlink.App) {
	v.nodesReset(app)
	v.w.Clear()
	rect := v.rootRect()
	v.nodes[v.m.Root].Draw(rect)
	rootBorder := image.Rectangle{Min: rect.Min, Max: rect.Min.Add(v.nodes[v.m.Root].Size(image.Point{}))}
	lrIndex := v.nodesSplit()
	v.nodesDraw(lrIndex.LeftNodes(v.m.Root.Childs), v.childRect(v.m.Root, rect, Left), Left, rootBorder)
	v.nodesDraw(lrIndex.RightNodes(v.m.Root.Childs), v.childRect(v.m.Root, rect, Right), Right, rootBorder)

	if v.draggingNode != nil {
		if v.raisingNode != v.draggingNode {
			v.raisingNode = v.draggingNode
			v.nodeRaise(v.raisingNode)
		}

		lr := v.nodeDirection(v.draggingNode, lrIndex)
		rect := v.nodeRect(v.draggingNode, lrIndex, lr)
		rect = rect.Add(v.draggingOffset)
		v.nodes[v.draggingNode].Draw(rect)
		v.nodesDraw(v.draggingNode.Childs, v.childRect(v.draggingNode, rect, lr), lr, image.Rectangle{})
	}

	v.w.Show()
}

func (v *View) nodesDraw(nodes []*mapmodel.Node, rect image.Rectangle, lr Direction, parentBorder image.Rectangle) {
	for _, node := range nodes {
		if node != v.draggingNode {
			v.nodes[node].Draw(rect)
			nodeBorder := image.Rectangle{}
			if !parentBorder.Eq(nodeBorder) { // parentBorder != nil
				nodeBorder = image.Rectangle{Min: rect.Min, Max: rect.Min.Add(v.nodes[node].Size(image.Point{}))}
				v.arrowDraw(parentBorder, nodeBorder)
			}
			v.nodesDraw(node.Childs, v.childRect(node, rect, lr), lr, nodeBorder)
		}
		childSize := v.sizeWithChilds(node)
		rect = rect.Add(geo.ColSpan(childSize)).Add(geo.ColSpan(v.pal.BoxOffset()))
	}
}

func (v *View) nodeRaise(node *mapmodel.Node) {
	for _, child := range node.Childs {
		v.nodeRaise(child)
	}
	v.nodes[node].Raise()
}
