package drawutil

import (
	"image"
	"image/color"

	"github.com/codeation/impress"
)

func DrawRectEdge(w *impress.Window, rect image.Rectangle, color color.Color) {
	pointX0Y0 := image.Pt(rect.Min.X, rect.Min.Y)
	pointX1Y0 := image.Pt(rect.Max.X, rect.Min.Y)
	pointX0Y1 := image.Pt(rect.Min.X, rect.Max.Y)
	pointX1Y1 := image.Pt(rect.Max.X, rect.Max.Y)

	w.Line(pointX0Y0, pointX1Y0, color)
	w.Line(pointX1Y0, pointX1Y1, color)
	w.Line(pointX0Y1, pointX1Y1, color)
	w.Line(pointX0Y0, pointX0Y1, color)
}

func DrawLine3Elem(w *impress.Window, offset image.Point, from image.Point, to image.Point, color color.Color) {
	drawLine3Elem(w, from.Add(offset), to.Add(offset), color)
}

func drawLine3Elem(w *impress.Window, from image.Point, to image.Point, color color.Color) {
	if from.Y == to.Y {
		w.Line(from, to, color)
		return
	}
	width := to.X - from.X
	width1 := width / 2
	width2 := width - width1
	w.Line(from, from.Add(image.Pt(width1, 0)), color)
	w.Line(from.Add(image.Pt(width1, 0)), to.Add(image.Pt(-width2, 0)), color)
	w.Line(to.Add(image.Pt(-width2, 0)), to, color)
}

func DrawLine2Elem(w *impress.Window, offset image.Point, from image.Point, to image.Point, color color.Color) {
	drawLine2Elem(w, from.Add(offset), to.Add(offset), color)
}

func drawLine2Elem(w *impress.Window, from image.Point, to image.Point, color color.Color) {
	middle := image.Pt(from.X, to.Y)
	w.Line(from, middle, color)
	w.Line(middle, to, color)
}
