package draw

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

func (b *Box) down() *Box {
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

func (b *Box) up() *Box {
	cursor := b.prev()
	if cursor == nil {
		return b.parent
	}
	for len(cursor.childs) != 0 {
		cursor = cursor.childs[len(cursor.childs)-1]
	}
	return cursor
}

func (b *Box) Find(point image.Point) *Box {
	if point.In(b.rect()) {
		return b
	}
	for _, child := range b.childs {
		if next := child.Find(point); next != nil {
			return next
		}
	}
	return nil
}

func (b *Box) IsRight() bool {
	for node := b; node.level > 1; node = node.parent {
		if node.level == 2 {
			return node.isRight
		}
	}
	return true
}
