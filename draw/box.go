package draw

import (
	"image"

	"github.com/codeation/impress"
	"github.com/codeation/lineation/mindmap"
	"github.com/codeation/lineation/palette"
	"github.com/codeation/lineation/text"
)

type Box struct {
	textBox      *text.Text
	textOption   *text.TextOption
	cursorOption *text.CursorOption
	point        image.Point
	level        int
	isActive     bool
	isRight      bool
	parent       *Box
	childs       []*Box
	pal          *palette.Palette
	canvas       *impress.Application
}

func NewBox(root *mindmap.Node, canvas *impress.Application, pal *palette.Palette) *Box {
	return newBoxNode(root, nil, 1, canvas, pal)
}

func newTextOption(level int, pal *palette.Palette) *text.TextOption {
	return &text.TextOption{
		Font:       pal.DefaultFont(),
		LineHeight: pal.DefaultFont().Height + pal.TextLineOffset(),
		Margin:     image.Pt(pal.HorizontalTextAlign(), pal.VerticalTextAlign()),
		Size:       image.Pt(pal.BoxWidth(level), 100),
		Foreground: pal.Color(palette.DefaultText),
		Background: pal.Color(palette.ActiveBoxBackground),
		Border:     pal.Color(palette.ActiveEdge),
	}
}

func newCursorOption(pal *palette.Palette) *text.CursorOption {
	return &text.CursorOption{
		Foreground: pal.Color(palette.CursorBlock),
		Size:       pal.CursorSize(),
	}
}

func newBoxNode(node *mindmap.Node, parent *Box, level int, canvas *impress.Application, pal *palette.Palette) *Box {
	textOption := newTextOption(level, pal)
	cursorOption := newCursorOption(pal)
	textBox := text.NewText(canvas, image.Pt(40, 40), node.Text,
		text.NewSimpleTextOption(textOption),
		text.NewSimpleCursorOption(cursorOption))
	b := &Box{
		textBox:      textBox,
		textOption:   textOption,
		cursorOption: cursorOption,
		parent:       parent,
		level:        level,
		pal:          pal,
		isActive:     level == 1,
		canvas:       canvas,
	}
	b.textBox.End()
	for _, child := range node.Childs {
		b.childs = append(b.childs, newBoxNode(child, b, level+1, canvas, pal))
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

func (b *Box) AddChildNode() *Box {
	textOption := newTextOption(b.level+1, b.pal)
	cursorOption := newCursorOption(b.pal)
	textBox := text.NewText(b.canvas, image.Pt(40, 40), "",
		text.NewSimpleTextOption(textOption),
		text.NewSimpleCursorOption(cursorOption))

	next := &Box{
		textOption:   textOption,
		cursorOption: cursorOption,
		textBox:      textBox,
		parent:       b,
		level:        b.level + 1,
		pal:          b.pal,
		canvas:       b.canvas,
	}
	if len(b.childs) == 0 {
		b.childs = []*Box{next}
		return next
	}
	b.childs = append([]*Box{next}, b.childs...)
	return next
}

func (b *Box) AddNextNode() *Box {
	textOption := newTextOption(b.level, b.pal)
	cursorOption := newCursorOption(b.pal)
	textBox := text.NewText(b.canvas, image.Pt(40, 40), "",
		text.NewSimpleTextOption(textOption),
		text.NewSimpleCursorOption(cursorOption))
	next := &Box{
		textOption:   textOption,
		cursorOption: cursorOption,
		textBox:      textBox,
		parent:       b.parent,
		level:        b.level,
		pal:          b.pal,
		canvas:       b.canvas,
	}
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

func (b *Box) Insert(alpha rune) {
	b.textBox.Insert(alpha)
}

func (b *Box) Backspace() bool {
	return b.textBox.Backspace()
}
