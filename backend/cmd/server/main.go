package main

import (
	"github.com/doryngal/CourseFlow/backend/config"
	"github.com/doryngal/CourseFlow/backend/internal/repository"
	"github.com/doryngal/CourseFlow/backend/internal/routers"
	"log"
)

func main() {
	cfg := config.InitConfig()

	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	repo := repository.InitRepositories(db)

	router := routers.SetupRouter(repo)

	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
