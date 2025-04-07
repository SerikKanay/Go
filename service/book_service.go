package service

import "rest-api/models"

type BookRepository interface {
	Create(book *models.Book) error
	FindAll() ([]models.Book, error)
	Delete(id int) error
	Update(id int, newBook *models.Book) error
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
func (b *BookService) DeleteBook(id int) error {
	return b.repo.Delete(id)
}
func (b *BookService) Update(id int, newBook *models.Book) (*models.Book, error) {
	err := b.repo.Update(id, newBook)
	if err != nil {
		return nil, err
	}
	return newBook, nil
}
