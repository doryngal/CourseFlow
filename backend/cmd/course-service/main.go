package main

import (
	"flag"
	"github.com/doryngal/CourseFlow/backend/config"
	"github.com/doryngal/CourseFlow/backend/internal/course-service/repository"
	"github.com/doryngal/CourseFlow/backend/internal/course-service/routers"
	"github.com/doryngal/CourseFlow/backend/internal/shared/migrations"
	"log"
)

func main() {
	migrate := flag.Bool("migrate", false, "Apply migrations")
	flag.Parse()

	cfg := config.InitCourseConfig()

	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	if !*migrate {
		err := migrations.ApplyMigrationsForCourse(db, cfg)
		if err != nil {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
		log.Println("Migrations applied successfully")
	}

	//redisService := redisclient.NewRedisService(cfg)

	repo := repository.InitRepositories(db)
	router := routers.SetupRouter(repo)

	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start course-service: %v\n", err)
	}
}
