package main

import (
	"log"
	"os"
	"time"

	"github.com/negrel/ginger/v1"
	"github.com/negrel/ginger/v1/layout"
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
		// Center the children widget
		Root: &layout.Center{
			// Column that contain the multiple row
			Child: &layout.Column{
				// First row of hello world text
				Children: []widget.Widget{
					&layout.Row{
						Children: []widget.Widget{
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
					// Second row of hello world text
					&layout.Row{
						Children: []widget.Widget{

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
					// Third row of hello world text
					&layout.Row{
						Children: []widget.Widget{
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
		},
	})

	time.Sleep(time.Second * 5)

	// Stop the app
	app.Stop()
}
