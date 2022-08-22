package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	Name        string
	ListenPort  string
	MongoConfig MongoConfig
}

type MongoConfig struct {
	Host        string
	Port        string
	Username    string
	Password    string
	Database    string
	MaxPoolSize string
}

func LoadConfig() *App {
	// Set Default Value
	app := &App{
		Name:       "default-name",
		ListenPort: "3001",
		MongoConfig: MongoConfig{
			Host:        "localhost",
			Port:        "27017",
			Username:    "",
			Password:    "",
			Database:    "",
			MaxPoolSize: "1",
		},
	}
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error load env:", err.Error())
		return app
	}

	app.Name = os.Getenv("APP_NAME")
	app.ListenPort = os.Getenv("APP_PORT")
	app.MongoConfig.Host = os.Getenv("MONGO_HOST")
	app.MongoConfig.Port = os.Getenv("MONGO_PORT")
	app.MongoConfig.Username = os.Getenv("MONGO_USERNAME")
	app.MongoConfig.Password = os.Getenv("MONGO_PASSWORD")
	app.MongoConfig.Database = os.Getenv("MONGO_DATABASE")
	app.MongoConfig.MaxPoolSize = os.Getenv("MONGO_MAX_POOL_SIZE")

	return app
}
