package routers

import (
	"github.com/doryngal/CourseFlow/backend/internal/course-service/controllers"
	"github.com/doryngal/CourseFlow/backend/internal/course-service/repository"
	"github.com/doryngal/CourseFlow/backend/internal/shared/middleware"
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
