package persistence

import "goworkshop/model"

func (store *GormDataStore) GetBooks() ([]model.Book, error) {
	var books []model.Book
	err := store.DBInstance.Preload("Author").Find(&books).Error
	return books, err
}

func (store *GormDataStore) CreateBook(book *model.Book) (error)  {
	return store.DBInstance.Create(&book).Error
}