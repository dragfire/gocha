package chat

import (
	"fmt"
	"strconv"

	"github.com/jroimartin/gocui"
)

var y int = 0

func fromView(g *gocui.Gui, msg string, vy int) error {
	maxX, _ := g.Size()

	fname := "from" + strconv.Itoa(vy)
	v, err := g.SetView(fname, 1, vy, maxX-40, vy+2)

	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	v.FgColor = gocui.ColorCyan
	fmt.Fprintf(v, "%s", msg)

	return nil
}

func FromView(g *gocui.Gui, msg string, v *gocui.View) error {
	fromView(g, msg, y)
	y += 3
	return nil
}
