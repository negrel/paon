package main

import (
	"io"
	"log"
	"os"

	"github.com/gdamore/tcell"
)

var logf io.Writer

func init() {
	logf, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logf)
}

func main() {
	scr, err := tcell.NewTerminfoScreen()

	if err != nil {
		log.Fatal(err)
	}

	err = scr.Init()

	if err != nil {
		log.Fatal(err)
	}

	scr.Clear()
	scr.SetContent(0, 0, '|', nil, tcell.StyleDefault.Foreground(tcell.ColorAqua))
	scr.SetContent(1, 0, '-', nil, tcell.StyleDefault.Foreground(tcell.ColorAqua))
	scr.Show()

	for {
		ev := scr.PollEvent()

		log.Printf("%+v", ev)

		switch ev := ev.(type) {
		case *tcell.EventMouse:
			x, y := ev.Position()
			log.Println("Mouse event : ", x, y)

		case *tcell.EventKey:
			key := ev.Key()
			log.Println("Key event : ", key)

			if key == tcell.KeyCtrlC {
				goto end
			}

		case *tcell.EventResize:
			w, h := ev.Size()
			log.Println("Resize event : ", w, h)

		case *tcell.EventInterrupt:
			log.Println("Interrupt event : ", ev.Data())

		case *tcell.EventError:
			log.Println("Error event : ", ev.Error())
		}
	}

end:
	scr.Fini()
	log.Println("Exit.")
	os.Exit(0)
}
