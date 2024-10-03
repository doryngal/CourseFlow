package tests

import (
	"CourseFlow/backend/internal/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	//"github.com/your_project/internal/controllers"
)

func TestCreateCourse(t *testing.T) {
	router := gin.Default()
	courseController := controllers.NewCourseController()
	router.POST("/courses", courseController.CreateCourse)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/courses", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
