package mindmap

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Node struct {
	Text   string  `xml:"text,attr"`
	Childs []*Node `xml:"node"`
}

type MindMapTree struct {
	XMLName string `xml:"map"`
	Root    *Node  `xml:"node"`
}

type MindMap struct {
	filename string
	root     *Node
}

func Open(filename string) (*MindMap, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: %w", err)
	}

	var tree MindMapTree
	if err := xml.Unmarshal(data, &tree); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	return NewMindMap(filename, tree.Root), nil
}

func NewMindMap(filename string, root *Node) *MindMap {
	return &MindMap{
		filename: filename,
		root:     root,
	}
}

func (mm *MindMap) Filename() string {
	return mm.filename
}

func (mm *MindMap) Root() *Node {
	return mm.root
}

func (mm *MindMap) Save() error {
	data, err := xml.MarshalIndent(&MindMapTree{Root: mm.root}, "", "    ")
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	if err := os.WriteFile(mm.filename, data, 0644); err != nil {
		return fmt.Errorf("WriteFile: %w", err)
	}

	return nil
}
