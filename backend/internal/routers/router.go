package routers

import (
	"github.com/doryngal/CourseFlow/backend/internal/controllers"
	"github.com/doryngal/CourseFlow/backend/internal/middleware"
	"github.com/doryngal/CourseFlow/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *repository.RepositoryInitializer) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Logger())

	courseController := controllers.NewCourseController(db.CourseRepo)

	//routers
	router.POST("/courses", courseController.CreateCourse)
	router.GET("/courses", courseController.GetCourses)

	return router
}
