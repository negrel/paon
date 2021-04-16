package main

import (
	"html/template"
	"log"
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
		template.New("").
			Funcs(funcMap).
			ParseFiles("../observables.go.template"),
	)

	for _, observableType := range types {
		data := make(map[string]string)
		data["Type"] = observableType

		filename := "observables." + observableType + ".go"

		f, err := os.Create(path.Join("..", filename))
		if err != nil {
			log.Fatal(err)
		}

		err = eventTemplate.ExecuteTemplate(f, "observables.go.template", data)
		if err != nil {
			log.Fatal(err)
		}
	}
}
