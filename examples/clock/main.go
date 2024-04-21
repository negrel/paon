package main

import (
	"context"
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/negrel/paon"
	"github.com/negrel/paon/colors"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/widgets"
	"github.com/negrel/paon/widgets/border"
	"github.com/negrel/paon/widgets/button"
	"github.com/negrel/paon/widgets/hbox"
	"github.com/negrel/paon/widgets/span"
	"github.com/negrel/paon/widgets/vbox"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// Create application.
	app, err := paon.NewApp()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Span widget containing current time.
	timeSpan := span.New(time.Now().Format(time.RFC3339), span.WithStyle(
		widgets.Style{}.Bold(true).Underline(true),
	))

	// Update span's text every second.
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-ctx.Done():
				return
			case now := <-ticker.C:
				app.DoChannel() <- func() {
					timeSpan.SetText(now.Format(time.RFC3339Nano))
				}
			}
		}
	}()

	// Start application.
	err = app.Start(ctx, widgets.NewRoot(
		vbox.New(
			vbox.WithStyle(widgets.Style{}.Background(colors.Color100)),
			vbox.WithChildren(
				hbox.New(
					hbox.WithStyle(widgets.Style{}.Border(border.RoundedBorder).Background(colors.Color98)),
					hbox.WithChildren(
						span.New("Current datetime: "),
						timeSpan,
					),
				),
				button.New("Click to exit", button.OnClick(func(event mouse.ClickEvent) {
					cancel()
				})),
			),
		),
	))
	if err != nil {
		panic(err)
	}
}
