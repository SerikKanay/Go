package main

import (
	"github.com/gin-gonic/gin"
	"rest-api/config"
	"rest-api/controller"
	_ "rest-api/models"
	"rest-api/repository"
	"rest-api/service"
)

func main() {

	db := config.DbConnect()
	r := gin.Default()
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookController := controller.NewBook(bookService)

	r.GET("/book", bookController.GetAllBook)
	r.GET("book/:id", bookController.FindById)
	r.POST("/book", bookController.CreateBook)
	r.DELETE("/book/:id", bookController.Delete)
	r.PUT("/book/:id", bookController.UpdateBook)
	r.Run()
}
