package controllers

import (
	"github.com/doryngal/CourseFlow/backend/internal/course-service/repository"
	"github.com/doryngal/CourseFlow/backend/internal/course-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CourseController struct {
	Service *services.CourseService
}

func NewCourseController(repo repository.CourseRepository) *CourseController {
	service := services.NewCourseService(repo)
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

func (cc *CourseController) GetCourseById(c *gin.Context) {
	//TODO logic
	c.JSON(http.StatusOK, gin.H{"message": "OKKK"})
}
