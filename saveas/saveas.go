package saveas

import (
	"image"

	"github.com/codeation/tile/control/dialog"
	"github.com/codeation/tile/elem/field"
	"github.com/codeation/tile/eventlink"
	"github.com/codeation/tile/view/area/fixedwindow"

	"github.com/codeation/lineation/palette"
)

func New(appFramer eventlink.AppFramer, fields []*dialog.Field, buttons []*dialog.Button,
	header string, errorField *field.Field, pal *palette.Palette,
) *dialog.Control {
	d := dialog.New(fields, buttons)
	d.Select(fields[len(fields)-1])

	viewer := dialogView(d, fields, buttons, header, errorField, pal)

	dialogSize := viewer.Size()
	rect := image.Rectangle{Max: dialogSize}.Add(appFramer.InnerRect().Size().Sub(dialogSize).Div(2))
	w := fixedwindow.New(appFramer, rect, pal.Color(palette.DefaultLine))

	return dialog.NewControl(w, d, viewer, rect)
}
