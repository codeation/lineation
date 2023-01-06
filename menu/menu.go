package menu

import (
	"github.com/codeation/impress"
	"github.com/codeation/impress/event"
)

var (
	Save     = event.NewMenu("Save")
	Exit     = event.NewMenu("Exit")
	NewChild = event.NewMenu("NewChild")
	NewNext  = event.NewMenu("NewNext")
	Delete   = event.NewMenu("Delete")
)

func Init(app *impress.Application) {
	file := app.NewMenu("File")
	file.NewItem("Save", Save)
	file.NewItem("Exit", Exit)

	node := app.NewMenu("Node")
	node.NewItem("New child node", NewChild)
	node.NewItem("New next node", NewNext)
	node.NewItem("Delete node", Delete)
}
