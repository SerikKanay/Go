package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
	"rest-api/service"
	"strconv"
)

func NewBook(service *service.BookService) *BookController {
	return &BookController{service: service}
}

type BookController struct {
	service *service.BookService
}

func (b *BookController) GetAllBook(c *gin.Context) {
	book, _ := b.service.GetAll()
	c.JSON(http.StatusOK, book)
}
func (b *BookController) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	book, err := b.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)

}
func (b *BookController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if err := b.service.DeleteBook(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "book deleted successful"})
}

func (b *BookController) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBook, err := b.service.Create(book.Title, book.Author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}
func (b BookController) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var newBook *models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request "})
		return
	}
	update, err := b.service.Update(id, newBook)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"book": update})
}
