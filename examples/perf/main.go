package main

import (
	"context"

	"github.com/negrel/paon"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/widgets"
	"github.com/negrel/paon/widgets/button"
	"github.com/negrel/paon/widgets/hbox"
	"github.com/negrel/paon/widgets/span"
	"github.com/negrel/paon/widgets/vbox"
)

var charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	// Create application.
	app, err := paon.NewApp()
	if err != nil {
		panic(err)
	}

	root := vbox.New()

	for i := 0; i < 1000; i++ {
		line := hbox.New()
		for j := 0; j < 1000; j++ {
			idx := (i + j) % len(charSet)
			txt := charSet[idx : idx+1]
			line.Node().AppendChild(span.New(txt).Node())
		}
		root.Node().AppendChild(line.Node())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Prepend exit button.
	root.Node().InsertBefore(button.New(
		"Click to exit", button.OnClick(func(event mouse.ClickEvent) {
			cancel()
		}),
	).Node(), root.Node().FirstChild())

	app.Start(ctx, widgets.NewRoot(root))
}
