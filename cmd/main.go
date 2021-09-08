package main

import (
	"flag"
	"fmt"

	_ "github.com/codeation/impress/duo"

	"github.com/codeation/lineation/application"
	"github.com/codeation/lineation/mindmap"
)

func run() error {
	flag.Parse()
	if len(flag.Args()) != 1 {
		return fmt.Errorf("filename argument is missing")
	}

	mm, err := mindmap.Open(flag.Args()[0])
	if err != nil {
		return fmt.Errorf("NewMindMap: %w", err)
	}

	app := application.NewApplication(mm)
	defer app.Close()
	app.Run()
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
