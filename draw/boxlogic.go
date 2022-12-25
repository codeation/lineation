package draw

import "image"

func (b *Box) SetActive(isActive bool) {
	if !b.isActive {
		b.textBox.End()
	}
	b.isActive = isActive
}

func (b *Box) down() *Box {
	if len(b.childs) != 0 {
		return b.childs[0]
	}
	for cursor := b; cursor != nil; {
		if cursor.next != nil {
			return cursor.next
		}
		cursor = cursor.parent
	}
	return nil
}

func (b *Box) up() *Box {
	if b.prev == nil {
		return b.parent
	}
	cursor := b.prev
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
