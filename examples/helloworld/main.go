package main

import (
	"context"
	"time"

	"github.com/negrel/paon"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/widgets"
)

func main() {
	app, err := paon.NewApp()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Start application.
	err = app.Start(ctx, widgets.NewVBox(
		widgets.NewHBox(
			widgets.NewSpan("English    "),
			widgets.NewSpan(" | "),
			widgets.NewSpan("French"),
		),
		widgets.NewSpan("-----------------------------------"),
		widgets.NewHBox(
			widgets.NewSpan("Hello world"),
			widgets.NewSpan(" | "),
			widgets.NewSpan("Bonjour tout le monde"),
		),
		widgets.NewButton("Click to exit", func(event events.Click) {
			cancel()
		}),
	))
	if err != nil {
		panic(err)
	}
}
