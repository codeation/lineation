package mapmodel

import (
	"github.com/codeation/lineation/xmlfile"
)

type MindMap struct {
	Root     *Node
	Selected *Node
	Filename string
}

func New(source *xmlfile.Node, filename string) *MindMap {
	root := newXMLNode(source, nil)
	return &MindMap{
		Root:     root,
		Selected: root,
		Filename: filename,
	}
}

func (m *MindMap) Export() *xmlfile.Node {
	return m.Root.export()
}

func (m *MindMap) Select(node *Node) {
	m.Selected = node
	m.Selected.Value.End()
}

func (m *MindMap) Every(fn func(node *Node)) {
	m.Root.every(fn)
}

func (m *MindMap) Down() {
	if node := m.Selected.down(); node != nil {
		m.Selected = node
		m.Selected.Value.End()
	}
}

func (m *MindMap) Up() {
	if node := m.Selected.up(); node != nil {
		m.Selected = node
		m.Selected.Value.End()
	}
}

func (m *MindMap) NewChildNode() {
	m.Selected = m.Selected.newChild()
}

func (m *MindMap) NewNextNode() {
	if m.Selected == m.Root {
		return
	}
	m.Selected = m.Selected.newNext()
}

func (m *MindMap) DeleteNode() {
	if m.Selected == m.Root {
		return
	}
	m.Selected = m.Selected.remove()
	m.Selected.Value.End()
}

func (m *MindMap) Adopt(node *Node, newParent *Node, beforeNode *Node) {
	newParent.insertBefore(node, beforeNode)
	m.Selected = node
	m.Selected.Value.End()
}

func (m *MindMap) Right() {
	m.Selected.Value.Right()
}

func (m *MindMap) Left() {
	m.Selected.Value.Left()
}

func (m *MindMap) Insert(alpha rune) {
	m.Selected.Value.Insert(alpha)
}

func (m *MindMap) Backspace() {
	m.Selected.Value.Backspace()
}
