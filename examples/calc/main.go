package main

import (
	"context"

	"github.com/negrel/paon"
	"github.com/negrel/paon/widgets"
)

func main() {
	app, err := paon.NewApp()
	if err != nil {
		panic(err)
	}

	err = app.Start(context.Background(), widgets.NewRoot(NewCalc()))
	if err != nil {
		panic(err)
	}
}
