package layouts

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func GetCmd(g *gocui.Gui, v *gocui.View) error {
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

func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
