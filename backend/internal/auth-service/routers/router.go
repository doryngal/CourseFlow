package routers

import (
	"github.com/doryngal/CourseFlow/backend/internal/auth-service/controllers"
	"github.com/doryngal/CourseFlow/backend/internal/auth-service/repository"
	"github.com/doryngal/CourseFlow/backend/internal/shared/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupRouter(db *repository.IAuthRepository) *gin.Engine {
	router := gin.Default()

	// Настройка CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Замените на адрес фронтенда
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//router.Use(middleware.JWTMiddleware())

	authController := controllers.NewAuthController(db.UserRepo)

	//public routers
	router.POST("auth/login", authController.Login)
	router.POST("auth/register", authController.Register)

	//private routers
	privateGroup := router.Group("/auth")
	privateGroup.Use(middleware.JWTMiddleware())
	{
		privateGroup.GET("/profile", authController.GetProfile)
		privateGroup.POST("/logout", authController.Logout)
	}

	return router
}
