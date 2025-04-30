package main

import (
	"github.com/gin-gonic/gin"
	"rest-api/config"
	"rest-api/config/auth"
	"rest-api/controller"
	"rest-api/middleware"
	"rest-api/repository"
	"rest-api/service"
)

func main() {
	config.DbConnect()

	bookRepo := repository.NewBookRepository(config.DB)
	bookService := service.NewBookService(bookRepo)
	bookCtrl := controller.NewBook(bookService)

	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userCtrl := controller.NewUserController(userService)

	r := gin.Default()
	r.Use(middleware.LoggingMiddleware(), middleware.CORSMiddleware())

	r.POST("/login", auth.Login)
	r.POST("/register", auth.Register)

	userRoutes := r.Group("/user", middleware.AuthMiddleware())
	{
		userRoutes.GET("/books", bookCtrl.GetAllBook)
		userRoutes.GET("/books/:id", bookCtrl.FindById)
	}

	adminRoutes := r.Group("/admin", middleware.AdminMiddleware())
	{
		adminRoutes.POST("/books", bookCtrl.CreateBook)
		adminRoutes.GET("/books", bookCtrl.GetAllBook)
		adminRoutes.PUT("/books/:id", bookCtrl.UpdateBook)
		adminRoutes.DELETE("/books/:id", bookCtrl.Delete)
		adminRoutes.GET("/users", userCtrl.GetUser)
		adminRoutes.DELETE("/users/:id", userCtrl.DeleteUser)
	}

	r.Run()
}
