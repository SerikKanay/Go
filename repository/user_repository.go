package repository

import (
	"gorm.io/gorm"
	"rest-api/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
func (u UserRepository) FindAll() ([]models.User, error) {
	var user []models.User
	err := u.db.Find(&user).Error
	return user, err
}
func (u UserRepository) Delete(id int) error {
	return u.db.Delete(&models.User{}, id).Error
}
