package chat

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

var y int = 0

func fromView(g *gocui.Gui, vy int) error {
	maxX, _ := g.Size()

	fname := "from" + string(vy)

	if v, err := g.SetView(fname, 1, vy, maxX-40, vy+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintf(v, "Hello there!")
	}
	return nil
}

func FromView(g *gocui.Gui, v *gocui.View) error {
	y += 3
	fromView(g, y)
	return nil
}
