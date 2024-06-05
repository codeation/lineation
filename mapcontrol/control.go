package mapcontrol

import (
	"context"
	"fmt"
	"image"

	"github.com/codeation/impress/event"
	"github.com/codeation/tile/eventlink"
	"github.com/codeation/tile/eventlink/ctxchan"

	"github.com/codeation/lineation/mapmodel"
	"github.com/codeation/lineation/mapview"
	"github.com/codeation/lineation/menuevent"
	"github.com/codeation/lineation/xmlfile"
)

type Control struct {
	mapModel *mapmodel.MindMap
	mapView  *mapview.View
}

func New(app eventlink.AppFramer, mapModel *mapmodel.MindMap, mapView *mapview.View) *Control {
	return &Control{
		mapModel: mapModel,
		mapView:  mapView,
	}
}

func (c *Control) Wait() {
}

func (c *Control) Action(ctx context.Context, app eventlink.App) {
	maybeChanged := true
	for {
		if len(app.Chan()) == 0 && maybeChanged {
			c.mapView.Draw(app)
			app.Sync()
			maybeChanged = false
		}

		e, ok := ctxchan.Get(ctx, app.Chan())
		if !ok {
			return
		}

		anyEvent := true
		switch e {
		case event.DestroyEvent, event.KeyExit, menuevent.Exit:
			app.Cancel()
			return

		case event.KeySave, menuevent.Save:
			data := c.mapModel.Export()
			if err := xmlfile.Save(data, c.mapModel.Filename); err != nil {
				fmt.Println(err)
			}

		case event.KeyUp:
			c.mapModel.Up()
		case event.KeyDown:
			c.mapModel.Down()
		case event.KeyTab, menuevent.NewChild:
			c.mapModel.NewChildNode()
		case event.KeyEnter, menuevent.NewNext:
			c.mapModel.NewNextNode()
		case event.KeyDelete, menuevent.Delete:
			c.mapModel.DeleteNode()

		case event.KeyLeft:
			c.mapModel.Left()
		case event.KeyRight:
			c.mapModel.Right()
		case event.KeyBackSpace:
			c.mapModel.Backspace()

		default:
			switch ev := e.(type) {
			case event.Configure:
				c.mapView.Configure(ev.InnerSize)

			case event.Keyboard:
				if ev.IsGraphic() {
					c.mapModel.Insert(ev.Rune)
				}

			case event.Button:
				if ev.Action == event.ButtonActionPress && ev.Button == event.ButtonLeft {
					if node, ok := c.mapView.Select(ev.Point); ok {
						c.mapModel.Select(node)
						if node != c.mapModel.Root {
							c.dragAndDrop(ctx, app, node, ev.Point)
						}
					}
				}

			default:
				anyEvent = false
			}
		}

		maybeChanged = maybeChanged || anyEvent
	}
}

func (c *Control) dragAndDrop(ctx context.Context, app eventlink.App, node *mapmodel.Node, from image.Point) {
	defer c.mapView.DragRelease()

	for {
		if len(app.Chan()) == 0 {
			c.mapView.Draw(app)
			app.Sync()
		}

		e, ok := ctxchan.Get(ctx, app.Chan())
		if !ok {
			return
		}

		switch ev := e.(type) {
		case event.Motion:
			c.mapView.Drag(node, from, ev.Point)

		case event.Button:
			if ev.Action == event.ButtonActionRelease && ev.Button == event.ButtonLeft {
				c.mapView.Drop(from, ev.Point)
			}
			return

		default:
			return
		}
	}
}
