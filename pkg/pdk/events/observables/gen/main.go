package main

import (
	"github.com/negrel/debuggo/pkg/log"
	"html/template"
	"os"
	"path"
	"strings"
)

func main() {
	funcMap := template.FuncMap{
		"title": strings.Title,
		"lower": strings.ToLower,
	}

	eventTemplate := template.Must(
		template.New("event-template").
			Funcs(funcMap).
			ParseFiles("observables.template"),
	)

	f, err := os.Create(path.Join("..", "observables.go"))
	if err != nil {
		log.Fatal(err)
	}

	//f = os.Stdout
	err = eventTemplate.ExecuteTemplate(f, "observables.template", types)
	if err != nil {
		log.Fatal(err)
	}
}
