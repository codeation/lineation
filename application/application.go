package application

import (
	"github.com/codeation/impress"

	"github.com/codeation/lineation/control"
	"github.com/codeation/lineation/draw"
	"github.com/codeation/lineation/mindmap"
	"github.com/codeation/lineation/palette"
)

type Application struct {
	guiApplication *impress.Application
	guiWindow      *impress.Window
	control        *control.Control
}

func NewApplication(mm *mindmap.MindMap) *Application {
	pal := palette.NewPalette()
	guiApp := impress.NewApplication(pal.DefaultAppRect(), "lineation "+mm.Filename())
	w := guiApp.NewWindow(pal.DefaultAppRect(), pal.Color(palette.DefaultBackground))

	root := draw.NewBox(mm.Root(), pal)
	v := draw.NewView(w, root)

	c := control.NewControl(guiApp, v, mm)

	return &Application{
		guiApplication: guiApp,
		guiWindow:      w,
		control:        c,
	}
}

func (a *Application) Run() {
	a.guiApplication.Start(a.control.Loop)
	a.guiApplication.Wait()
}

func (a *Application) Close() {
	a.control.Done()
	a.guiWindow.Drop()
	a.guiApplication.Close()
}
