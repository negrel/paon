package main

import (
	"context"
	"time"

	"github.com/negrel/paon"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/widgets/button"
	"github.com/negrel/paon/widgets/hbox"
	"github.com/negrel/paon/widgets/span"
	"github.com/negrel/paon/widgets/vbox"
)

func main() {
	app, err := paon.NewApp()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Start application.
	err = app.Start(ctx, vbox.New(
		hbox.New(
			span.New("English    "),
			span.New(" | "),
			span.New("French"),
		),
		span.New("-----------------------------------"),
		hbox.New(
			span.New("Hello world"),
			span.New(" | "),
			span.New("Bonjour tout le monde"),
		),
		button.New("Click to exit", func(event mouse.ClickEvent) {
			cancel()
		}),
	))
	if err != nil {
		panic(err)
	}
}
