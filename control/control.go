package control

import (
	"github.com/codeation/impress"

	"github.com/codeation/lineation/draw"
	"github.com/codeation/lineation/mindmap"
)

type eventer interface {
	Event() impress.Eventer
}

type Control struct {
	eventer interface{ Event() impress.Eventer }
	view    *draw.View
	mm      *mindmap.MindMap
}

func NewControl(eventer eventer, v *draw.View, mm *mindmap.MindMap) *Control {
	return &Control{
		eventer: eventer,
		view:    v,
		mm:      mm,
	}
}

func (c *Control) Done() {
}

func (c *Control) Loop() {
	for {
		event := c.eventer.Event()
		if event == impress.DoneEvent || event == impress.KeyExit {
			return
		}
		c.do(event)
	}
}
