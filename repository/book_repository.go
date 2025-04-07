package repository

import (
	"gorm.io/gorm"
	"rest-api/models"
)

type BookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{db: db}
}

func (b BookRepositoryImpl) Create(book *models.Book) error {
	return b.db.Create(book).Error
}
func (b BookRepositoryImpl) FindAll() ([]models.Book, error) {
	var book []models.Book
	err := b.db.Find(&book).Error
	return book, err
}
func (b BookRepositoryImpl) Delete(id int) error {
	return b.db.Delete(&models.Book{}, id).Error
}
func (b BookRepositoryImpl) Update(id int, newBook *models.Book) error {
	result := b.db.Model(&models.Book{}).Where("id=?", id).Omit("id,CreatedAt").Updates(newBook)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
