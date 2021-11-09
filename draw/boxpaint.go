package draw

import (
	"image"
	"image/color"

	"github.com/codeation/impress"

	"github.com/codeation/lineation/draw/drawutil"
	"github.com/codeation/lineation/palette"
)

func (b *Box) Align(since image.Point) {
	b.point = since.Add(image.Pt(-b.width()/2, 0))
	left := since.Add(image.Pt(-b.pal.HorizontalBoxOffset(b.level+1), 0))
	right := since.Add(image.Pt(b.pal.HorizontalBoxOffset(b.level+1), 0))
	if b.level > 2 {
		left = left.Add(image.Pt(0, b.height()+b.pal.VerticalBoxOffset()))
		right = right.Add(image.Pt(0, b.height()+b.pal.VerticalBoxOffset()))
	}
	for _, child := range b.childs {
		if child.IsRight() {
			child.Align(right)
			right = right.Add(image.Pt(0, child.heightWithChilds()+b.pal.VerticalBoxOffset()))
		} else {
			child.Align(left)
			left = left.Add(image.Pt(0, child.heightWithChilds()+b.pal.VerticalBoxOffset()))
		}
	}
}

func (b *Box) WarpText() {
	row, col := b.content.Cursor()
	if b.warpRow != row || b.warpCol != col {
		b.warpTextSize = b.pal.DefaultFont().Size(b.content.Line(row)[:col])
		b.warpRow = row
		b.warpCol = col
	}
}

func (b *Box) Draw(w *impress.Window, offset image.Point) {
	rect := b.rect()
	if b.isActive {
		w.Fill(rect.Add(offset), b.pal.Color(palette.ActiveBoxBackground))
		drawutil.DrawRectEdge(w, rect.Add(offset), b.pal.Color(palette.ActiveEdge))
	} else {
		drawutil.DrawRectEdge(w, rect.Add(offset), b.pal.Color(palette.DefaultEdge))
	}
	if b.isActive {
		cursorPoint := rect.Min.
			Add(offset).
			Add(b.pal.TextLinePoint(b.level, b.warpRow)).
			Add(image.Pt(b.warpTextSize.X, 0))
		w.Fill(
			image.Rect(0, 0, b.pal.CursorSize().X, b.pal.CursorSize().Y).
				Add(cursorPoint).
				Add(b.pal.CursorPoint()),
			b.pal.Color(palette.CursorBlock))
	}
	for i := 0; i < b.content.Lines(); i++ {
		w.Text(b.content.Line(i), b.pal.DefaultFont(),
			rect.Min.
				Add(offset).
				Add(b.pal.TextLinePoint(b.level, i)),
			b.pal.Color(palette.DefaultText))
	}
	b.drawLine(w, offset, b.pal.Color(palette.DefaultLine))
	for _, child := range b.childs {
		child.Draw(w, offset)
	}
}

func (b *Box) drawLine(w *impress.Window, offset image.Point, color color.Color) {
	if b.level <= 1 {
		return
	}
	if b.level == 2 || b.level == 3 {
		drawutil.DrawLine3Elem(w, offset, b.parent.lineFrom(b.IsRight()), b.lineTo(b.IsRight()), color)
		return
	}
	drawutil.DrawLine2Elem(w, offset, b.parent.lineFrom(b.IsRight()), b.lineTo(b.IsRight()), color)
}

func (b *Box) lineFrom(toRight bool) image.Point {
	rect := b.rect()
	if b.level == 1 || b.level == 2 {
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
	if b.level == 2 || b.level == 3 {
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
