package control

import (
	"github.com/codeation/impress/event"
)

func (c *Control) dragDrop(initAction event.Button) {
	state, ok := c.view.Catch(initAction.Point)
	if !ok {
		return
	}

	c.background()
	c.app.Sync()

	for {
		action := <-c.app.Chan()
		if action.Type() == event.MotionType {
			c.view.Drag(state, action.(event.Motion).Point)
		} else if action.Type() == event.ButtonType {
			buttonEvent := action.(event.Button)
			if state.IsDroppable() && buttonEvent.Action == event.ButtonActionRelease && buttonEvent.Button == event.ButtonLeft {
				c.view.Drop(state, buttonEvent.Point)
				c.view.Modified(true)
				return
			}
			break
		} else {
			break
		}

		if len(c.app.Chan()) == 0 {
			c.background()
			c.view.DrawDrag(state)
			c.app.Sync()
		}
	}

	c.view.DrawRemain(state)
}
