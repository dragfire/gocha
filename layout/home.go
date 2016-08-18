package layout

import (
	"fmt"
	_ "log"

	"github.com/jroimartin/gocui"
)

func getCmd(g *gocui.Gui, v *gocui.View) error {
	var (
		cmd string
	)

	cmd = v.ViewBuffer()

	maxX, maxY := g.Size()

	if v, err := g.SetView("msg", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintf(v, "%s", cmd)
		if err := g.SetCurrentView("msg"); err != nil {
			return err
		}
	}
	return nil
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("cmdbar", gocui.KeyEnter, gocui.ModNone, getCmd); err != nil {
		return err
	}
	if err := g.SetKeybinding("msg", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
