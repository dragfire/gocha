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
	
	//if err := g.SetKeybinding("cmdbar", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
	//g.SetLayout(chat.Layout)
	//return nil
	//}); err != nil {
	//return err
	//}
	
	if err := g.SetKeybinding("cmdbar", gocui.KeyEnter, gocui.ModNone, processCmd); err != nil {
		return err
	}
	
	if err := g.SetKeybinding("msg", gocui.KeyCtrlC, gocui.ModNone, Quit); err != nil {
		return err
	}
	
	if err := g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, chatBoxHandlers); err != nil {
		return err
	}
	
	return nil
}
