package mapcontrol

import (
	"context"
	"fmt"
	"image"

	"github.com/codeation/impress/clipboard"
	"github.com/codeation/impress/event"
	"github.com/codeation/tile/eventlink"
	"github.com/codeation/tile/eventlink/ctxchan"

	"github.com/codeation/lineation/mapmodel"
	"github.com/codeation/lineation/mapview"
	"github.com/codeation/lineation/menuevent"
	"github.com/codeation/lineation/modified"
	"github.com/codeation/lineation/xmlfile"
)

var keypadEnter = event.Keyboard{Name: "KP_Enter"}

type Control struct {
	mapModel *mapmodel.MindMap
	mapView  *mapview.View
	modView  *modified.View
}

func New(app eventlink.AppFramer, mapModel *mapmodel.MindMap, mapView *mapview.View, modView *modified.View,
) *Control {
	return &Control{
		mapModel: mapModel,
		mapView:  mapView,
		modView:  modView,
	}
}

func (c *Control) Wait() {}

func (c *Control) Action(ctx context.Context, app eventlink.App) {
	maybeChanged := true
	for {
		if len(app.Chan()) == 0 && maybeChanged {
			c.mapView.Draw(app)
			c.modView.Draw()
			app.Application().Sync()
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
			} else {
				c.modView.Set(false)
			}

		case event.KeyUp:
			c.mapModel.Up()
		case event.KeyDown:
			c.mapModel.Down()
		case event.KeyTab, menuevent.NewChild:
			c.mapModel.NewChildNode()
			c.modView.Set(true)
		case event.KeyEnter, keypadEnter, menuevent.NewNext:
			c.mapModel.NewNextNode()
			c.modView.Set(true)
		case menuevent.Delete:
			c.mapModel.DeleteNode()
			c.modView.Set(true)
		case event.KeyDelete:
			if len(c.mapModel.Selected.Value.String()) == c.mapModel.Selected.Value.Cursor() {
				c.mapModel.DeleteNode()
				c.modView.Set(true)
			}

		case event.KeyCopy, menuevent.Copy:
			app.Application().ClipboardPut(clipboard.Text(c.mapModel.Selected.Value.String()))
		case event.KeyPaste, menuevent.Paste:
			app.Application().ClipboardGet(clipboard.TextType)

		default:
			switch ev := e.(type) {
			case event.Configure:
				c.mapView.Configure(ev.InnerSize)

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
				oldValue := c.mapModel.Selected.Value.String()
				c.mapModel.NodeControl(ctx, app, e, func(ctx context.Context, app eventlink.App, e event.Eventer) {
					anyEvent = false
				})
				if oldValue != c.mapModel.Selected.Value.String() {
					c.modView.Set(true)
				}
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
			app.Application().Sync()
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
				c.modView.Set(true)
			}
			return

		default:
			return
		}
	}
}
