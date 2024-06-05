package mapview

import (
	"github.com/codeation/lineation/mapmodel"
)

type Direction int

const (
	_ Direction = iota
	Left
	Right
)

type LRIndex struct {
	index  int
	length int
}

func NewLRIndex(index int, length int) LRIndex {
	return LRIndex{
		index:  index,
		length: length,
	}
}

func (index LRIndex) Direction(i int) Direction {
	if index.length <= 1 {
		return Right
	}
	if i <= int(index.index) {
		return Left
	}
	return Right
}

func (index LRIndex) LeftNodes(nodes []*mapmodel.Node) []*mapmodel.Node {
	if len(nodes) <= 1 {
		return nil
	}
	return nodes[:index.index+1]
}

func (index LRIndex) RightNodes(nodes []*mapmodel.Node) []*mapmodel.Node {
	if len(nodes) <= 1 {
		return nodes
	}
	return nodes[index.index+1:]
}
