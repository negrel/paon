package main

import (
	"context"
	"time"

	"github.com/negrel/paon"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/widgets"
)

func main() {
	// Create application.
	app, err := paon.NewApp()
	if err != nil {
		panic(err)
	}

	// Stop automatically application after 15sec.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Span widget containing current time.
	span := widgets.NewSpan(time.Now().Format(time.RFC3339))

	// Update span's text every second.
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			case now := <-ticker.C:
				app.DoChannel() <- func() {
					span.SetText(now.Format(time.RFC3339))
				}
			}
		}
	}()

	// Start application.
	err = app.Start(ctx, widgets.NewVBox(
		widgets.NewHBox(
			widgets.NewSpan("Current datetime: "),
			span,
		),
		widgets.NewButton("Click to exit", func(event events.Click) {
			cancel()
		}),
	))
	if err != nil {
		panic(err)
	}
}
