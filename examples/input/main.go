package main

import (
	"context"

	"github.com/negrel/paon"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/widgets/button"
	"github.com/negrel/paon/widgets/hbox"
	"github.com/negrel/paon/widgets/input"
	"github.com/negrel/paon/widgets/vbox"
)

func main() {
	app, err := paon.NewApp()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	input := input.New("")

	// Start application.
	err = app.Start(ctx, vbox.New(
		hbox.New(
			input,
		),
		button.New("Click to exit", func(_ mouse.ClickEvent) {
			cancel()
		}),
	))
	if err != nil {
		panic(err)
	}
}
