package menuevent

import (
	"github.com/codeation/impress"
	"github.com/codeation/impress/event"
)

var (
	Save     = event.NewMenu("Save")
	SaveAs   = event.NewMenu("SaveAs")
	Exit     = event.NewMenu("Exit")
	Copy     = event.NewMenu("Copy")
	Paste    = event.NewMenu("Paste")
	NewChild = event.NewMenu("NewChild")
	NewNext  = event.NewMenu("NewNext")
	Delete   = event.NewMenu("Delete")
)

func New(app *impress.Application) {
	file := app.NewMenu("File")
	file.NewItem("Save", Save)
	//file.NewItem("Save As", SaveAs) // TODO
	file.NewItem("Exit", Exit)

	edit := app.NewMenu("Edit")
	edit.NewItem("Copy", Copy)
	edit.NewItem("Paste", Paste)

	node := app.NewMenu("Node")
	node.NewItem("New child node", NewChild)
	node.NewItem("New next node", NewNext)
	node.NewItem("Delete node", Delete)
}
