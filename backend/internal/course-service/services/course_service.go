package services

import (
	"errors"
	"github.com/doryngal/CourseFlow/backend/internal/course-service/models"
	"github.com/doryngal/CourseFlow/backend/internal/course-service/repository"
	"time"
)

type CourseService struct {
	Repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) *CourseService {
	return &CourseService{
		Repo: repo,
	}
}

func (s *CourseService) CreateCourse(course models.Course) error {
	if course.Title == "" || course.InstructorID > 0 {
		return errors.New("course name and author are required")
	}

	course.CreatedAt = time.Now()
	course.UpdatedAt = time.Now()

	return s.Repo.Save(course)
}

func (s *CourseService) GetAllCourses() ([]models.Course, error) {
	return s.Repo.GetAll()
}
