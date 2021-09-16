package main

import (
	"time"

	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon"
	"github.com/negrel/paon/layouts"
	"github.com/negrel/paon/widgets"
)

func main() {
	app, err := paon.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	txt := widgets.NewText("Hello, World!")
	vbox := layouts.NewVBox()
	vbox.AppendChild(txt)
	vbox.AppendChild(widgets.NewText("another text"))
	vbox.AppendChild(widgets.NewText("another text again"))

	go func() {
		time.Sleep(time.Second * 5)
		app.Stop()
	}()

	err = app.Start(vbox)
	if err != nil {
		log.Fatal(err)
	}
}
