package main

import (
	"context"

	"github.com/negrel/paon"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/widgets"
)

func main() {
	app, err := paon.NewApp()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	input := widgets.NewInput("")

	// Start application.
	err = app.Start(ctx, widgets.NewVBox(
		widgets.NewHBox(
			input,
		),
		widgets.NewButton("Click to exit", func(_ events.Click) {
			cancel()
		}),
	))
	if err != nil {
		panic(err)
	}
}
