package main

import (
	"log"
	"time"

	"github.com/negrel/paon/pkg/paon"
	"github.com/negrel/paon/pkg/tui"
)

func main() {
	app := paon.New(tui.TextWidget("Hello world"))

	go func() {
		time.Sleep(time.Second * 3)
		app.Stop()
	}()

	log.Fatal(app.Start())
}
