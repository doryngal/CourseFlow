package services

import (
	//"github.com/your_project/internal/models"
	//"github.com/your_project/internal/repository"
	"CourseFlow/backend/internal/models"
	"CourseFlow/backend/internal/repository"
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
	return s.Repo.Save(course)
}

func (s *CourseService) GetAllCourses() ([]models.Course, error) {
	return s.Repo.GetAll()
}
