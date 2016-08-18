package layout

import (
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
	app.SetLayout(Init)

	if err := keybindings(app); err != nil {
		log.Panicln(err)
	}

	if err := app.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
