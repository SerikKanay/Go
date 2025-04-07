package main

import (
	"github.com/gin-gonic/gin"
	"rest-api/config"
	"rest-api/controller"
	"rest-api/models"
	"rest-api/repository"
	"rest-api/service"
)

func main() {
	r := gin.Default()
	db := config.DbConnect()
	db.Table("books").AutoMigrate(&models.Book{})

	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookController := controller.NewBook(bookService)

	r.GET("/book", bookController.GetAllBook)
	r.POST("/book", bookController.CreateBook)
	r.DELETE("/book/:id", bookController.Delete)
	r.PUT("/book/:id", bookController.UpdateBook)
	r.Run()
}
