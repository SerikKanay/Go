package service

import (
	"rest-api/models"
	"rest-api/repository"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	Delete(id int) error
}
type UserService struct {
	repo UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{repo: userRepo}
}

func (u UserService) GetAll() ([]models.User, error) {
	return u.repo.FindAll()
}
func (u UserService) Delete(id int) error {
	return u.repo.Delete(id)
}
