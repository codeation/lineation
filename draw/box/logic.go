package box

import "image"

func (b *Box) SetActive(isActive bool) {
	if !b.isActive {
		b.textBox.End()
	}
	b.isActive = isActive
}

func (b *Box) prev() *Box {
	if b.parent == nil {
		return nil
	}
	for i, node := range b.parent.childs {
		if node == b {
			if i > 0 {
				return b.parent.childs[i-1]
			}
			break
		}
	}
	return nil
}

func (b *Box) next() *Box {
	if b.parent == nil {
		return nil
	}
	for i, node := range b.parent.childs {
		if node == b {
			if i < len(b.parent.childs)-1 {
				return b.parent.childs[i+1]
			}
		}
	}
	return nil
}

func (b *Box) Down() *Box {
	if len(b.childs) != 0 {
		return b.childs[0]
	}
	for cursor := b; cursor != nil; {
		if next := cursor.next(); next != nil {
			return next
		}
		cursor = cursor.parent
	}
	return nil
}

func (b *Box) Up() *Box {
	cursor := b.prev()
	if cursor == nil {
		return b.parent
	}
	for len(cursor.childs) != 0 {
		cursor = cursor.childs[len(cursor.childs)-1]
	}
	return cursor
}

func (b *Box) In(point image.Point) bool {
	return point.In(b.rect())
}

func (b *Box) find(point image.Point, ignoredBox *Box) *Box {
	if b == ignoredBox {
		return nil
	}
	if point.In(b.rect()) {
		return b
	}
	for _, child := range b.childs {
		if next := child.find(point, ignoredBox); next != nil {
			return next
		}
	}
	return nil
}

func (b *Box) Find(point image.Point) *Box {
	return b.find(point, nil)
}

func (b *Box) FindOther(point image.Point, other *Box) *Box {
	return b.find(point, other)
}

func (b *Box) IsRight() bool {
	for node := b; node != nil; node = node.parent {
		if node.parent != nil && node.parent.parent == nil {
			return node.isRight
		}
	}
	return true
}
