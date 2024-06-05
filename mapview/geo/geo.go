package geo

import (
	"image"
)

func RowSpan(size image.Point) image.Point { return image.Pt(size.X, 0) }

func RowSize(rects ...image.Point) image.Point {
	var output image.Point
	for _, rect := range rects {
		output.X += rect.X
		output.Y = max(output.Y, rect.Y)
	}
	return output
}

func ColSpan(size image.Point) image.Point { return image.Pt(0, size.Y) }

func ColSize(rects ...image.Point) image.Point {
	var output image.Point
	for _, rect := range rects {
		output.X = max(output.X, rect.X)
		output.Y += rect.Y
	}
	return output
}

func Mirror(pt image.Point) image.Point {
	return image.Pt(-pt.X, pt.Y)
}
