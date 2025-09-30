package handlers

import (
	"todo/internal/requests"
	"todo/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	As *service.AuthService
}

func NewAuthHandler(as *service.AuthService) *AuthHandler {
	return &AuthHandler{As: as}
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var input requests.LoginRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := ah.As.Login(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func (ah *AuthHandler) Register(c *gin.Context) {
	var input requests.RegisterRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := ah.As.Register(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"token": token})
}

func (ah *AuthHandler) Me(c *gin.Context) {
	user, exist := c.Get("user")
	if !exist {
		c.JSON(401, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{"user": user})
}
