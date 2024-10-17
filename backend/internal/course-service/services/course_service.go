package service

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
	if course.Name == "" || course.Author == "" {
		return errors.New("course name and author are required")
	}

	if course.Duration <= 0 {
		return errors.New("course duration must be greater than 0")
	}

	course.CreatedAt = time.Now()
	course.UpdatedAt = time.Now()

	return s.Repo.Save(course)
}

func (s *CourseService) GetAllCourses() ([]models.Course, error) {
	return s.Repo.GetAll()
}
