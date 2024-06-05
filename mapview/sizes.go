package mapview

import (
	"image"
	"slices"

	"github.com/codeation/lineation/mapmodel"
	"github.com/codeation/lineation/mapview/geo"
)

func (v *View) rootRect() image.Rectangle {
	rootWidth := v.pal.BoxWidth(1)
	return image.Rectangle{Max: image.Pt(rootWidth, 0)}.Add(image.Pt((v.rect.Dx()-rootWidth)/2, v.pal.BoxAlign().Y))

}

func (v *View) childRect(node *mapmodel.Node, rect image.Rectangle, lr Direction) image.Rectangle {
	switch {
	case len(node.Childs) == 0:
		return rect
	case lr == Left:
		childSize := image.Pt(v.pal.BoxWidth(node.Level()+1), 0)
		if node.Level() < v.pal.Columns() {
			return rect.Sub(childSize).Sub(geo.RowSpan(v.pal.BoxOffset()))
		}
		nodeSize := v.nodes[node].Size(image.Point{})
		return rect.Sub(childSize).Add(nodeSize).Add(geo.Mirror(v.pal.BoxOffset()))
	default: // lr == Right
		if node.Level() < v.pal.Columns() {
			nodeSize := image.Pt(v.pal.BoxWidth(node.Level()), 0)
			return rect.Add(nodeSize).Add(geo.RowSpan(v.pal.BoxOffset()))
		}
		nodeSize := v.nodes[node].Size(image.Point{})
		return rect.Add(geo.ColSpan(nodeSize)).Add(v.pal.BoxOffset())
	}
}

func (v *View) sizeWithChilds(node *mapmodel.Node) image.Point {
	if node.Level() == 1 || len(node.Childs) == 0 {
		return v.nodes[node].Size(image.Point{})
	}
	nodeSize := v.nodes[node].Size(image.Point{})
	childsSize := v.sizeOfChilds(node)
	if node.Level() < v.pal.Columns() {
		return geo.RowSize(nodeSize, geo.RowSpan(v.pal.BoxOffset()), childsSize)
	}
	return geo.ColSize(nodeSize, geo.ColSpan(v.pal.BoxOffset()), childsSize)
}

func (v *View) sizeOfChilds(node *mapmodel.Node) image.Point {
	output := geo.ColSpan(v.pal.BoxOffset().Mul(-1))
	for _, child := range node.Childs {
		childSize := v.sizeWithChilds(child)
		output = geo.ColSize(output, geo.ColSpan(v.pal.BoxOffset()), childSize)
	}
	return output
}

func (v *View) nodesSplit() LRIndex {
	if len(v.m.Root.Childs) <= 1 {
		return NewLRIndex(0, len(v.m.Root.Childs))
	}
	heights := make([]int, 0, len(v.m.Root.Childs))
	totalHeight := -v.pal.BoxOffset().Y
	for _, child := range v.m.Root.Childs {
		totalHeight += v.pal.BoxOffset().Y + v.sizeWithChilds(child).Y
		heights = append(heights, totalHeight)
	}
	halfHeight := (totalHeight - v.pal.BoxOffset().Y) / 2
	minIndex, minDelta := 0, totalHeight
	for i, height := range heights {
		if abs(height-halfHeight) < minDelta {
			minIndex, minDelta = i, abs(height-halfHeight)
		}
	}
	return NewLRIndex(minIndex, len(v.m.Root.Childs))
}

func (v *View) nodeDirection(node *mapmodel.Node, lrIndex LRIndex) Direction {
	for node.Parent != v.m.Root {
		node = node.Parent
	}
	return lrIndex.Direction(slices.Index(node.Parent.Childs, node))
}

func (v *View) nodeRect(node *mapmodel.Node, lrIndex LRIndex, lr Direction) image.Rectangle {
	rect := v.rootRect()
	for node.Parent != nil {
		rect = v.childRect(node.Parent, rect, lr)
		for i, child := range node.Parent.Childs {
			if child == node {
				break
			}
			if node.Parent == v.m.Root && lrIndex.Direction(i) != lr {
				continue
			}
			childSize := v.sizeWithChilds(child)
			rect = rect.Add(geo.ColSpan(childSize)).Add(geo.ColSpan(v.pal.BoxOffset()))
		}
		node = node.Parent
	}
	return rect
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
