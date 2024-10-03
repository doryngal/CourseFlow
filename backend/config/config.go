package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
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
}

func LoadConfig() (*Config, error) {
	var cfg Config

	viper.SetConfigFile("backend/config/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func InitConfig() *Config {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return cfg
}
