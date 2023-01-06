package box

import (
	"image"
	"image/color"

	"github.com/codeation/impress"

	"github.com/codeation/lineation/draw/drawutil"
	"github.com/codeation/lineation/palette"
)

func (b *Box) Align(since image.Point) bool {
	level := b.level()
	if level == 1 {
		since = since.Add(image.Pt(0, b.pal.VerticalBoxAlign()))
	}
	shifted := false
	boxPoint := since.Add(image.Pt(-b.width()/2, 0))
	if b.point != boxPoint {
		shifted = true
		b.point = boxPoint
	}
	left := since.Add(image.Pt(-b.pal.HorizontalBoxOffset(level+1), 0))
	right := since.Add(image.Pt(b.pal.HorizontalBoxOffset(level+1), 0))
	if level > 2 {
		left = left.Add(image.Pt(0, b.height()+b.pal.VerticalBoxOffset()))
		right = right.Add(image.Pt(0, b.height()+b.pal.VerticalBoxOffset()))
	}
	for _, child := range b.childs {
		if child.IsRight() {
			childShifted := child.Align(right)
			shifted = shifted || childShifted
			right = right.Add(image.Pt(0, child.heightWithChilds()+b.pal.VerticalBoxOffset()))
		} else {
			childShifted := child.Align(left)
			shifted = shifted || childShifted
			left = left.Add(image.Pt(0, child.heightWithChilds()+b.pal.VerticalBoxOffset()))
		}
	}
	return shifted
}

func (b *Box) Draw(offset image.Point) {
	b.textBox.Move(b.point.Add(offset))
	b.textBox.Show()
	for _, child := range b.childs {
		child.Draw(offset)
	}
}

func (b *Box) DrawGrid(w *impress.Window, offset image.Point, other *Box) {
	if b == other {
		return
	}
	b.drawLine(w, offset, b.pal.Color(palette.DefaultLine))
	for _, child := range b.childs {
		child.DrawGrid(w, offset, other)
	}
}

func (b *Box) drawLine(w *impress.Window, offset image.Point, color color.Color) {
	level := b.level()
	if level <= 1 {
		return
	}
	if level == 2 || level == 3 {
		drawutil.DrawLine3Elem(w, offset, b.parent.lineFrom(b.IsRight()), b.lineTo(b.IsRight()), color)
		return
	}
	drawutil.DrawLine2Elem(w, offset, b.parent.lineFrom(b.IsRight()), b.lineTo(b.IsRight()), color)
}

func (b *Box) lineFrom(toRight bool) image.Point {
	rect := b.rect()
	level := b.level()
	if level == 1 || level == 2 {
		if toRight {
			return image.Pt(rect.Max.X, rect.Min.Y+b.pal.VerticalTextAlign())
		}
		return image.Pt(rect.Min.X, rect.Min.Y+b.pal.VerticalTextAlign())
	}
	if toRight {
		return image.Pt(rect.Min.X+b.pal.HorizontalTextAlign(), rect.Max.Y)
	}
	return image.Pt(rect.Max.X-b.pal.HorizontalTextAlign(), rect.Max.Y)
}

func (b *Box) lineTo(toRight bool) image.Point {
	rect := b.rect()
	level := b.level()
	if level == 2 || level == 3 {
		if toRight {
			return image.Pt(rect.Min.X, rect.Min.Y+b.pal.VerticalTextAlign())
		}
		return image.Pt(rect.Max.X, rect.Min.Y+b.pal.VerticalTextAlign())
	}

	if toRight {
		return image.Pt(rect.Min.X, rect.Min.Y+b.pal.VerticalTextAlign())
	}
	return image.Pt(rect.Max.X, rect.Min.Y+b.pal.VerticalTextAlign())
}

func (b *Box) Raise() {
	b.textBox.Raise()
	for _, child := range b.childs {
		child.Raise()
	}
}

func (b *Box) Drag(plus image.Point) {
	b.textBox.Move(b.point.Add(plus))
	for _, child := range b.childs {
		child.Drag(plus)
	}
}
