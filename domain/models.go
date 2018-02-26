package domain

import (
	"fmt"
	//"golang.org/x/tools/go/gcimporter15/testdata"
)

type AuthorDto struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
	Death     string `json:"death"`
}

type BookDto struct {
	UUID        string `json:"uuid"`
	Title       string `json:"title"`
	NoPages     int `json:"noPages"`
	ReleaseDate string `json:"releaseDate"`
	Author      AuthorDto
}

func (b BookDto) String() string {
	return fmt.Sprintf("BookDtp{UUID=%s, Title=%s, NoPages=%d, ReleaseDate=%s, Author=%s}",
		b.UUID, b.Title, b.NoPages, b.ReleaseDate, b.Author)
}

func (a AuthorDto) String() string {
	return fmt.Sprintf("AuthorDto{UUID=%s, FirstName=%s, LastName=%s, Birthday=%s, Death=%s}",
		a.UUID, a.FirstName, a.LastName, a.Birthday, a.Death)
}

var Books []BookDto

var Authors []AuthorDto