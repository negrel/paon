package main

import (
	"context"
	"fmt"

	"github.com/negrel/paon"
	"github.com/negrel/paon/events/click"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Counter value
	counter := 0

	// Counter span and a function to sync value and span.
	counterSpan := span.New(" 0 ")
	updateCounterSpan := func() {
		counterSpan.SetText(fmt.Sprintf(" %v ", counter))
	}

	// Start application.
	err = app.Start(ctx, vbox.New(
		hbox.New(
			button.New(" - |", func(_ click.Event) {
				counter--
				updateCounterSpan()
			}),
			counterSpan,
			button.New("| + ", func(_ click.Event) {
				counter++
				updateCounterSpan()
			}),
		),
		button.New("Click to exit", func(_ click.Event) {
			cancel()
		}),
	))
	if err != nil {
		panic(err)
	}
}
