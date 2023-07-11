package control

import (
	"github.com/codeation/impress"
	"github.com/codeation/impress/event"

	"github.com/codeation/lineation/draw"
	"github.com/codeation/lineation/menu"
	"github.com/codeation/lineation/mindmap"
)

type Control struct {
	app  *impress.Application
	view *draw.View
	mm   *mindmap.MindMap
}

func NewControl(app *impress.Application, v *draw.View, mm *mindmap.MindMap) *Control {
	c := &Control{
		app:  app,
		view: v,
		mm:   mm,
	}
	return c
}

func (c *Control) Done() {
}

func (c *Control) Loop() {
	for {
		if len(c.app.Chan()) == 0 {
			c.app.Sync()
		}

		action := <-c.app.Chan()
		if action == event.DestroyEvent || action == event.KeyExit || action == menu.Exit {
			return
		}

		c.do(action)

		if len(c.app.Chan()) == 0 {
			c.reDraw()
		}
	}
}
