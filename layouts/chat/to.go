package chat

import (
	"fmt"
	"strconv"

	"github.com/jroimartin/gocui"
)

func toView(g *gocui.Gui, msg string, vy int) error {
	maxX, _ := g.Size()

	tname := "to" + strconv.Itoa(vy)
	// fmt.Println(tname)
	v, err := g.SetView(tname, 40, vy, maxX-1, vy+2)

	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	v.FgColor = gocui.ColorMagenta
	fmt.Fprintf(v, "%s", msg)

	return nil
}

func ToView(g *gocui.Gui, msg string, v *gocui.View) error {
	toView(g, msg, y)
	y += 3
	return nil
}
