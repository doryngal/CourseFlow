package routers

import (
	"github.com/doryngal/CourseFlow/backend/internal/course-service/controllers"
	"github.com/doryngal/CourseFlow/backend/internal/course-service/repository"
	"github.com/doryngal/CourseFlow/backend/internal/shared/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupRouter(db *repository.RepositoryInitializer) *gin.Engine {
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

	router.Use(middleware.Logger())

	courseController := controllers.NewCourseController(db.CourseRepo)

	//routers
	router.POST("/courses", courseController.CreateCourse)
	router.GET("/courses", courseController.GetCourses)

	return router
}
