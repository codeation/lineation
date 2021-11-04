package control

import (
	"github.com/codeation/impress"

	"github.com/codeation/lineation/draw"
	"github.com/codeation/lineation/mindmap"
)

type Control struct {
	eventChan <-chan impress.Eventer
	view      *draw.View
	mm        *mindmap.MindMap
	modified  bool
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
		event := <-c.eventChan
		if event == impress.DestroyEvent || event == impress.KeyExit {
			return
		}

		c.do(event)

		if len(c.eventChan) == 0 {
			c.background()
		}
	}
}
