package chat

import (
	"github.com/jroimartin/gocui"
)

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("chat", 1, 1, maxX-1, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Autoscroll = true
		v.Title = "Gocha"
	}

	if v, err := g.SetView("input", 1, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.BgColor = gocui.ColorDefault
		v.Autoscroll = false
		v.Editable = true
		v.Wrap = false
		v.Title = "Message"
		if err := g.SetCurrentView("input"); err != nil {
			return err
		}
	}

	return nil
}

func Handler(g *gocui.Gui, v *gocui.View) error {
	g.SetLayout(Layout)
	return nil
}
