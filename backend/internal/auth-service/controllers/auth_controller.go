package controllers

import (
	"github.com/doryngal/CourseFlow/backend/internal/auth-service/models"
	"github.com/doryngal/CourseFlow/backend/internal/auth-service/repository"
	"github.com/doryngal/CourseFlow/backend/internal/auth-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type AuthController struct {
	Service *services.AuthService
}

func NewAuthController(repo repository.UserRepository) *AuthController {
	service := services.NewAuthService(repo)
	return &AuthController{Service: service}
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	token, err := ac.Service.Login(loginData.Email, loginData.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ac *AuthController) Register(c *gin.Context) {
	var registerData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		Name     string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:        registerData.Email,
		Name:         registerData.Name,
		PasswordHash: registerData.Password,
		UpdatedAt:    time.Now(),
		CreatedAt:    time.Now(),
	}

	err := ac.Service.Register(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (ac *AuthController) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is missing"})
		return
	}

	tokenString := strings.Split(authHeader, "Bearer ")
	if len(tokenString) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is malformed"})
		return
	}

	err := ac.Service.Logout(tokenString[1])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}

func (ac *AuthController) GetProfile(c *gin.Context) {
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := ac.Service.GetProfile(userId.(int))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"profile": user})
}
