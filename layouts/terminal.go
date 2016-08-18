package layouts

import (
	"github.com/dragfire/gocha/layouts/home"
	"github.com/jroimartin/gocui"
	"log"
)

func Main() {
	app := gocui.NewGui()

	if err := app.Init(); err != nil {
		log.Fatal(err)
	}

	defer app.Close()

	app.Cursor = true
	app.SetLayout(home.Layout)

	if err := keybindings(app); err != nil {
		log.Panicln(err)
	}

	if err := app.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}