package service

import (
	"os"
	"time"
	"todo/internal/models"
	"todo/internal/repository"
	"todo/internal/requests"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Ar *repository.AuthRepository
}

func NewAuthService(ar *repository.AuthRepository) *AuthService {
	return &AuthService{Ar: ar}
}

func (as *AuthService) Login(lr requests.LoginRequest) (string, error) {
	user, err := as.Ar.GetByEmail(lr.Email)
	if err != nil {
		return "", err
	}

	err = as.checkPasswordHash(lr.Password, user.Password)

	if err != nil {
		return "", err
	}

	token, err := as.GenerateToken(user.ID, user.Email)

	return token, err
}

func (as *AuthService) Register(user requests.RegisterRequest) (string, error) {
	hashPassword, err := as.hashPassword(user.Password)
	if err != nil {
		return "", err
	}

	nu := models.User{
		Name:       user.Name,
		Lastname:   user.Lastname,
		Patronymic: user.Patronymic,
		Email:      user.Email,
		Password:   hashPassword,
		Age:        user.Age,
	}

	err = as.Ar.Register(nu)
	if err != nil {
		return "", err
	}

	token, err := as.GenerateToken(nu.ID, nu.Email)

	return token, err
}

func (as *AuthService) GenerateToken(id int64, email string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func (as *AuthService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (as *AuthService) checkPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
