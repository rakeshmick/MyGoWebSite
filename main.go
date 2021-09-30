package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Fatal(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(nf, "one.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}
