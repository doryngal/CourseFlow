package repository

import (
	"database/sql"
	"github.com/doryngal/CourseFlow/backend/internal/course-service/models"
	_ "github.com/lib/pq"
	"log"
)

type CourseRepository interface {
	Save(course models.Course) error
	GetAll() ([]models.Course, error)
}

type PgCourseRepository struct {
	DB *sql.DB
}

func NewPgCourseRepository(db *sql.DB) *PgCourseRepository {
	return &PgCourseRepository{DB: db}
}

func (r *PgCourseRepository) Save(course models.Course) error {
	query := `
		INSERT INTO courses (id, title, description, instructor_id, created_at, updated_at, modules)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.DB.Exec(query, course.ID, course.Title, course.Description, course.InstructorID, course.CreatedAt, course.UpdatedAt, course.ModulesID)
	if err != nil {
		log.Printf("Error while inserting course: %v\n", err)
		return err
	}
	return nil
}

func (r *PgCourseRepository) GetAll() ([]models.Course, error) {
	query := `
		SELECT id, title, description, instructor_id, created_at, updated_at, modules 
		FROM courses`
	rows, err := r.DB.Query(query)
	if err != nil {
		log.Printf("Error while retrieving courses: %v", err)
		return nil, err
	}
	defer rows.Close()

	var courses []models.Course
	for rows.Next() {
		var course models.Course
		err := rows.Scan(&course.ID, &course.Title, &course.Description,
			&course.InstructorID, &course.CreatedAt, &course.UpdatedAt, &course.ModulesID)
		if err != nil {
			log.Printf("Error while scanning course: %v", err)
			return nil, err
		}
		courses = append(courses, course)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows error: %v", err)
		return nil, err
	}
	return courses, nil
}
