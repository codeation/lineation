package xmlfile

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Node struct {
	Text   string  `xml:"text,attr"`
	Childs []*Node `xml:"node"`
}

type xmlTree struct {
	XMLName string `xml:"map"`
	Root    *Node  `xml:"node"`
}

func Open(filename string) (*Node, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %w", err)
	}

	var mm xmlTree
	if err := xml.Unmarshal(data, &mm); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal: %w", err)
	}

	return mm.Root, nil
}

func Save(root *Node, filename string) error {
	data, err := xml.MarshalIndent(xmlTree{Root: root}, "", "    ")
	if err != nil {
		return fmt.Errorf("xml.MarshalIndent: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("os.WriteFile: %w", err)
	}

	return nil
}
