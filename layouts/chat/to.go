package chat

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

var ty int = y

func toView(g *gocui.Gui, vy int) error {
	maxX, _ := g.Size()

	tname := "to" + string(vy)

	if v, err := g.SetView(tname, 40, vy, maxX-1, vy+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintf(v, "Hi! How are you?")
	}
	return nil
}

func ToView(g *gocui.Gui, v *gocui.View) error {
	ty += 3
	toView(g, ty)
	return nil
}
