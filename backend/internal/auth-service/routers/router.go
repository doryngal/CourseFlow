package routers

import (
	"github.com/doryngal/CourseFlow/backend/internal/auth-service/controllers"
	"github.com/doryngal/CourseFlow/backend/internal/auth-service/repository"
	"github.com/doryngal/CourseFlow/backend/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *repository.IAuthRepository) *gin.Engine {
	router := gin.Default()

	//router.Use(middleware.JWTMiddleware())

	authController := controllers.NewAuthController(db.UserRepo)

	//public routers
	router.POST("/login", authController.Login)
	router.POST("/register", authController.Register)

	//private routers
	privateGroup := router.Group("/auth")
	privateGroup.Use(middleware.JWTMiddleware())
	{
		privateGroup.GET("/profile", authController.GetProfile)
		privateGroup.POST("/logout", authController.Logout)
	}

	return router
}
