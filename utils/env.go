package utils

import (
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

type Config struct {
	Logs LogConfig
	DB   SQLiteConfig
	Port string
}

type LogConfig struct {
	Style string
	Level string
}

//type PostgresConfig struct {
//	Username string
//	Password string
//	URL      string
//	Port     string
//}

type SQLiteConfig struct {
	Path string
}

var cfg *Config

func LoadConfig() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{
		Port: os.Getenv("PORT"),
		Logs: LogConfig{
			Style: os.Getenv("LOG_STYLE"),
			Level: os.Getenv("LOG_LEVEL"),
		},
		DB: SQLiteConfig{
			Path: os.Getenv("DB_URL"),
		},
	}

	if cfg.DB.Path == "" {
		log.Println("DB_URL env variable not set, use memory db")
		cfg.DB.Path = ":memory:"
	}

	log.Println("Config loaded")
	return cfg, nil
}
