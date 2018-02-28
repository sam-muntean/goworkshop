package main

import (
	"fmt"
	"projects/goworkshop/importer"
	"projects/goworkshop/web"
	"projects/goworkshop/model"
)

func main() {
	model.Authors = importer.ImportAuthors()
	fmt.Printf("Imported authors are: %s\n", model.Authors)
	model.Books = importer.ImportBooks()
	fmt.Printf("Imported books are: %s\n", model.Books)
	web.StartServer()
}
