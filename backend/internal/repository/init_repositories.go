package repository

import (
	"database/sql"
	"log"
)

type RepositoryInitializer struct {
	CourseRepo CourseRepository
}

func NewRepositoryInitializer(db *sql.DB) *RepositoryInitializer {
	return &RepositoryInitializer{
		CourseRepo: NewPgCourseRepository(db),
	}
}

func InitRepositories(db *sql.DB) *RepositoryInitializer {
	repos := NewRepositoryInitializer(db)
	log.Println("Repositories have been initialized")
	return repos
}
