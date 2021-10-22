package main

import (
	"image/jpeg"
	"log"
	"os"
	"time"

	"github.com/negrel/paon"
)

func main() {
	file, _ := os.Open("img.jpg")

	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}
	file.Close()

	app, err := paon.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		time.Sleep(time.Second * 5)
		app.Stop()
	}()

	err = app.Start(NewIMV(img))
	if err != nil {
		log.Fatal(err)
	}
}
