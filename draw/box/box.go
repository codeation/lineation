package box

import (
	"image"

	"github.com/codeation/impress"

	"github.com/codeation/lineation/mindmap"
	"github.com/codeation/lineation/palette"
	"github.com/codeation/lineation/text"
)

type Box struct {
	textBox    *text.Text
	point      image.Point
	isActive   bool
	isRight    bool
	emphasized bool
	parent     *Box
	childs     []*Box
	pal        *palette.Palette
}

func NewBox(root *mindmap.Node, app *impress.Application, pal *palette.Palette) *Box {
	return newBoxNode(root, nil, app, pal)
}

func newBoxNode(node *mindmap.Node, parent *Box, app *impress.Application, pal *palette.Palette) *Box {
	b := &Box{
		parent:   parent,
		isActive: parent == nil,
		pal:      pal,
	}
	b.textBox = text.NewText(app, node.Text, newTextOption(b), newCursorOption(b))
	b.textBox.End()
	for _, child := range node.Childs {
		b.childs = append(b.childs, newBoxNode(child, b, app, pal))
	}
	return b
}

func (b *Box) GetNodes() *mindmap.Node {
	node := &mindmap.Node{
		Text: b.textBox.String(),
	}
	for _, child := range b.childs {
		node.Childs = append(node.Childs, child.GetNodes())
	}
	return node
}

func (b *Box) level() int {
	level := 0
	for node := b; node != nil; node = node.parent {
		level++
	}
	return level
}

func (b *Box) AddChildNode(app *impress.Application) *Box {
	next := &Box{
		parent: b,
		pal:    b.pal,
	}
	next.textBox = text.NewText(app, "", newTextOption(next), newCursorOption(next))
	if len(b.childs) == 0 {
		b.childs = []*Box{next}
		return next
	}
	b.childs = append([]*Box{next}, b.childs...)
	return next
}

func (b *Box) AddNextNode(app *impress.Application) *Box {
	next := &Box{
		parent: b.parent,
		pal:    b.pal,
	}
	next.textBox = text.NewText(app, "", newTextOption(next), newCursorOption(next))
	childs := make([]*Box, 0, len(b.parent.childs)+1)
	for _, node := range b.parent.childs {
		childs = append(childs, node)
		if node == b {
			childs = append(childs, next)
		}
	}
	b.parent.childs = childs
	return next
}

func (b *Box) DeleteNode() *Box {
	var nextPos int
	childs := make([]*Box, 0, len(b.parent.childs)-1)
	for i, node := range b.parent.childs {
		if node == b {
			nextPos = i
			continue
		}
		childs = append(childs, node)
	}
	b.parent.childs = childs
	b.deleteChildNodes()
	b.parent.point = image.Point{}
	b.textBox.Drop()
	if nextPos < len(b.parent.childs) {
		return b.parent.childs[nextPos]
	} else if len(b.parent.childs) > 0 {
		return b.parent.childs[len(b.parent.childs)-1]
	}
	return b.parent
}

func (b *Box) deleteChildNodes() {
	for _, child := range b.childs {
		child.deleteChildNodes()
		child.textBox.Drop()
	}
}

func (b *Box) Right() bool {
	return b.textBox.Right()
}

func (b *Box) Left() bool {
	return b.textBox.Left()
}

func (b *Box) Insert(alpha rune) {
	b.textBox.Insert(alpha)
}

func (b *Box) Backspace() bool {
	return b.textBox.Backspace()
}

func (b *Box) Adopt(child *Box, beforeBox *Box) {
	if child.parent != nil {
		childs := make([]*Box, 0, len(child.parent.childs)-1)
		for _, node := range child.parent.childs {
			if node == child {
				continue
			}
			childs = append(childs, node)
		}
		child.parent.childs = childs
	}
	b.childs = append(b.childs, child)
	child.parent = b
	child.point = image.Point{}
	if beforeBox != nil {
		childs := make([]*Box, 0, len(b.childs))
		for _, node := range b.childs {
			if node == child {
				continue
			}
			if node == beforeBox {
				childs = append(childs, child)
			}
			childs = append(childs, node)
		}
		b.childs = childs
	}
}

func (b *Box) Emphasize() {
	b.emphasized = true
	b.textBox.Show()
}

func (b *Box) DeEmphasize() {
	b.emphasized = false
	b.textBox.Show()
}
