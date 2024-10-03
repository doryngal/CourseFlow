package repository

import (
	//"github.com/your_project/internal/models"
	"CourseFlow/backend/internal/models"
)

type CourseRepository interface {
	Save(course models.Course) error
	GetAll() ([]models.Course, error)
}

type PgCourseRepository struct {
	// добавьте необходимые поля, например, соединение с БД
}

func (r *PgCourseRepository) Save(course models.Course) error {
	// TODO logic of course saving in PostgreSQL
	return nil
}

func (r *PgCourseRepository) GetAll() ([]models.Course, error) {
	// TODO logic for retrieving all courses from PostgreSQL
	return []models.Course{}, nil
}
