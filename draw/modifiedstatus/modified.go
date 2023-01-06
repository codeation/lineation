package modifiedstatus

import (
	"image"
	"image/color"

	"github.com/codeation/impress"

	"github.com/codeation/lineation/palette"
)

type ModifiedStatus struct {
	w          *impress.Window
	rect       image.Rectangle
	foreground color.Color
	isModified bool
}

func NewModifiedStatus(app *impress.Application, pal *palette.Palette) *ModifiedStatus {
	rect := image.Rectangle{Min: image.Pt(0, 0), Max: image.Pt(pal.HorizontalBoxAlign()/2, pal.VerticalBoxAlign()/4)}
	w := app.NewWindow(rect, pal.Color(palette.DefaultBackground))
	w.Show()
	return &ModifiedStatus{
		w:          w,
		rect:       rect,
		foreground: pal.Color(palette.ActiveEdge),
	}
}

func (ms *ModifiedStatus) Modified(ok bool) {
	if ms.isModified == ok {
		return
	}
	ms.isModified = ok
	ms.w.Clear()
	if ms.isModified {
		ms.w.Fill(ms.rect, ms.foreground)
	}
	ms.w.Show()
}
