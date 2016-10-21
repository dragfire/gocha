package layouts

import (
	"fmt"
	"github.com/dragfire/gocha/layouts/chat"
	"github.com/dragfire/gocha/socket"
	"github.com/jroimartin/gocui"
	"golang.org/x/net/websocket"
	"strings"
)

const (
	NEW_CHANNEL          = "/create:"
	JOIN_CHANNEL         = "/join:"
	LEAVE_CHANNEL        = "/leave:"
	NEW_PRIVATE_CHANNEL  = "/private:create:"
	JOIN_PRIVATE_CHANNEL = "/private:join:"
	START                = "/start" // choose a default channel
)

func GetCmd(g *gocui.Gui, v *gocui.View) error {
	var cmd string

	cmd = v.ViewBuffer()

	fmt.Println("Get CMD")

	maxX, maxY := g.Size()

	if v, err := g.SetView("msg", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintf(v, "%s", cmd)
		if err := g.SetCurrentView("msg"); err != nil {
			return err
		}
	}
	return nil
}

func processCmd(g *gocui.Gui, v *gocui.View) error {
	cmd := v.ViewBuffer()
	if strings.Contains(cmd, START) {
		g.SetLayout(chat.Layout)
	}

	fmt.Println("Process CMD")
	
	return nil
}

func chatBoxHandlers(g *gocui.Gui, v *gocui.View) error {
	text := v.ViewBuffer()
	websocket.Message.Send(socket.Conn, &text)

	v.Clear()
	//fmt.Printf(text, strings.Contains(text, "to:"))
	if strings.Contains(text, "to") {
		chat.ToView(g, text, v)
	} else {
		chat.FromView(g, text, v)
	}
	return nil
}

func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
