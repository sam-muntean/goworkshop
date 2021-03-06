package web

import (
	"net/http"
	"goworkshop/importer"
	"github.com/gorilla/mux"
	"goworkshop/model"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"goworkshop/persistence"
)

type Books []model.Book

var books Books = importer.ImportBooks()

//Demonstrates the basic functionality of private and public modifiers in GO
func Index(w http.ResponseWriter, r *http.Request) {
	helloWorkshop := struct {
		Message        string `json:"message"`
		privateMessage string `json:"privateMessage"`
		NoTagField     string `json:"-"`
	}{
		Message:        "Hello workshop!",
		privateMessage: "Message that does not appear in response :).",
		NoTagField:     "This message won't appear either",
	}
	WriteJson(w, helloWorkshop)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := persistence.Storage.GetBooks()
	if err != nil {
		panic(err)
	}
	WriteJson(w, books)
}

func GetBookByUUID(w http.ResponseWriter, r *http.Request) {
	var bookUUID = mux.Vars(r)["uuid"]
	book, err := books.get(bookUUID)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	} else {
		WriteJson(w, book)
	}
}

func DeleteBookByUUID(w http.ResponseWriter, r *http.Request) {
	var bookUUID = mux.Vars(r)["uuid"]
	err := books.delete(bookUUID)
	if err != nil {
		fmt.Fprintf(w, "Failed to delete book: %s", err)
	} else {
		WriteJson(w, books)
	}
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(bytes, &book); err != nil {
		panic(err)
	} else if err := persistence.Storage.CreateBook(&book); err != nil {
		panic(err)
	} else {
		WriteJson(w, book)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &book)
	if err != nil {
		fmt.Fprintf(w, "Failed to update book: %s", err)
		return
	}
	book, err = books.update(book)
	if err != nil {
		fmt.Fprintf(w, "Failed to update book: %s", err)
		return
	}
	WriteJson(w, book)
}

func (b *Books) get(bookUUID string) (model.Book, error) {
	err := fmt.Errorf("could not find book by uuid %s", bookUUID)
	for _, book := range *b {
		if book.UUID == bookUUID {
			return book, nil
		}
	}
	var book model.Book
	return book, err
}

func (b *Books) delete(bookUUID string) error {
	var err error = fmt.Errorf("could not find book by uuid %s", bookUUID)
	var updatedBooks Books
	for _, book := range *b {
		if book.UUID == bookUUID {
			err = nil
		} else {
			updatedBooks = append(updatedBooks, book)
		}
	}
	if err == nil {
		*b = updatedBooks
	}
	return err
}

func (b *Books) update(updatedBook model.Book) (model.Book, error) {
	var err error = fmt.Errorf("could not find book by uuid %s", updatedBook.UUID)
	var newBooks Books
	for _, book := range *b {
		if book.UUID == updatedBook.UUID {
			newBooks = append(newBooks, updatedBook)
			err = nil
		} else {
			newBooks = append(newBooks, book)
		}
	}
	if err == nil {
		*b = newBooks
	}
	return updatedBook, err
}