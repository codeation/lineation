package main

// go build github.com/codeation/lineation/cmd
// go build --tags impresslink github.com/codeation/lineation/cmd

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/codeation/impress"
	"github.com/codeation/impress/driver"
	"github.com/codeation/tile/eventlink"

	"github.com/codeation/lineation/mapcontrol"
	"github.com/codeation/lineation/mapmodel"
	"github.com/codeation/lineation/mapview"
	"github.com/codeation/lineation/menuevent"
	"github.com/codeation/lineation/modified"
	"github.com/codeation/lineation/palette"
	"github.com/codeation/lineation/xmlfile"
)

func run(ctx context.Context, d driver.Driver) error {
	flag.Parse()
	if len(flag.Args()) != 1 {
		return fmt.Errorf("filename argument is missing")
	}
	filename := filepath.Clean(flag.Args()[0])

	pal := palette.NewPalette()
	title := fmt.Sprintf("lineation %s", filename)
	app := eventlink.MainApp(impress.MakeApplication(d, pal.DefaultAppRect(), title))
	defer app.Close()
	pal.FontLink(app.Application())
	defer pal.FontClose()
	menuevent.New(app.Application())

	mapRoot, err := xmlfile.Open(filename)
	if err != nil {
		return fmt.Errorf("xmlfile.Open: %w", err)
	}
	mapModel := mapmodel.New(mapRoot, filename)
	mapView := mapview.New(app, mapModel, pal)
	defer mapView.Destroy()
	modView := modified.NewView(app, pal)
	defer modView.Destroy()

	mapControl := mapcontrol.New(app, mapModel, mapView, modView)
	app.Run(ctx, mapControl)

	return nil
}

func root(d driver.Driver) {
	if err := run(context.Background(), d); err != nil {
		fmt.Println(err)
	}
}
