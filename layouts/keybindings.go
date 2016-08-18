package layouts

import (
	"github.com/dragfire/gocha/layouts/chat"
	"github.com/jroimartin/gocui"
)

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, Quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyPgdn, gocui.ModNone, chat.Handler); err != nil {
		return err
	}

	if err := g.SetKeybinding("cmdbar", gocui.KeyEnter, gocui.ModNone, GetCmd); err != nil {
		return err
	}

	if err := g.SetKeybinding("msg", gocui.KeyCtrlC, gocui.ModNone, Quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, chat.FromView); err != nil {
		return err
	}

	return nil
}
