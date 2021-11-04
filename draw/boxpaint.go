package draw

import (
	"github.com/codeation/impress"

	"github.com/codeation/lineation/draw/drawutil"
	"github.com/codeation/lineation/palette"
)

func (b *Box) Align(since impress.Point) {
	b.rect = since.MoveX(-b.width() / 2).Size(b.rect.Size)
	b.markRight()
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
	if b.level != 1 {
		return
	}
}

func (b *Box) Draw(w *impress.Window, offset impress.Point) {
	if b.isActive {
		w.Fill(b.rect.Move(offset), b.pal.Color(palette.ActiveBoxBackground))
		drawutil.DrawRectEdge(w, b.rect.Move(offset), b.pal.Color(palette.ActiveEdge))
	} else {
		drawutil.DrawRectEdge(w, b.rect.Move(offset), b.pal.Color(palette.DefaultEdge))
	}
	for i := range b.texts {
		w.Text(b.texts[i], b.pal.DefaultFont(),
			b.rect.Point.Move(offset).Move(b.pal.TextLinePoint(b.level, i)),
			b.pal.Color(palette.DefaultText))
	}
	if b.isActive {
		if true || (b.cursorPoint.X == 0 && b.cursorPoint.Y == 0) {
			var textSize impress.Size
			if len(b.texts) != 0 {
				textSize = b.pal.DefaultFont().Size(b.texts[len(b.texts)-1])
			}
			b.cursorPoint = b.rect.Point.Move(offset).Move(b.pal.TextLinePoint(b.level, len(b.texts)-1)).MoveX(textSize.Width)
		}
		w.Fill(b.cursorPoint.Move(b.pal.CursorPoint()).Size(b.pal.CursorSize()), b.pal.Color(palette.CursorBlock))
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
	if b.level == 1 || b.level == 2 {
		if toRight {
			return impress.NewPoint(b.rect.X+b.rect.Width, b.rect.Y+b.pal.VerticalTextAlign())
		}
		return impress.NewPoint(b.rect.X-1, b.rect.Y+b.pal.VerticalTextAlign())
	}
	if toRight {
		return impress.NewPoint(b.rect.X+b.pal.HorizontalTextAlign(), b.rect.Y+b.rect.Height)
	}
	return impress.NewPoint(b.rect.X+b.rect.Width-b.pal.HorizontalTextAlign(), b.rect.Y+b.rect.Height)
}

func (b *Box) lineTo(toRight bool) impress.Point {
	if b.level == 2 || b.level == 3 {
		if toRight {
			return impress.NewPoint(b.rect.X-1, b.rect.Y+b.pal.VerticalTextAlign())
		}
		return impress.NewPoint(b.rect.X+b.rect.Width, b.rect.Y+b.pal.VerticalTextAlign())
	}

	if toRight {
		return impress.NewPoint(b.rect.X-1, b.rect.Y+b.pal.VerticalTextAlign())
	}
	return impress.NewPoint(b.rect.X+b.rect.Width, b.rect.Y+b.pal.VerticalTextAlign())
}
