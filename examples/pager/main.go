package main

import (
	"context"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/negrel/paon"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/widgets"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	if len(os.Args) == 1 {
		println("please specify a single file path")
		os.Exit(1)
	}

	fpath := os.Args[1]
	f, err := os.Open(fpath)
	if err != nil {
		println("failed to open file", fpath, err.Error())
		os.Exit(1)
	}

	// Read file.
	content, err := io.ReadAll(f)
	if err != nil {
		println("failed to read file", fpath, err.Error())
		os.Exit(1)
	}

	// Split lines.
	lines := strings.Split(string(content), "\n")

	app, err := paon.NewApp()
	if err != nil {
		panic(err)
	}

	// Application context.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Root widget.
	root := widgets.NewRoot(NewFrame(lines))

	// Close app on ctrl-c.
	root.AddEventListener(events.KeyListener(func(ev events.Event, data events.KeyEventData) {
		if data.Key.String() == "Ctrl-C" || data.Rune == 'q' {
			cancel()
		}
	}))

	// Start application.
	err = app.Start(ctx, root)
	if err != nil {
		panic(err)
	}
}
