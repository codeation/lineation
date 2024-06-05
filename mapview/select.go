package mapview

import (
	"image"

	"github.com/codeation/lineation/mapmodel"
	"github.com/codeation/lineation/mapview/geo"
)

func (v *View) Select(point image.Point) (*mapmodel.Node, bool) {
	rect := v.rootRect()
	rootSize := v.nodes[v.m.Root].Size(image.Point{})
	if point.In(image.Rectangle{Min: rect.Min, Max: rect.Min.Add(rootSize)}) {
		return v.m.Root, true
	}
	lrIndex := v.nodesSplit()
	if node, ok := v.nodesSelect(lrIndex.LeftNodes(v.m.Root.Childs), v.childRect(v.m.Root, rect, Left), Left, point); ok {
		return node, true
	}
	if node, ok := v.nodesSelect(lrIndex.RightNodes(v.m.Root.Childs), v.childRect(v.m.Root, rect, Right), Right, point); ok {
		return node, true
	}
	return nil, false
}

func (v *View) nodesSelect(nodes []*mapmodel.Node, rect image.Rectangle, lr Direction, point image.Point) (*mapmodel.Node, bool) {
	for _, node := range nodes {
		if node != v.draggingNode {
			nodeSize := v.nodes[node].Size(image.Point{})
			if point.In(image.Rectangle{Min: rect.Min, Max: rect.Min.Add(nodeSize)}) {
				return node, true
			}
			if node, ok := v.nodesSelect(node.Childs, v.childRect(node, rect, lr), lr, point); ok {
				return node, true
			}
		}
		childSize := v.sizeWithChilds(node)
		rect = rect.Add(geo.ColSpan(childSize)).Add(geo.ColSpan(v.pal.BoxOffset()))
	}
	return nil, false
}

func (v *View) nearSelect(point image.Point) (*mapmodel.Node, bool) {
	rect := v.rootRect()
	lrIndex := v.nodesSplit()
	if node, ok := v.nodesNear(lrIndex.LeftNodes(v.m.Root.Childs), v.childRect(v.m.Root, rect, Left), Left, point); ok {
		return node, true
	}
	if node, ok := v.nodesNear(lrIndex.RightNodes(v.m.Root.Childs), v.childRect(v.m.Root, rect, Right), Right, point); ok {
		return node, true
	}
	return nil, false
}

func (v *View) nodesNear(nodes []*mapmodel.Node, rect image.Rectangle, lr Direction, point image.Point) (*mapmodel.Node, bool) {
	if len(nodes) == 0 {
		return nil, false
	}
	rowSpan := image.Pt(v.pal.BoxWidth(nodes[0].Level()), 0)
	colSpan := geo.ColSpan(v.pal.BoxOffset())
	for _, node := range nodes {
		if node != v.draggingNode {
			if point.In(image.Rectangle{Min: rect.Min.Sub(colSpan), Max: rect.Min.Add(rowSpan)}) {
				return node, true
			}
			if node, ok := v.nodesNear(node.Childs, v.childRect(node, rect, lr), lr, point); ok {
				return node, true
			}
		}
		childSize := v.sizeWithChilds(node)
		rect = rect.Add(geo.ColSpan(childSize)).Add(geo.ColSpan(v.pal.BoxOffset()))
	}
	return nil, false
}
