package config

import (
	"belajar-mongodb/pkg/envar"
	"fmt"

	"github.com/joho/godotenv"
)

type App struct {
	Name        string
	AppHost     string
	ListenPort  int
	MongoConfig MongoConfig
}

type MongoConfig struct {
	Host        string
	Port        int
	Username    string
	Password    string
	Database    string
	MaxPoolSize int
}

// Loading konfigurasi services
func LoadConfig() *App {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error load env:", err.Error())
	}

	return &App{
		Name:       envar.GetEnv("APP_NAME", "default-app-name"),
		AppHost:    envar.GetEnv("APP_HOST", "localhost"),
		ListenPort: envar.GetEnv("APP_PORT", 5000),
		MongoConfig: MongoConfig{
			Host:        envar.GetEnv("MONGO_HOST", "localhost"),
			Port:        envar.GetEnv("MONGO_PORT", 0),
			Username:    envar.GetEnv("MONGO_USERNAME", ""),
			Password:    envar.GetEnv("MONGO_PASSWORD", ""),
			Database:    envar.GetEnv("MONGO_DATABASE", ""),
			MaxPoolSize: envar.GetEnv("MONGO_MAX_POOL_SIZE", 0),
		},
	}
}
