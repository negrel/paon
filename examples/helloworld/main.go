package main

import (
	"context"
	"math/rand"
	"os"
	"runtime/debug"
	"runtime/pprof"
	"time"

	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon"
	"github.com/negrel/paon/layouts"
	"github.com/negrel/paon/widgets"
)

const charSet = "abcdefghijklmnopqrstuvwxyz"

func randStr(size int) string {
	str := make([]byte, size)

	for i := 0; i < size; i++ {
		str[i] = charSet[rand.Int()%len(charSet)]
	}

	return string(str)
}

func main() {
	rand.Seed(time.Now().Unix())

	profile, err := os.Create("cpu.profile")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(profile)
	defer pprof.StopCPUProfile()

	app, err := paon.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	txt := widgets.NewText("Hello, World!")
	txt2 := widgets.NewText("Bye, World!")
	txt3 := widgets.NewText("Hello, World!")
	txt4 := widgets.NewText("Bye, World!")
	hbox := layouts.NewHBox()
	vbox := layouts.NewVBox()
	vbox2 := layouts.NewVBox()
	err = vbox.AppendChild(txt)
	if err != nil {
		log.Fatal(err)
	}
	err = vbox.AppendChild(txt2)
	if err != nil {
		log.Fatal(err)
	}

	err = vbox2.AppendChild(txt3)
	if err != nil {
		log.Fatal(err)
	}
	err = vbox2.AppendChild(txt4)
	if err != nil {
		log.Fatal(err)
	}

	err = hbox.AppendChild(vbox)
	if err != nil {
		log.Fatal(err)
	}
	err = hbox.AppendChild(vbox2)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	time.AfterFunc(5*time.Second, func() {
		debug.PrintStack()
		os.Exit(1)
	})

	err = app.Start(ctx, hbox)
	if err != nil {
		log.Fatal(err)
	}
}
