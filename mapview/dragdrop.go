package mapview

import (
	"image"

	"github.com/codeation/lineation/mapmodel"
)

func (v *View) Drag(node *mapmodel.Node, from, to image.Point) {
	if v.draggingNode == nil {
		v.draggingNode = node
		lrIndex := v.nodesSplit()
		lr := v.nodeDirection(v.draggingNode, lrIndex)
		v.draggingRect = v.nodeRect(v.draggingNode, lrIndex, lr)
	}

	v.draggingOffset = to.Sub(from)
	v.droppingNode, _ = v.droppable()
}

func (v *View) droppable() (*mapmodel.Node, *mapmodel.Node) {
	nodePoint := v.draggingRect.Min.Add(image.Pt(v.draggingRect.Dx()/2, 0)).Add(v.draggingOffset)
	if node, ok := v.Select(nodePoint); ok {
		return node, nil
	}
	if node, ok := v.nearSelect(nodePoint); ok {
		return node.Parent, node
	}
	return nil, nil
}

func (v *View) Drop(from, to image.Point) {
	if v.draggingNode == nil {
		return
	}
	v.draggingOffset = to.Sub(from)
	var beforeNode *mapmodel.Node
	v.droppingNode, beforeNode = v.droppable()
	if v.droppingNode == nil {
		return
	}
	v.m.Adopt(v.draggingNode, v.droppingNode, beforeNode)
}

func (v *View) DragRelease() {
	v.draggingNode = nil
	v.droppingNode = nil
	v.raisingNode = nil
}
