package control

import (
	"fmt"

	"github.com/codeation/impress/event"

	"github.com/codeation/lineation/menu"
	"github.com/codeation/lineation/mindmap"
)

func (c *Control) do(action event.Eventer) {
	switch {
	case action == event.KeySave || action == menu.Save:
		mm := mindmap.NewMindMap(c.mm.Filename(), c.view.GetNodes())
		if err := mm.Save(); err != nil {
			fmt.Printf("save: %v\n", err)
		}
		c.view.Modified(false)

	case action.Type() == event.ConfigureType:
		configureEvent := action.(event.Configure)
		c.view.ConfigureSize(configureEvent.InnerSize)

	case action == event.KeyDown:
		c.view.KeyDown()
	case action == event.KeyUp:
		c.view.KeyUp()
	case action.Type() == event.ButtonType:
		buttonEvent := action.(event.Button)
		if buttonEvent.Action == event.ButtonActionPress && buttonEvent.Button == event.ButtonLeft {
			c.view.Click(buttonEvent.Point)
			c.dragDrop(buttonEvent)
		}

	case action == event.KeyTab || action == menu.NewChild:
		c.view.AddChildNode(c.app)
		c.view.Modified(true)
	case action == event.KeyEnter || action == menu.NewNext:
		c.view.AddNextNode(c.app)
		c.view.Modified(true)
	case action == event.KeyDelete || action == menu.Delete:
		c.view.DeleteNode()
		c.view.Modified(true)

	case action == event.KeyBackSpace:
		c.view.RemoveLastChar()
		c.view.Modified(true)
	case action == event.KeyLeft:
		c.view.KeyLeft()
	case action == event.KeyRight:
		c.view.KeyRight()
	case action.Type() == event.KeyboardType:
		keyboardEvent := action.(event.Keyboard)
		if keyboardEvent.IsGraphic() {
			c.view.InsertChar(keyboardEvent.Rune)
			c.view.Modified(true)
		}
	}
}

func (c *Control) background() {
	c.view.ReDraw(c.app)
}
