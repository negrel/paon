package main

import (
	"context"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/debug"
	"runtime/pprof"
	"time"

	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon"
	"github.com/negrel/paon/layouts"
	pdkwidgets "github.com/negrel/paon/pdk/widgets"
	"github.com/negrel/paon/styles/property"
	"github.com/negrel/paon/widgets"
)

const iter = 512

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomRune() rune {
	const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return rune(charSet[rand.Intn(len(charSet))])
}

func randomColor() property.Color {
	return property.ColorFromHex(rand.Uint32())
}

func childCount(layout widgets.Layout) int {
	i := 0
	child := layout.FirstChild()
	for child != nil {
		i++
		child = child.NextSibling()
	}

	return i
}

var counter = 0
var done = false

func main() {
	go func() {
		log.Debug(http.ListenAndServe("localhost:6060", nil))
	}()

	profile, err := os.Create("cpu.profile")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(profile)
	defer pprof.StopCPUProfile()
	debug.SetGCPercent(-1)

	app, err := paon.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	vbox := layouts.NewVBox()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			<-ticker.C
			log.Debugf("%v text widgets", counter)
		}
	}()

	defer func() {
		log.Debugf("%v text widgets", counter)
	}()

	go func() {
		texts := make([]string, iter+1)
		colors := make([]property.Color, iter+1)
		for i := 0; i < iter; i++ {
			texts[i] = string(randomRune())
			colors[i] = randomColor()
		}

		for i := 0; i < iter; i++ {
			j := i
			_ = j
			app.DoChannel() <- func() {
				// err = vbox.InsertBefore(vbox.FirstChild(), layouts.NewHBox())
				err = vbox.AppendChild(layouts.NewHBox())
				if err != nil {
					log.Fatal(err)
				}

				for child := vbox.FirstChild(); child != nil; child = child.NextSibling() {
					box := child.(pdkwidgets.Layout)
					txt := widgets.NewText(texts[rand.Int()%iter])
					txt.Style().SetColor(property.ForegroundColor(), &colors[rand.Int()%iter])

					counter++

					// err = box.AppendChild(txt)
					err = box.InsertBefore(box.FirstChild(), txt)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}

		// txt := widgets.NewText("Hello world")
		// theme := txt.Theme()
		// z1 := property.NewInt(1)
		// theme.SetInt(property.ZIndex(), &z1)
		// hbox := vbox.FirstChild().(pdkwidgets.Layout)
		// hbox.InsertBefore(hbox.FirstChild(), txt)
		// log.Debug("HELLO WORLD prepended")

		time.Sleep(5 * time.Second)
		cancel()
	}()

	err = app.Start(ctx, vbox)
	if err != nil {
		log.Fatal(err)
	}
}
