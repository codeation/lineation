package draw

import (
	"github.com/codeation/impress"

	"github.com/codeation/lineation/mindmap"
	"github.com/codeation/lineation/palette"
	"github.com/codeation/lineation/wrap"
)

type Box struct {
	content          *wrap.Wrap
	point            impress.Point
	warpRow, warpCol int
	warpTextSize     impress.Size
	level            int
	isActive         bool
	isRight          bool
	parent           *Box
	next, prev       *Box
	childs           []*Box
	pal              *palette.Palette
}

func NewBox(root *mindmap.Node, pal *palette.Palette) *Box {
	return newBoxNode(root, nil, 1, pal)
}

func newBoxNode(node *mindmap.Node, parent *Box, level int, pal *palette.Palette) *Box {
	b := &Box{
		content:  wrap.NewWrap(node.Text, pal.DefaultFont(), pal.BoxWidth(level)-pal.HorizontalTextAlign()*2),
		parent:   parent,
		level:    level,
		pal:      pal,
		isActive: level == 1,
	}
	b.content.End()
	for _, child := range node.Childs {
		b.childs = append(b.childs, newBoxNode(child, b, level+1, pal))
	}
	for i := range b.childs {
		if i > 0 {
			b.childs[i].prev = b.childs[i-1]
		}
		if i < len(b.childs)-1 {
			b.childs[i].next = b.childs[i+1]
		}
	}
	return b
}

func (b *Box) GetNodes() *mindmap.Node {
	node := &mindmap.Node{
		Text: b.content.String(),
	}
	for _, child := range b.childs {
		node.Childs = append(node.Childs, child.GetNodes())
	}
	return node
}

func (b *Box) AddChildNode() *Box {
	next := &Box{
		content: wrap.NewWrap("", b.pal.DefaultFont(), b.pal.BoxWidth(b.level+1)-b.pal.HorizontalTextAlign()*2),
		parent:  b,
		level:   b.level + 1,
		pal:     b.pal,
	}
	if len(b.childs) == 0 {
		b.childs = []*Box{next}
		return next
	}
	next.next = b.childs[0]
	b.childs[0].prev = next
	b.childs = append([]*Box{next}, b.childs...)
	return next
}

func (b *Box) AddNextNode() *Box {
	next := &Box{
		content: wrap.NewWrap("", b.pal.DefaultFont(), b.pal.BoxWidth(b.level)-b.pal.HorizontalTextAlign()*2),
		parent:  b.parent,
		level:   b.level,
		pal:     b.pal,
		prev:    b,
		next:    b.next,
	}
	if b.next != nil {
		b.next.prev = next
	}
	b.next = next
	childs := make([]*Box, 0, len(b.parent.childs)+1)
	for node := b.parent.childs[0]; node != nil; node = node.next {
		childs = append(childs, node)
	}
	b.parent.childs = childs
	return next
}

func (b *Box) DeleteNode() *Box {
	next := b.next
	if next == nil {
		next = b.prev
	}
	if next == nil {
		next = b.parent
	}
	childs := make([]*Box, 0, len(b.parent.childs)-1)
	for node := b.parent.childs[0]; node != nil; node = node.next {
		if node == b {
			continue
		}
		childs = append(childs, node)
	}
	b.parent.childs = childs
	if b.prev != nil {
		b.prev.next = b.next
	}
	if b.next != nil {
		b.next.prev = b.prev
	}
	return next
}

func (b *Box) Insert(alpha rune) {
	b.content.Insert(alpha)
}

func (b *Box) Backspace() bool {
	return b.content.Backspace()
}
