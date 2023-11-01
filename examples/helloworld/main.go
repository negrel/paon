package main

import (
	"context"
	"time"

	"github.com/negrel/paon"
	"github.com/negrel/paon/widgets"
)

func main() {
	app, err := paon.NewApp()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Start application.
	err = app.Start(ctx, widgets.NewSpan("Hello world"))
	if err != nil {
		panic(err)
	}
}
