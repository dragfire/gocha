package home

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("title", maxX/2-7, 0, maxX/2+7, 3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
		fmt.Fprintf(v, "\t\t\tGocha\nTerminal IRC")
	}

	if v, err := g.SetView("sidebar_one", 0, 3, maxX/2-5, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintf(v, "%50s", "Available Channels:\n\n")
		fmt.Fprintf(v, "Active Channels:\n")
		fmt.Fprintf(v, "Winterfell\n")
		fmt.Fprintf(v, "North\n")
		fmt.Fprintf(v, "South\n")
		fmt.Fprintf(v, "KingsLanding\n\n")
		fmt.Fprintf(v, "Inactive Channels:\n")
		fmt.Fprintf(v, "Castle Rock\n")
		fmt.Fprintf(v, "\n\n\n\nNote: Only public channels will be shown here.\n")
	}

	if v, err := g.SetView("sidebar_two", maxX/2+5, 3, maxX-1, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		fmt.Fprintf(v, "%50v", "Welcome to Gocha.\n\n")
		fmt.Fprintln(v, "\tAvailble commands:\n")
		fmt.Fprintln(v, "\t\t\t/create:channel_name/public [for private: /create:channel_name/private]\n")
		fmt.Fprintln(v, "\t\t\t/join:channel_name\n")
		fmt.Fprintln(v, "\t\t\t/switch:channel_name [switch to another channel]\n")
		fmt.Fprintln(v, "\t\t\t/leave [leave current channel and return to home]\n")
	}

	if v, err := g.SetView("cmdbar", 0, maxY-3, maxX, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Enter Command"
		v.Editable = true
		//fmt.Fprintf(v, "Enter Command:")
	}

	g.SetCurrentView("cmdbar")
	return nil
}
