package config

import (
	"database/sql"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type CourseConfig struct {
	Server struct {
		Port string
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DbName   string
	}
	Migrations struct {
		Path   string
		DbName string
	}
	Redis struct {
		Addr     string
		Password string
		DB       int
	}
}

func loadCourseConfig() (*CourseConfig, error) {
	var cfg CourseConfig

	viper.SetConfigFile("backend/config/course-config.yml")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("CourseConfig file changed:", e.Name)
	})

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func InitCourseConfig() *CourseConfig {
	cfg, err := loadCourseConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return cfg
}

func ConnectDB(cfg *CourseConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return db, nil
}
