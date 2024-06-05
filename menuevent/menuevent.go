package menuevent

import (
	"github.com/codeation/impress"
	"github.com/codeation/impress/event"
)

var (
	Save     = event.NewMenu("Save")
	SaveAs   = event.NewMenu("SaveAs")
	Exit     = event.NewMenu("Exit")
	NewChild = event.NewMenu("NewChild")
	NewNext  = event.NewMenu("NewNext")
	Delete   = event.NewMenu("Delete")
)

func New(app *impress.Application) {
	file := app.NewMenu("File")
	file.NewItem("Save", Save)
	file.NewItem("Save As", SaveAs)
	file.NewItem("Exit", Exit)

	node := app.NewMenu("Node")
	node.NewItem("New child node", NewChild)
	node.NewItem("New next node", NewNext)
	node.NewItem("Delete node", Delete)
}
