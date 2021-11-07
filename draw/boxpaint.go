package draw

import (
	"github.com/codeation/impress"

	"github.com/codeation/lineation/draw/drawutil"
	"github.com/codeation/lineation/palette"
)

func (b *Box) Align(since impress.Point) {
	b.point = since.MoveX(-b.width() / 2)
	left := since.MoveX(-b.pal.HorizontalBoxOffset(b.level + 1))
	right := since.MoveX(b.pal.HorizontalBoxOffset(b.level + 1))
	if b.level > 2 {
		left = left.MoveY(b.height() + b.pal.VerticalBoxOffset())
		right = right.MoveY(b.height() + b.pal.VerticalBoxOffset())
	}
	for _, child := range b.childs {
		if child.IsRight() {
			child.Align(right)
			right = right.MoveY(child.heightWithChilds() + b.pal.VerticalBoxOffset())
		} else {
			child.Align(left)
			left = left.MoveY(child.heightWithChilds() + b.pal.VerticalBoxOffset())
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

func (b *Box) Draw(w *impress.Window, offset impress.Point) {
	rect := b.rect()
	if b.isActive {
		w.Fill(rect.Move(offset), b.pal.Color(palette.ActiveBoxBackground))
		drawutil.DrawRectEdge(w, rect.Move(offset), b.pal.Color(palette.ActiveEdge))
	} else {
		drawutil.DrawRectEdge(w, rect.Move(offset), b.pal.Color(palette.DefaultEdge))
	}
	if b.isActive {
		cursorPoint := rect.Point.Move(offset).Move(b.pal.TextLinePoint(b.level, b.warpRow)).MoveX(b.warpTextSize.Width)
		w.Fill(cursorPoint.Move(b.pal.CursorPoint()).Size(b.pal.CursorSize()), b.pal.Color(palette.CursorBlock))
	}
	for i := 0; i < b.content.Lines(); i++ {
		w.Text(b.content.Line(i), b.pal.DefaultFont(),
			rect.Point.Move(offset).Move(b.pal.TextLinePoint(b.level, i)),
			b.pal.Color(palette.DefaultText))
	}
	b.drawLine(w, offset, b.pal.Color(palette.DefaultLine))
	for _, child := range b.childs {
		child.Draw(w, offset)
	}
}

func (b *Box) drawLine(w *impress.Window, offset impress.Point, color impress.Color) {
	if b.level <= 1 {
		return
	}
	if b.level == 2 || b.level == 3 {
		drawutil.DrawLine3Elem(w, offset, b.parent.lineFrom(b.IsRight()), b.lineTo(b.IsRight()), color)
		return
	}
	drawutil.DrawLine2Elem(w, offset, b.parent.lineFrom(b.IsRight()), b.lineTo(b.IsRight()), color)
}

func (b *Box) lineFrom(toRight bool) impress.Point {
	rect := b.rect()
	if b.level == 1 || b.level == 2 {
		if toRight {
			return impress.NewPoint(rect.X+rect.Width, rect.Y+b.pal.VerticalTextAlign())
		}
		return impress.NewPoint(rect.X-1, rect.Y+b.pal.VerticalTextAlign())
	}
	if toRight {
		return impress.NewPoint(rect.X+b.pal.HorizontalTextAlign(), rect.Y+rect.Height)
	}
	return impress.NewPoint(rect.X+rect.Width-b.pal.HorizontalTextAlign(), rect.Y+rect.Height)
}

func (b *Box) lineTo(toRight bool) impress.Point {
	rect := b.rect()
	if b.level == 2 || b.level == 3 {
		if toRight {
			return impress.NewPoint(rect.X-1, rect.Y+b.pal.VerticalTextAlign())
		}
		return impress.NewPoint(rect.X+rect.Width, rect.Y+b.pal.VerticalTextAlign())
	}

	if toRight {
		return impress.NewPoint(rect.X-1, rect.Y+b.pal.VerticalTextAlign())
	}
	return impress.NewPoint(rect.X+rect.Width, rect.Y+b.pal.VerticalTextAlign())
}
