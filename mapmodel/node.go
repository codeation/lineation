package mapmodel

import (
	"slices"

	"github.com/codeation/tile/elem/field"

	"github.com/codeation/lineation/xmlfile"
)

type Node struct {
	Value  *field.Field
	Parent *Node
	Childs []*Node
}

func newXMLNode(elem *xmlfile.Node, parent *Node) *Node {
	node := &Node{
		Value:  field.New(elem.Text),
		Parent: parent,
		Childs: make([]*Node, 0, len(elem.Childs)),
	}
	if len(elem.Childs) == 0 {
		return node
	}
	node.Childs = make([]*Node, 0, len(elem.Childs))
	for _, child := range elem.Childs {
		node.Childs = append(node.Childs, newXMLNode(child, node))
	}
	return node
}

func (node *Node) export() *xmlfile.Node {
	output := &xmlfile.Node{
		Text: node.Value.String(),
	}
	for _, child := range node.Childs {
		output.Childs = append(output.Childs, child.export())
	}
	return output
}

func (parent *Node) newChild() *Node {
	newNode := &Node{
		Value:  field.New(""),
		Parent: parent,
	}
	parent.Childs = append([]*Node{newNode}, parent.Childs...)
	return newNode
}

func (node *Node) newNext() *Node {
	i := slices.Index(node.Parent.Childs, node)
	newNode := &Node{
		Value:  field.New(""),
		Parent: node.Parent,
	}
	node.Parent.Childs = slices.Insert(node.Parent.Childs, i+1, newNode)
	return newNode
}

func (node *Node) remove() *Node {
	i := slices.Index(node.Parent.Childs, node)
	node.Parent.Childs = slices.Delete(node.Parent.Childs, i, i+1)
	if len(node.Parent.Childs) >= i+1 {
		return node.Parent.Childs[i]
	}
	if len(node.Parent.Childs) != 0 {
		return node.Parent.Childs[len(node.Parent.Childs)-1]
	}
	return node.Parent
}

func (node *Node) insertBefore(newNode *Node, beforeNode *Node) {
	i := slices.Index(newNode.Parent.Childs, newNode)
	newNode.Parent.Childs = slices.Delete(newNode.Parent.Childs, i, i+1)
	newNode.Parent = node
	i = len(node.Childs)
	if beforeNode != nil {
		i = slices.Index(node.Childs, beforeNode)
	}
	node.Childs = slices.Insert(node.Childs, i, newNode)
}

func (node *Node) Level() int {
	level := 0
	for node != nil {
		level++
		node = node.Parent
	}
	return level
}

func (node *Node) next() *Node {
	if node.Parent == nil {
		return nil
	}
	i := slices.Index(node.Parent.Childs, node)
	if i < 0 || i+1 >= len(node.Parent.Childs) {
		return nil
	}
	return node.Parent.Childs[i+1]
}

func (node *Node) prev() *Node {
	if node.Parent == nil {
		return nil
	}
	i := slices.Index(node.Parent.Childs, node)
	if i <= 0 { // same as i < 0 || i-1 < 0
		return nil
	}
	return node.Parent.Childs[i-1]
}

func (node *Node) down() *Node {
	if len(node.Childs) != 0 {
		return node.Childs[0]
	}
	for cursor := node; cursor != nil; {
		if next := cursor.next(); next != nil {
			return next
		}
		cursor = cursor.Parent
	}
	return nil
}

func (node *Node) up() *Node {
	cursor := node.prev()
	if cursor == nil {
		return node.Parent
	}
	for len(cursor.Childs) != 0 {
		cursor = cursor.Childs[len(cursor.Childs)-1]
	}
	return cursor
}

func (node *Node) every(fn func(node *Node)) {
	fn(node)
	for _, child := range node.Childs {
		child.every(fn)
	}
}
