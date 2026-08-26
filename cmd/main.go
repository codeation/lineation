package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/codeation/impress"
	"github.com/codeation/impress/duo/duodriver"
	"github.com/codeation/tile/eventlink"

	"github.com/codeation/lineation/mapcontrol"
	"github.com/codeation/lineation/mapmodel"
	"github.com/codeation/lineation/mapview"
	"github.com/codeation/lineation/menuevent"
	"github.com/codeation/lineation/modified"
	"github.com/codeation/lineation/palette"
	"github.com/codeation/lineation/xmlfile"
)

func run(ctx context.Context) error {
	var filename string
	mapRoot := &xmlfile.Node{Text: "New Topic"}

	flag.Parse()
	if len(flag.Args()) == 1 {
		filename = filepath.Clean(flag.Args()[0])
		xmlNode, err := xmlfile.Open(filename)
		if err != nil {
			return fmt.Errorf("xmlfile.Open: %w", err)
		}
		mapRoot = xmlNode
	}

	pal, err := palette.NewPalette()
	if err != nil {
		return fmt.Errorf("palette.NewPalette: %w", err)
	}

	d, err := duodriver.New()
	if err != nil {
		return fmt.Errorf("duodriver.New: %w", err)
	}

	title := fmt.Sprintf("lineation %s", filename)
	app := eventlink.MainApp(impress.MakeApplication(d, pal.DefaultAppRect(), title))
	defer app.Close()
	menuevent.New(app.Application())

	pal.FontsOpen(app.Application())
	defer pal.FontsClose()

	mapModel := mapmodel.New(mapRoot, filename)
	mapView := mapview.New(app, mapModel, pal)
	defer mapView.Destroy()
	modView := modified.NewView(app, pal)
	defer modView.Destroy()

	mapControl := mapcontrol.New(app, mapModel, mapView, modView, pal)
	app.Run(ctx, mapControl)

	return nil
}

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Println(err)
	}
}
