package main

import (
	"context"
	"fmt"

	"github.com/negrel/paon"
	"github.com/negrel/paon/colors"
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/widgets"
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
	counterSpan := span.New("0", span.WithStyle(widgets.Style{}.Border(styles.BorderSide{
		Size:      10,
		Style:     0,
		CellStyle: draw.CellStyle{},
	})))
	updateCounterSpan := func() {
		counterSpan.SetText(fmt.Sprintf("%v", counter))
	}

	// Button style.
	btnStyle := widgets.Style{}.Padding(0, 1).Border(styles.BorderSide{
		Size:  1,
		Style: styles.BorderHidden,
		CellStyle: draw.CellStyle{
			Background: colors.ColorGreen,
		},
	}).Margin(1)

	// Start application.
	err = app.Start(ctx, widgets.NewRoot(
		vbox.New(
			vbox.WithChildren(
				hbox.New(
					hbox.WithChildren(
						button.New(" - ", button.OnClick(func(_ mouse.ClickEvent) {
							counter--
							updateCounterSpan()
						}), button.WithStyle(btnStyle)),
						counterSpan,
						button.New(" + ", button.OnClick(func(_ mouse.ClickEvent) {
							counter++
							updateCounterSpan()
						}), button.WithStyle(btnStyle)),
					),
					hbox.WithStyle(
						widgets.Style{}.
							Border(styles.BorderSide{
								Size:  3,
								Style: styles.BorderHidden,
								CellStyle: draw.CellStyle{
									Background: colors.ColorRed,
								},
							}).
							Background(colors.ColorRebeccaPurple).
							Foreground(colors.ColorBlack),
					)),
				button.New("Click to exit", button.OnClick(func(_ mouse.ClickEvent) {
					cancel()
				})),
			),
		)))
	if err != nil {
		panic(err)
	}
}
