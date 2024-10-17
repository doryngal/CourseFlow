package main

import (
	"github.com/doryngal/CourseFlow/backend/config"
	"github.com/doryngal/CourseFlow/backend/internal/auth-service/repository"
	"github.com/doryngal/CourseFlow/backend/internal/auth-service/routers"
	"log"
)

func main() {
	cfg := config.InitCourseConfig()

	db, err := config.ConnectAuthDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	repo := repository.IAuthRepositories(db)

	router := routers.SetupRouter(repo)

	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start course-service: %v\n", err)
	}
}
