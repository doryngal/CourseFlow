package main

import (
	"flag"
	"github.com/doryngal/CourseFlow/backend/config"
	"github.com/doryngal/CourseFlow/backend/internal/auth-service/repository"
	"github.com/doryngal/CourseFlow/backend/internal/auth-service/routers"
	"github.com/doryngal/CourseFlow/backend/internal/shared/migrations"
	"log"
)

func main() {
	migrate := flag.Bool("migrate", false, "Apply migrations")
	flag.Parse()

	cfg := config.InitAuthConfig()

	db, err := config.ConnectAuthDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	if !*migrate {
		err := migrations.ApplyMigrationsForAuth(db, cfg)
		if err != nil {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
		log.Println("Migrations applied successfully")
	}

	repo := repository.IAuthRepositories(db)

	router := routers.SetupRouter(repo)

	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start course-service: %v\n", err)
	}
}
