package main

import (
	"fmt"
	"io/ioutil"
	"goworkshop/domain"
	"encoding/json"
)

func main() {
	fileC, err := ioutil.ReadFile("main/authors.json")

	if err != nil {
		panic(err)
	}

	//fmt.Println(string(fileC))

	if err = json.Unmarshal(fileC, &domain.Authors); err != nil {
		panic(err)
	}

	fmt.Println("the authors are: ")
	fmt.Println(domain.Authors)

	fileC, err = ioutil.ReadFile("main/books.json")

	if err != nil {
		panic(err)
	}

	//fmt.Println(string(fileC))

	if err = json.Unmarshal(fileC, &domain.Books); err != nil {
		panic(err)
	}

	fmt.Println("the books are: ")
	fmt.Println(domain.Books)

	//serializedData, err := json.Marshal(books)
	//if err != nil{
	//	panic(err)
	//}
	//
	//fmt.Println("ser books are:")
	//fmt.Println(string(serializedData))
}
