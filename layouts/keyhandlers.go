package layouts

import (
	"fmt"
	"strings"
	"time"

	"github.com/dragfire/gocha/layouts/chat"
	"github.com/dragfire/gocha/socket"
	"github.com/fatih/color"
	"github.com/jroimartin/gocui"
	"golang.org/x/net/websocket"
)

const (
	newChannel         = "/create:"
	joinChannel        = "/join:"
	leaveChannel       = "/leave:"
	newPrivateChannel  = "/private:create:"
	joinPrivateChannel = "/private:join:"
	start              = "/start" // choose a default channel
)

func GetCmd(g *gocui.Gui, v *gocui.View) error {
	var cmd string

	cmd = v.ViewBuffer()

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
	if strings.Contains(cmd, start) {
		g.SetLayout(chat.Layout)
	}
	return nil
}

func chatBoxHandlers(g *gocui.Gui, v *gocui.View) error {
	text := strings.TrimSpace(v.ViewBuffer())
	websocket.Message.Send(socket.Conn, &text)

	v.Clear()
	chatView, err := g.View("chat")

	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	TSColor := color.New(color.FgYellow).SprintFunc()

	timestamp := time.Now().Format("03:04")
	fmt.Fprintf(chatView, "-> [%s] * %s\n", TSColor(timestamp), color.CyanString(text))
	return nil
}

func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
