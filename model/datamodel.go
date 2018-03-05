package model

import (
	"fmt"
)

//Book - The DTO used to access books
type Book struct {
	UUID        string `json:"uuid" gorm:"primary_key"`
	Title       string `json:"title"`
	NoPages     int    `json:"noPages"`
	ReleaseDate string `json:"releaseDate"`
	Author      Author `json:"author" gorm:"foreignkey:AuthorUUID"`
	AuthorUUID string `json:"-"`
}

//Author - The DTO used to access authors
type Author struct {
	UUID      string `json:"uuid" gorm:"primary_key"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
	Death     string `json:"death"`
}

func (author Author) String() string {
	return fmt.Sprintf("Author{UUID='%s', FirstName='%s', LastName='%s', Birthday='%s', Death='%s'}", author.UUID,
		author.FirstName, author.LastName, author.Birthday, author.Death)
}

func (book Book) String() string {
	return fmt.Sprintf("Book{UUID='%s', Title='%s', NoPages=%d, ReleaseDate='%s',Author=%s}", book.UUID, book.Title, book.NoPages, book.ReleaseDate, book.Author)
}

//Books - the list of available books
var Books []Book

// Authors - the list of available authors
var Authors []Author
