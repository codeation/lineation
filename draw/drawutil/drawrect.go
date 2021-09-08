package drawutil

import (
	"github.com/codeation/impress"
)

func DrawRectEdge(w *impress.Window, rect impress.Rect, color impress.Color) {
	pointX0Y0 := impress.NewPoint(rect.X, rect.Y)
	pointX1Y0 := impress.NewPoint(rect.X+rect.Width-1, rect.Y)
	pointX0Y1 := impress.NewPoint(rect.X, rect.Y+rect.Height-1)
	pointX1Y1 := impress.NewPoint(rect.X+rect.Width-1, rect.Y+rect.Height-1)

	w.Line(pointX0Y0, pointX1Y0, color)
	w.Line(pointX1Y0, pointX1Y1, color)
	w.Line(pointX0Y1, pointX1Y1, color)
	w.Line(pointX0Y0, pointX0Y1, color)
}

func DrawLine3Elem(w *impress.Window, offset impress.Point, from impress.Point, to impress.Point, color impress.Color) {
	drawLine3Elem(w, from.Move(offset), to.Move(offset), color)
}

func drawLine3Elem(w *impress.Window, from impress.Point, to impress.Point, color impress.Color) {
	if from.Y == to.Y {
		w.Line(from, to, color)
		return
	}
	width := to.X - from.X
	width1 := width / 2
	width2 := width - width1
	w.Line(from, from.MoveX(width1), color)
	w.Line(from.MoveX(width1), to.MoveX(-width2), color)
	w.Line(to.MoveX(-width2), to, color)
}

func DrawLine2Elem(w *impress.Window, offset impress.Point, from impress.Point, to impress.Point, color impress.Color) {
	drawLine2Elem(w, from.Move(offset), to.Move(offset), color)
}

func drawLine2Elem(w *impress.Window, from impress.Point, to impress.Point, color impress.Color) {
	middle := impress.NewPoint(from.X, to.Y)
	w.Line(from, middle, color)
	w.Line(middle, to, color)
}
