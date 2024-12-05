package services

import (
	"github.com/doryngal/CourseFlow/backend/internal/course-service/repository"
	"log"
)

type ServiceInitializer struct {
	courseService *CourseService
}

func NewServiceInitializer(repos *repository.RepositoryInitializer) *ServiceInitializer {
	return &ServiceInitializer{
		courseService: NewCourseService(repos.CourseRepo),
	}
}

func InitServices(repos *repository.RepositoryInitializer) *ServiceInitializer {
	services := NewServiceInitializer(repos)
	log.Println("Services have been initialized")
	return services
}
