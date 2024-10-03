package controllers

import (
	"CourseFlow/backend/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	//"github.com/your_project/internal/models"
	//"github.com/your_project/internal/services"
)

type CourseController struct {
	Service *services.CourseService
}

func NewCourseController() *CourseController {
	service := services.NewCourseService(nil) //TODO repository transfer
	return &CourseController{Service: service}
}

func (cc *CourseController) CreateCourse(c *gin.Context) {
	//TODO course creation logic
	c.JSON(http.StatusOK, gin.H{"message": "Course created successfully"})
}

func (cc *CourseController) GetCourses(c *gin.Context) {
	//TODO course delivery logic
	c.JSON(http.StatusOK, gin.H{"message": "OKKK"})
}
