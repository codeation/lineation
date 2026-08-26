package saveas

import (
	"image"

	"github.com/codeation/tile/control/dialog"
	"github.com/codeation/tile/elem/field"
	"github.com/codeation/tile/view"
	"github.com/codeation/tile/view/align"
	"github.com/codeation/tile/view/box"
	"github.com/codeation/tile/view/column"
	"github.com/codeation/tile/view/empty"
	"github.com/codeation/tile/view/fieldview"
	"github.com/codeation/tile/view/fn"
	"github.com/codeation/tile/view/row"
	"github.com/codeation/tile/view/solid"
	"io.local/text/gen/article/elemview/chunkview"

	"github.com/codeation/lineation/palette"
)

func dialogView(
	d *dialog.Dialog, fields []*dialog.Field, buttons []*dialog.Button,
	header string, errorField *field.Field, pal *palette.Palette,
) view.Viewer {
	var lines []view.Viewer

	var blank view.Viewer = empty.New()
	blank = align.New(blank).Width(pal.DialogLeftColumnWidth)
	blank = box.New(blank).Padding(image.Pt(pal.DialogElemPadding().X, 0))

	var title view.Viewer = fieldview.New(
		fieldview.WithoutFocused(field.New(header)),
		pal.BoldFont(),
		fn.Const(pal.Color(palette.DefaultText))).
		Width(pal.DialogRightColumnWidth)
	title = box.New(title).Padding(image.Pt(0, pal.DialogElemPadding().Y))

	lines = append(lines, row.New(blank, title).Margin(pal.DialogElemMargin().X))

	for _, f := range fields {
		var prompt view.Viewer = fieldview.New(
			fieldview.WithoutFocused(field.New(f.Prompt)),
			pal.DefaultFont(),
			fn.Const(pal.Color(palette.DefaultText)))
		prompt = align.New(prompt).Width(pal.DialogLeftColumnWidth).Align(align.Right)
		prompt = box.New(prompt).Padding(pal.DialogElemPadding())

		selectedFn := func() bool { return d.Selected() == f }
		var value view.Viewer = fieldview.New(
			fieldview.WithFocused(f.Field, selectedFn),
			pal.DefaultFont(),
			fn.Const(pal.Color(palette.DefaultText))).
			Width(pal.DialogRightColumnWidth).
			WithCursor(solid.New(pal.CursorSize(), fn.Const(pal.Color(palette.CursorBlock))))
		value = box.New(value).
			Padding(pal.DialogElemPadding()).
			Foreground(fn.If(selectedFn, pal.Color(palette.ActiveEdge), pal.Color(palette.DefaultEdge))).
			Background(fn.Const(pal.Color(palette.DefaultBackground)))
		value = chunkview.New(value).Model(f)

		lines = append(lines, row.New(prompt, value).Margin(pal.DialogElemMargin().X))
	}

	var message view.Viewer = fieldview.New(fieldview.WithoutFocused(errorField),
		pal.DefaultFont(),
		fn.Const(pal.Color(palette.ActiveEdge))).
		Width(pal.DialogRightColumnWidth).
		MaxRows(1)

	lines = append(lines, row.New(blank, message).Margin(pal.DialogElemMargin().X))

	var buttonList []view.Viewer
	for _, b := range buttons {
		readyFn := func() bool {
			return (d.IsAnyButtonSelected() || d.IsLastFieldSelected()) && d.ButtonSelected() == b
		}
		selectedFn := func() bool {
			return d.IsAnyButtonSelected() && d.ButtonSelected() == b
		}
		var button view.Viewer = fieldview.New(
			fieldview.WithoutFocused(field.New(b.Prompt)),
			pal.DefaultFont(),
			fn.Const(pal.Color(palette.DefaultText)))
		button = align.New(button).Width(pal.DialogButtonWidth).Align(align.Center)
		button = box.New(button).
			Padding(pal.DialogElemPadding()).
			Foreground(fn.If(selectedFn, pal.Color(palette.DefaultLine), pal.Color(palette.DefaultEdge))).
			Background(fn.If(readyFn, pal.Color(palette.CursorBlock), pal.Color(palette.ActiveBoxBackground)))
		button = chunkview.New(button).Model(b)
		buttonList = append(buttonList, button)
	}
	buttonLine := row.New(buttonList...).Margin(pal.DialogElemMargin().X)

	lines = append(lines, row.New(blank, buttonLine).Margin(pal.DialogElemMargin().X))

	return box.New(column.New(lines...).Margin(pal.DialogElemMargin().Y)).
		Padding(pal.DialogPadding()).
		Foreground(fn.Const(pal.Color(palette.DefaultEdge)))
}
