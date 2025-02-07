package utils

import (
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

type Config struct {
	Logs        LogConfig
	DB          SQLiteConfig
	Port        string
	ProjectPath string
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

func GetConfig() *Config {
	if cfg != nil {
		return cfg
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
		ProjectPath: os.Getenv("PROJECT_PATH"),
	}

	if cfg.DB.Path == "" {
		cfg.DB.Path = ":memory:"
		log.Println("DB_URL env variable not set, use memory db")
	}

	if cfg.ProjectPath == "" {
		cfg.ProjectPath = "/kitchen-sink-server"
		log.Println("PROJECT_PATH env variable not set, may cause some tests to FAIL")
	}

	if cfg.Port == "" {
		cfg.Port = ":8080"
		log.Println("PORT env variable not set, default :8080")
	}

	log.Println("Config loaded")

	return cfg
}
