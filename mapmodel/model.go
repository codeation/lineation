package mapmodel

import (
	"context"

	"github.com/codeation/impress/event"
	"github.com/codeation/lineation/xmlfile"
	"github.com/codeation/tile/control"
	"github.com/codeation/tile/eventlink"
)

type MindMap struct {
	Root     *Node
	Selected *Node
	Filename string
}

func New(source *xmlfile.Node, filename string) *MindMap {
	root := newXMLNode(source, nil)
	root.Value.End()
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
	m.Selected.Value.Home()
	m.Selected = node
	m.Selected.Value.End()
}

func (m *MindMap) Every(fn func(node *Node)) {
	m.Root.every(fn)
}

func (m *MindMap) Down() {
	if node := m.Selected.down(); node != nil {
		m.Selected.Value.Home()
		m.Selected = node
		m.Selected.Value.End()
	}
}

func (m *MindMap) Up() {
	if node := m.Selected.up(); node != nil {
		m.Selected.Value.Home()
		m.Selected = node
		m.Selected.Value.End()
	}
}

func (m *MindMap) NewChildNode() {
	m.Selected.Value.Home()
	m.Selected = m.Selected.newChild()
}

func (m *MindMap) NewNextNode() {
	if m.Selected == m.Root {
		return
	}
	m.Selected.Value.Home()
	m.Selected = m.Selected.newNext()
}

func (m *MindMap) DeleteNode() {
	if m.Selected == m.Root {
		return
	}
	m.Selected.Value.Home()
	m.Selected = m.Selected.remove()
	m.Selected.Value.End()
}

func (m *MindMap) Adopt(node *Node, newParent *Node, beforeNode *Node) {
	newParent.insertBefore(node, beforeNode)
	m.Selected.Value.Home()
	m.Selected = node
	m.Selected.Value.End()
}

func (m *MindMap) NodeControl(ctx context.Context, app eventlink.App, e event.Eventer, prior control.DoFunc) {
	m.Selected.Control.Control(ctx, app, e, prior)
}
