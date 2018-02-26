package main

import (
	"fmt"
	"goworkshop/importer"
	"goworkshop/web"
	"goworkshop/model"
)

func main() {
	model.Authors = importer.ImportAuthors()
	fmt.Printf("Imported authors are: %s\n", model.Authors)
	model.Books = importer.ImportBooks()
	fmt.Printf("Imported books are: %s\n", model.Books)
	web.StartServer()
}
