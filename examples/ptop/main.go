package main

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/negrel/paon"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/widgets"
	"github.com/negrel/paon/widgets/button"
	"github.com/negrel/paon/widgets/hbox"
	"github.com/negrel/paon/widgets/span"
	"github.com/negrel/paon/widgets/vbox"
	"github.com/prometheus/procfs"
)

func main() {
	procFs, err := procfs.NewFS("/proc")
	if err != nil {
		panic(err)
	}

	// Create application.
	app, err := paon.NewApp()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	uptimeSpan := span.New("0")

	header :=
		hbox.New(
			hbox.WithChildren(
				span.New("uptime: "),
				uptimeSpan,
			),
		)

	procsVbox := vbox.New()
	body := vbox.New(
		vbox.WithChildren(
			hbox.New(
				hbox.WithChildren(
					span.New("PID |	"),
					span.New("USER |	"),
					span.New("CMD"),
				),
			),
			procsVbox,
		),
	)

	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				b, err := os.ReadFile("/proc/uptime")
				if err != nil {
					panic(err)
				}
				splitted := bytes.Split(b, []byte(" "))
				uptime, err := strconv.ParseFloat(string(splitted[0]), 64)
				if err != nil {
					panic(err)
				}

				now := time.Now()
				uptimeFormatted := now.Sub(now.Add(-time.Duration(math.Round(uptime)) * time.Second)).String()

				procs, err := procFs.AllProcs()
				if err != nil {
					panic(err)
				}

				app.DoChannel() <- func() {
					node := procsVbox.Node()
					for node.FirstChild() != nil {
						node.RemoveChild(node.FirstChild())
					}

					for _, proc := range procs {
						procStatus, err := proc.NewStatus()
						if err != nil {
							panic(err)
						}

						row := hbox.New(hbox.WithChildren(
							span.New(fmt.Sprint(proc.PID)),
							span.New("	"),
							span.New(procStatus.Name),
						))

						node.AppendChild(row.Node())
					}

					uptimeSpan.SetText(uptimeFormatted)
				}

			case <-ctx.Done():
				return
			}
		}
	}()

	// Start application.
	err = app.Start(ctx, widgets.NewRoot(vbox.New(
		vbox.WithChildren(
			button.New("Click to exit", button.OnClick(func(event mouse.ClickEvent) {
				cancel()
			})),
			header,
			body,
		),
	)))
	if err != nil {
		panic(err)
	}
}
