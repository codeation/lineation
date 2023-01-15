package box

import (
	"image"
)

func (b *Box) Point() image.Point {
	return b.point
}

func (b *Box) width() int {
	return b.pal.BoxWidth(b.level())
}

func (b *Box) height() int {
	return b.pal.BoxHeight(b.level(), b.textBox.Lines())
}

func (b *Box) rect() image.Rectangle {
	return image.Rect(b.point.X, b.point.Y, b.point.X+b.width(), b.point.Y+b.height())
}

func (b *Box) heightWithChilds() int {
	level := b.level()
	switch {
	case level == 1:
		return b.height()
	case level < b.pal.Columns():
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

func (b *Box) SplitLeftRight() bool {
	shifted := false
	if len(b.childs) < 2 {
		for i := range b.childs {
			shifted = shifted || !b.childs[i].isRight
			b.childs[i].isRight = true
		}
		return shifted
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
		shifted = shifted || b.childs[i].isRight != (i >= minPos)
		b.childs[i].isRight = i >= minPos
	}
	return shifted
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

func (b *Box) getEdge(current image.Rectangle) image.Rectangle {
	rect := b.rect()
	output := image.Rect(minInt(current.Min.X, rect.Min.X), minInt(current.Min.Y, rect.Min.Y),
		maxInt(current.Max.X, rect.Max.X), maxInt(current.Max.Y, rect.Max.Y))
	for _, child := range b.childs {
		output = child.getEdge(output)
	}
	return output
}

func (b *Box) Fit(windowSize image.Point, offset image.Point) image.Point {
	mapRect := b.getEdge(b.rect())

	x := offset.X
	if mapRect.Dx() <= windowSize.X-2*b.pal.HorizontalBoxAlign() {
		if mapRect.Min.X < b.pal.HorizontalBoxAlign() {
			x = b.pal.HorizontalBoxAlign() - mapRect.Min.X
		} else if mapRect.Max.X > windowSize.X-b.pal.HorizontalBoxAlign() {
			x = windowSize.X - b.pal.HorizontalBoxAlign() - mapRect.Max.X
		} else {
			x = 0
		}
	} else if x < 0 && x < windowSize.X-b.pal.HorizontalBoxAlign()-mapRect.Max.X {
		x = windowSize.X - b.pal.HorizontalBoxAlign() - mapRect.Max.X
	} else if x > 0 && x > b.pal.HorizontalBoxAlign()-mapRect.Min.X {
		x = b.pal.HorizontalBoxAlign() - mapRect.Min.X
	}

	y := offset.Y
	if mapRect.Max.Y < windowSize.Y-b.pal.VerticalBoxAlign() {
		y = 0
	} else if y < 0 && y < windowSize.Y-b.pal.VerticalBoxAlign()-mapRect.Max.Y {
		y = windowSize.Y - b.pal.VerticalBoxAlign() - mapRect.Max.Y
	}

	return image.Pt(x, y)
}
