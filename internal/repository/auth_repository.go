package repository

import (
	"errors"
	"todo/internal/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthkRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := ar.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (ar *AuthRepository) Register(user models.User) error {
	err := ar.db.Create(&user).Error

	return err
}

func (ar *AuthRepository) GetByID(id int64) (models.User, error) {
	var user models.User
	err := ar.db.First(&user, id).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
