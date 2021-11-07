package control

import (
	"fmt"

	"github.com/codeation/impress"

	"github.com/codeation/lineation/mindmap"
)

func (c *Control) do(event impress.Eventer) {
	switch {
	case event == impress.KeySave:
		mm := mindmap.NewMindMap(c.mm.Filename(), c.view.GetNodes())
		if err := mm.Save(); err != nil {
			fmt.Printf("save: %v\n", err)
		}
		c.view.Modified(false)

	case event.Type() == impress.ConfigureEventType:
		configureEvent := event.(impress.ConfigureEvent)
		c.view.ConfigureSize(configureEvent.Size)

	case event == impress.KeyDown:
		c.view.KeyDown()
	case event == impress.KeyUp:
		c.view.KeyUp()
	case event.Type() == impress.ButtonEventType:
		buttonEvent := event.(impress.ButtonEvent)
		if buttonEvent.Action == impress.ButtonActionRelease && buttonEvent.Button == impress.ButtonLeft {
			c.view.Click(buttonEvent.Point)
		}

	case event == impress.KeyTab:
		c.view.AddChildNode()
		c.view.Modified(true)
	case event == impress.KeyEnter:
		c.view.AddNextNode()
		c.view.Modified(true)
	case event == impress.KeyDelete:
		c.view.DeleteNode()
		c.view.Modified(true)

	case event == impress.KeyBackSpace:
		c.view.RemoveLastChar()
		c.view.Modified(true)
	case event == impress.KeyLeft:
		c.view.KeyLeft()
	case event == impress.KeyRight:
		c.view.KeyRight()
	case event.Type() == impress.KeyboardEventType:
		keyboardEvent := event.(impress.KeyboardEvent)
		if keyboardEvent.IsGraphic() {
			c.view.InsertChar(keyboardEvent.Rune)
			c.view.Modified(true)
		}

	case event.Type() == impress.MotionEventType:
	}
}

func (c *Control) background() {
	c.view.ReDraw()
}
