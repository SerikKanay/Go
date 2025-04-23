package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/service"
	"strconv"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}
func (u UserController) GetUser(c *gin.Context) {
	user, err := u.service.GetAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not fount"})
		return
	}
	c.JSON(http.StatusOK, user)
}
func (u UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user id error"})
		return
	}
	if err := u.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "user deleted success"})

}
