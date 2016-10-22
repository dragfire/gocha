package chat

import (
	"fmt"
	"strconv"

	"github.com/jroimartin/gocui"
)

func Bubble(g *gocui.Gui, msg string, vY int) error {
	maxX, _ := g.Size()

	fname := "bubble" + strconv.Itoa(vY)
	v, err := g.SetView(fname, 1, vY, maxX-1, vY+2)

	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	v.Frame = false
	v.FgColor = gocui.ColorCyan

	fmt.Fprintf(v, "%s", msg)
	return nil
}
