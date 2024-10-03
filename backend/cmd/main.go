package main

import (
	"CourseFlow/backend/config"
	"CourseFlow/backend/internal/routers"
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

func main() {
	cfg := config.InitConfig()
	conn, err := pgx.Connect(context.Background(),
		"postgres://"+cfg.Database.User+":"+
			cfg.Database.Password+"@"+
			cfg.Database.Host+"/"+
			cfg.Database.DbName)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	router := routers.SetupRouter()

	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
