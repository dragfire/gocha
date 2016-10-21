package chat

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"strconv"
)

var y int = 0

func fromView(g *gocui.Gui, msg string, vy int) error {
	maxX, _ := g.Size()

	fname := "from" + strconv.Itoa(vy)

	if v, err := g.SetView(fname, 1, vy, maxX-40, vy+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.FgColor = gocui.ColorCyan
		fmt.Fprintf(v, msg)
	}
	return nil
}

func FromView(g *gocui.Gui, msg string, v *gocui.View) error {
	fromView(g, msg, y)
	y += 3
	return nil
}
