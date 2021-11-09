package control

import (
	"github.com/codeation/impress"
	"github.com/codeation/impress/event"

	"github.com/codeation/lineation/draw"
	"github.com/codeation/lineation/mindmap"
)

type Control struct {
	eventChan <-chan event.Eventer
	view      *draw.View
	mm        *mindmap.MindMap
}

func NewControl(app *impress.Application, v *draw.View, mm *mindmap.MindMap) *Control {
	c := &Control{
		eventChan: app.Chan(),
		view:      v,
		mm:        mm,
	}
	return c
}

func (c *Control) Done() {
}

func (c *Control) Loop() {
	for {
		action := <-c.eventChan
		if action == event.DestroyEvent || action == event.KeyExit {
			return
		}

		c.do(action)

		if len(c.eventChan) == 0 {
			c.background()
		}
	}
}
