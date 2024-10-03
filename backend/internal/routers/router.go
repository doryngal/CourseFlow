package routers

import (
	"CourseFlow/backend/internal/controllers"
	//"github.com/CourseFlow/internal/controllers"
	//"github.com/CourseFlow/internal/middleware"
	"CourseFlow/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Logger())

	courseController := controllers.NewCourseController()

	//routers
	router.POST("/courses", courseController.CreateCourse)
	router.GET("/courses", courseController.GetCourses)

	return router
}
