package chat

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"strconv"
)

var y int = 0

func fromView(g *gocui.Gui, vy int) error {
	maxX, _ := g.Size()

	fname := "from" + strconv.Itoa(vy)

	if v, err := g.SetView(fname, 1, vy, maxX-40, vy+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintf(v, "Hello \033[36;7mcolors!\033[0m\n")
	}
	return nil
}

func FromView(g *gocui.Gui, v *gocui.View) error {
	fromView(g, y)
	y += 3
	return nil
}
