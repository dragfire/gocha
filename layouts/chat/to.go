package chat

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"strconv"
)

func toView(g *gocui.Gui, vy int) error {
	maxX, _ := g.Size()

	tname := "to" + strconv.Itoa(vy)
	// fmt.Println(tname)
	if v, err := g.SetView(tname, 40, vy, maxX-1, vy+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintf(v, "\x1b[0;36mHow are you?")
	}
	return nil
}

func ToView(g *gocui.Gui, v *gocui.View) error {
	toView(g, y)
	y += 3
	return nil
}
