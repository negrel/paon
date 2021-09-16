package main

import (
	"math/rand"
	"os"
	"runtime/debug"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon"
	"github.com/negrel/paon/layouts"
	pdkwidgets "github.com/negrel/paon/pdk/widgets"
	"github.com/negrel/paon/styles/property"
	"github.com/negrel/paon/styles/value"
	"github.com/negrel/paon/widgets"
)

const iter = 1024

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// var letters = []rune("\u4e16\u754c")

func randomRune() rune {
	return letters[rand.Intn(len(letters))]
}

func randomString() string {
	length := 1
	r := make([]rune, length)

	for i := 0; i < int(length); i++ {
		r[i] = randomRune()
	}

	return string(r)
}

func randomColor() property.Color {
	return property.NewColor(property.ForegroundColorID(), value.ColorFromHex(rand.Uint32()))
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

	var wg sync.WaitGroup
	wg.Add(1)
	vbox := layouts.NewVBox()

	go func() {
		wg.Wait()
		log.Debugf("TERMINE SIZE %+v", vbox.Box().MarginBox().Size())
		log.Debug("STOPPING APP")
		app.Stop()
		log.Debug("APP STOPPED")
		log.Debugf("%v text widgets", counter)
	}()

	go func() {
		texts := make([]string, iter+1)
		for i := 0; i < iter; i++ {
			texts[i] = randomString()
		}

		for i := 0; i < iter; i++ {
			app.DoChannel() <- func() {
				// err = vbox.InsertBefore(vbox.FirstChild(), layouts.NewHBox())
				err = vbox.AppendChild(layouts.NewHBox())
				if err != nil {
					log.Fatal(err)
				}

				for child := vbox.FirstChild(); child != nil; child = child.NextSibling() {
					box := child.(pdkwidgets.Layout)
					txt := widgets.NewText(texts[i])
					txt.Style().Set(randomColor())

					counter++

					// err = box.AppendChild(txt)
					err = box.InsertBefore(box.FirstChild(), txt)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
		time.Sleep(time.Second)
		wg.Done()
	}()

	err = app.Start(vbox)
	if err != nil {
		log.Fatal(err)
	}
}
