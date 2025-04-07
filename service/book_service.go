package service

import "rest-api/models"

type BookRepository interface {
	Create(book *models.Book) error
	FindAll() ([]models.Book, error)
	Delete(id int) error
}
type BookService struct {
	repo BookRepository
}

func NewBookService(bookRepo BookRepository) *BookService {
	return &BookService{repo: bookRepo}
}
func (b *BookService) GetAll() ([]models.Book, error) {
	return b.repo.FindAll()
}
func (b *BookService) Create(title, author string) (*models.Book, error) {
	book := &models.Book{
		Title:  title,
		Author: author,
	}
	err := b.repo.Create(book)
	return book, err
}
func (b *BookService) DeleteBook(bookId int) error {
	return b.repo.Delete(bookId)
}
