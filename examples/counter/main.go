package main

import (
	"context"
	"fmt"

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

	// Counter value
	counter := 0

	// Counter span and a function to sync value and span.
	counterSpan := widgets.NewSpan(" 0 ")
	updateCounterSpan := func() {
		counterSpan.SetText(fmt.Sprintf(" %v ", counter))
	}

	// Start application.
	err = app.Start(ctx, widgets.NewVBox(
		widgets.NewHBox(
			widgets.NewButton(" - |", func(_ events.Click) {
				counter--
				updateCounterSpan()
			}),
			counterSpan,
			widgets.NewButton("| + ", func(_ events.Click) {
				counter++
				updateCounterSpan()
			}),
		),

		widgets.NewButton("Click to exit", func(_ events.Click) {
			cancel()
		}),
	))
	if err != nil {
		panic(err)
	}
}
