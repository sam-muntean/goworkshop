package persistence

import "goworkshop/model"

func (store *GormDataStore) GetAuthors() ([]model.Author, error) {
	var authors []model.Author
	err := store.DBInstance.Find(&authors).Error
	return authors, err
}

func (store *GormDataStore) CreateAuthor(author *model.Author) (error)  {
	return store.DBInstance.Create(&author).Error
}
