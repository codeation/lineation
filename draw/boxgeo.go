package draw

import (
	"github.com/codeation/impress"
)

func (b *Box) width() int {
	return b.pal.BoxWidth(b.level)
}

func (b *Box) height() int {
	return b.pal.BoxHeight(b.level, b.content.Lines())
}

func (b *Box) rect() impress.Rect {
	return b.point.Size(impress.NewSize(b.width(), b.height()))
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

func (b *Box) GetOffset(windowSize impress.Size, offset impress.Point) impress.Point {
	rect := b.rect()
	x := offset.X
	right := x + rect.X + rect.Width
	if right > windowSize.Width-b.pal.HorizontalBoxAlign() {
		x += windowSize.Width - b.pal.HorizontalBoxAlign() - right
	}
	left := x + rect.X
	if left < b.pal.HorizontalBoxAlign() {
		x += b.pal.HorizontalBoxAlign() - left
	}

	y := offset.Y
	bottom := y + rect.Y + rect.Height
	if bottom > windowSize.Height-b.pal.VerticalBoxAlign() {
		y += windowSize.Height - b.pal.VerticalBoxAlign() - bottom
	}
	top := y + rect.Y
	if top < b.pal.VerticalBoxAlign() {
		y += b.pal.VerticalBoxAlign() - top
	}

	return impress.NewPoint(x, y)
}

func (b *Box) getEdge(left, top, right, bottom int) (int, int, int, int) {
	rect := b.rect()
	left, top, right, bottom = minInt(left, rect.X), minInt(top, rect.Y),
		maxInt(right, rect.X+rect.Width), maxInt(bottom, rect.Y+rect.Height)
	for _, child := range b.childs {
		left, top, right, bottom = child.getEdge(left, top, right, bottom)
	}
	return left, top, right, bottom
}

func (b *Box) getEdgeSize() impress.Size {
	rect := b.rect()
	left, top, right, bottom := rect.X, rect.Y, rect.X+rect.Width, rect.Y+rect.Height
	left, top, right, bottom = b.getEdge(left, top, right, bottom)
	return impress.NewSize(right-left, bottom-top)
}

func (b *Box) Fit(windowSize impress.Size, offset impress.Point) impress.Point {
	mapSize := b.getEdgeSize()

	x := offset.X
	tailX := mapSize.Width + 2*b.pal.HorizontalBoxAlign() - windowSize.Width
	if tailX <= 0 {
		x = 0
	} else if x < 0 && -x > tailX/2 {
		x = -tailX / 2
	} else if x > 0 && x > tailX/2 {
		x = tailX / 2
	}

	y := offset.Y
	tailY := mapSize.Height + 2*b.pal.VerticalBoxAlign() - windowSize.Height
	if tailY <= 0 {
		y = 0
	} else if y < 0 && -y > tailY {
		y = -tailY
	} else if y > 0 && y > tailY {
		y = tailY
	}

	return impress.NewPoint(x, y)
}
