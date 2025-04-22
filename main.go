package main

import (
	"github.com/gin-gonic/gin"
	"rest-api/config"
	"rest-api/config/auth"
	"rest-api/controller"
	"rest-api/middleware"
	_ "rest-api/models"
	"rest-api/repository"
	"rest-api/service"
)

func main() {

	config.DbConnect()
	r := gin.Default()
	bookRepo := repository.NewBookRepository(config.DB)
	bookService := service.NewBookService(bookRepo)
	bookController := controller.NewBook(bookService)
	r.Use(middleware.LoggingMiddleware())
	r.POST("/login", auth.Login)
	r.POST("/register", auth.Register)
	r.GET("/book", middleware.AuthMiddleware(), bookController.GetAllBook)
	r.GET("/book/:id", bookController.FindById)
	r.POST("/book", bookController.CreateBook)
	r.DELETE("/book/:id", bookController.Delete)
	r.PUT("/book/:id", bookController.UpdateBook)

	r.Run()
}
