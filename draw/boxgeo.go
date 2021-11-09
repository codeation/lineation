package draw

import (
	"image"
)

func (b *Box) width() int {
	return b.pal.BoxWidth(b.level)
}

func (b *Box) height() int {
	return b.pal.BoxHeight(b.level, b.content.Lines())
}

func (b *Box) rect() image.Rectangle {
	return image.Rect(b.point.X, b.point.Y, b.point.X+b.width(), b.point.Y+b.height())
}

func (b *Box) heightWithChilds() int {
	switch b.level {
	case 1:
		return b.height()
	case 2:
		return maxInt(b.height(), b.heightOfChilds())
	default:
		heightOfChilds := b.heightOfChilds()
		if heightOfChilds != 0 {
			heightOfChilds += b.pal.VerticalBoxOffset()
		}
		return b.height() + heightOfChilds
	}
}

func (b *Box) heightOfChilds() int {
	output := 0
	for i, child := range b.childs {
		if i != 0 {
			output += b.pal.VerticalBoxOffset()
		}
		output += child.heightWithChilds()
	}
	return output
}

func (b *Box) SplitLeftRight() {
	if len(b.childs) < 2 {
		for i := range b.childs {
			b.childs[i].isRight = true
		}
		return
	}

	heights := make([]int, len(b.childs))
	for i, child := range b.childs {
		heights[i] = child.heightWithChilds()
	}
	heads := make([]int, len(heights)+1)
	sum := -b.pal.VerticalBoxOffset()
	for i := range heights {
		sum += heights[i] + b.pal.VerticalBoxOffset()
		heads[i+1] = sum
	}
	tails := make([]int, len(heights)+1)
	sum = -b.pal.VerticalBoxOffset()
	for i := len(heights) - 1; i >= 0; i-- {
		sum += heights[i] + b.pal.VerticalBoxOffset()
		tails[i] = sum
	}

	minDiff := absInt(heads[0] - tails[0])
	minPos := 0
	for i := range heads {
		if absInt(heads[i]-tails[i]) < minDiff {
			minDiff = absInt(heads[i] - tails[i])
			minPos = i
		}
	}

	for i := range b.childs {
		b.childs[i].isRight = i >= minPos
	}
}

func (b *Box) GetOffset(windowSize image.Point, offset image.Point) image.Point {
	rect := b.rect()
	x := offset.X
	right := x + rect.Max.X
	if right > windowSize.X-b.pal.HorizontalBoxAlign() {
		x += windowSize.X - b.pal.HorizontalBoxAlign() - right
	}
	left := x + rect.Min.X
	if left < b.pal.HorizontalBoxAlign() {
		x += b.pal.HorizontalBoxAlign() - left
	}

	y := offset.Y
	bottom := y + rect.Max.Y
	if bottom > windowSize.Y-b.pal.VerticalBoxAlign() {
		y += windowSize.Y - b.pal.VerticalBoxAlign() - bottom
	}
	top := y + rect.Min.Y
	if top < b.pal.VerticalBoxAlign() {
		y += b.pal.VerticalBoxAlign() - top
	}

	return image.Pt(x, y)
}

func (b *Box) getEdge(left, top, right, bottom int) (int, int, int, int) {
	rect := b.rect()
	left, top, right, bottom = minInt(left, rect.Min.X), minInt(top, rect.Min.Y),
		maxInt(right, rect.Max.X), maxInt(bottom, rect.Max.Y)
	for _, child := range b.childs {
		left, top, right, bottom = child.getEdge(left, top, right, bottom)
	}
	return left, top, right, bottom
}

func (b *Box) getEdgeSize() image.Point {
	rect := b.rect()
	left, top, right, bottom := rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y
	left, top, right, bottom = b.getEdge(left, top, right, bottom)
	return image.Pt(right-left, bottom-top)
}

func (b *Box) Fit(windowSize image.Point, offset image.Point) image.Point {
	mapSize := b.getEdgeSize()

	x := offset.X
	tailX := mapSize.X + 2*b.pal.HorizontalBoxAlign() - windowSize.X
	if tailX <= 0 {
		x = 0
	} else if x < 0 && -x > tailX/2 {
		x = -tailX / 2
	} else if x > 0 && x > tailX/2 {
		x = tailX / 2
	}

	y := offset.Y
	tailY := mapSize.Y + 2*b.pal.VerticalBoxAlign() - windowSize.Y
	if tailY <= 0 {
		y = 0
	} else if y < 0 && -y > tailY {
		y = -tailY
	} else if y > 0 && y > tailY {
		y = tailY
	}

	return image.Pt(x, y)
}
