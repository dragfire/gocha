package chat

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("as", 1, 1, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Second Layout")
	}
	return nil
}

func Init(g *gocui.Gui, v *gocui.View) error {
	g.SetLayout(layout)
	return nil
}
