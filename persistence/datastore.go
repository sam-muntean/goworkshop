package persistence

import (
	"goworkshop/model"
	"github.com/jinzhu/gorm"
)

var Storage DataStore

type GormDataStore struct {
	DBInstance *gorm.DB
}

type DataStore interface{
	//books
	GetBooks() ([]model.Book, error)
	CreateBook(book *model.Book) error
	//authors
	GetAuthors() ([]model.Author, error)
	CreateAuthor(author *model.Author) error
}
