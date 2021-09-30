package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", 43)

	if err != nil {
		log.Fatal(err)
	}

}
