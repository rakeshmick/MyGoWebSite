package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Println("executable " + os.Args[0])
	fmt.Println("first argument " + os.Args[1])
	name := os.Args[1]
	htmlData := `
	<html>
	<head>
	<title>
	</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fileHTML, err := os.Create("index.html")

	if err != nil {
		log.Fatal("file not created ")
	}
	defer fileHTML.Close()

	data := strings.NewReader(htmlData)
	io.Copy(fileHTML, data)
}
