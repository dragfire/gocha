package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	fmt.Println("GOCHA")
	app := gocui.NewGui()

	if err := app.Init(); err != nil {
		log.Fatal(err)
	}

	defer app.Close()

	app.SetLayout(layout)

	if err := app.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := app.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("GOCHA", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Println(v, "Setting up GOCHA View")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
