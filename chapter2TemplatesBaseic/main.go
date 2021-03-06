package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseGlob("tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatal(err)
	}
}
