package main

import (
	"fmt"
	"io/ioutil"
	"goworkshop/domain"
	"encoding/json"
)

func main() {
	fileC, err := ioutil.ReadFile("main/books.json")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(fileC))
	var books []domain.BookDto

	if err = json.Unmarshal(fileC, &books); err != nil {
		panic(err)
	}

	fmt.Println("the books are: ")
	fmt.Println(books)

}
