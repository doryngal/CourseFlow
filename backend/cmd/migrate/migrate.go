package main

import (
	"flag"
	"github.com/doryngal/CourseFlow/backend/internal/course-service/migrations"
	"log"

	"github.com/doryngal/CourseFlow/backend/config"
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
		err := migrations.ApplyMigrations(db, cfg)
		if err != nil {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
		log.Println("Migrations applied successfully")
		return
	}

	log.Println("No migration flag provided. Exiting.")
}
