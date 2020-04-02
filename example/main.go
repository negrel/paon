package main

import (
	"log"
	"time"

	"github.com/gdamore/tcell"
	"github.com/negrel/ginger/v1"
	"github.com/negrel/ginger/v1/widget"
)

func main() {
	app, err := ginger.New()

	if err != nil {
		log.Fatal(err)
	}

	app.Start(&ginger.Activity{
		Root: &widget.Column{
			Childrens: []widget.Widget{
				&widget.Row{
					Childrens: []widget.Widget{
						&widget.Text{
							Content: "- ",
						},
						&widget.Text{
							Content: "Hello, how are you ?",
							Colors:  tcell.StyleDefault.Background(tcell.ColorPaleGreen),
						},
					},
				},
				&widget.Row{
					Childrens: []widget.Widget{
						&widget.Text{
							Content: "- ",
						},
						&widget.Text{
							Content: "Fine, fuck you",
							Colors:  tcell.StyleDefault.Background(tcell.ColorTurquoise),
						},
					},
				},
			},
		},
	})

	time.Sleep(time.Second * 5)

	app.Stop()
}
