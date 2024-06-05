package mapview

import (
	"image"

	"github.com/codeation/lineation/palette"
)

func (v *View) arrowDraw(parentBorder image.Rectangle, nodeBorder image.Rectangle) {
	padding := v.pal.TextAlign()
	padding.Y += v.pal.DefaultFont().Height / 2
	if parentBorder.Min.X < nodeBorder.Min.X {
		// to right
		if parentBorder.Max.X < nodeBorder.Min.X {
			// next column
			from := image.Pt(parentBorder.Max.X, parentBorder.Min.Y+padding.Y)
			to := image.Pt(nodeBorder.Min.X, nodeBorder.Min.Y+padding.Y)
			v.arrowThree(from, to)
		} else {
			// same column
			from := image.Pt(parentBorder.Min.X+padding.X, parentBorder.Max.Y)
			to := image.Pt(nodeBorder.Min.X, nodeBorder.Min.Y+padding.Y)
			v.arrowTwo(from, to)
		}
	} else {
		// to left
		if parentBorder.Min.X > nodeBorder.Max.X {
			// next column
			from := image.Pt(parentBorder.Min.X, parentBorder.Min.Y+padding.Y)
			to := image.Pt(nodeBorder.Max.X, nodeBorder.Min.Y+padding.Y)
			v.arrowThree(from, to)
		} else {
			// same column
			from := image.Pt(parentBorder.Max.X-padding.X, parentBorder.Max.Y)
			to := image.Pt(nodeBorder.Max.X, nodeBorder.Min.Y+padding.Y)
			v.arrowTwo(from, to)
		}
	}
}

func (v *View) arrowThree(p1, p2 image.Point) {
	if p1.Y == p2.Y {
		v.w.Line(p1, p2, v.pal.Color(palette.DefaultLine))
		return
	}
	middleX := (p1.X + p2.X) / 2
	v.w.Line(p1, image.Pt(middleX, p1.Y), v.pal.Color(palette.DefaultLine))
	v.w.Line(image.Pt(middleX, p1.Y), image.Pt(middleX, p2.Y), v.pal.Color(palette.DefaultLine))
	v.w.Line(image.Pt(middleX, p2.Y), p2, v.pal.Color(palette.DefaultLine))
}

func (v *View) arrowTwo(p1, p2 image.Point) {
	v.w.Line(p1, image.Pt(p1.X, p2.Y), v.pal.Color(palette.DefaultLine))
	v.w.Line(image.Pt(p1.X, p2.Y), p2, v.pal.Color(palette.DefaultLine))
}
