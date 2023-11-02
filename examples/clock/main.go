package main

import (
	"context"
	"time"

	"github.com/negrel/paon"
	"github.com/negrel/paon/events/click"
	"github.com/negrel/paon/widgets/button"
	"github.com/negrel/paon/widgets/hbox"
	"github.com/negrel/paon/widgets/span"
	"github.com/negrel/paon/widgets/vbox"
)

func main() {
	// Create application.
	app, err := paon.NewApp()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Span widget containing current time.
	timeSpan := span.New(time.Now().Format(time.RFC3339))

	// Update span's text every second.
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			case now := <-ticker.C:
				app.DoChannel() <- func() {
					timeSpan.SetText(now.Format(time.RFC3339))
				}
			}
		}
	}()

	// Start application.
	err = app.Start(ctx, vbox.New(
		hbox.New(
			span.New("Current datetime: "),
			timeSpan,
		),
		button.New("Click to exit", func(event click.Event) {
			cancel()
		}),
	))
	if err != nil {
		panic(err)
	}
}
