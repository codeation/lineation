package appclip

import (
	"github.com/codeation/impress"
	"github.com/codeation/impress/clipboard"
)

type Clipboard struct {
	app *impress.Application
}

func New(app *impress.Application) *Clipboard {
	return &Clipboard{
		app: app,
	}
}

func (c *Clipboard) Put(value string) {
	c.app.ClipboardPut(clipboard.Text(value))
}

func (c *Clipboard) Get() {
	c.app.ClipboardGet(clipboard.TextType)
}
