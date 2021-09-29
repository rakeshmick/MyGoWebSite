package main

import (
	"io"
	"log"
	"os"
)

func main(args []string) {
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
	fileHtml := os.Create("index.html")

	BadExpr
	if err != nil {
		log.Fatal("")
	}
	io.Copy(file.html, fileHtml)
}
