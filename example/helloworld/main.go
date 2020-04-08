package main

import (
	"log"
	"os"
	"time"

	"github.com/negrel/ginger/v1"
	"github.com/negrel/ginger/v1/widget"
)

// For debugging purpose.
func init() {
	file, err := os.OpenFile("debug.log", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
}

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
							Content:    "Hello world",
							Background: 0xFFFFFF,
							Foreground: 0x0,
						},
						&widget.Text{
							Content:    "Hello world",
							Background: 0x0,
							Foreground: 0xFFFFFF,
						},
						&widget.Text{
							Content:    "Hello world",
							Background: 0xFFFFFF,
							Foreground: 0x0,
						},
					},
				},
				&widget.Row{
					Childrens: []widget.Widget{

						&widget.Text{
							Content:    "Hello world",
							Background: 0x0,
							Foreground: 0xFFFFFF,
						},
						&widget.Text{
							Content:    "Hello world",
							Background: 0xFFFFFF,
							Foreground: 0x0,
						},
						&widget.Text{
							Content:    "Hello world",
							Background: 0x0,
							Foreground: 0xFFFFFF,
						},
					},
				},
				&widget.Row{
					Childrens: []widget.Widget{
						&widget.Text{
							Content:    "Hello world",
							Background: 0xFFFFFF,
							Foreground: 0x0,
						},
						&widget.Text{
							Content:    "Hello world",
							Background: 0x0,
							Foreground: 0xFFFFFF,
						},
						&widget.Text{
							Content:    "Hello world",
							Background: 0xFFFFFF,
							Foreground: 0x0,
						},
					},
				},
			},
		},
	})

	time.Sleep(time.Second * 5)

	app.Stop()
}
