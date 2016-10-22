package chat

import (
	"fmt"
	"strconv"

	"github.com/jroimartin/gocui"
)

var vY int

func Bubble(g *gocui.Gui, msg string) error {
	maxX, _ := g.Size()

	fname := "bubble" + strconv.Itoa(vY)
	v, err := g.SetView(fname, 1, vY, maxX-1, vY+4)

	vY += 2

	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	v.Frame = false
	v.FgColor = gocui.ColorCyan
	v.BgColor = gocui.ColorDefault

	fmt.Fprintf(v, "%s", msg)
	return nil
}
