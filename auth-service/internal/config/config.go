package config

import (
	"os"
)

const (
	EnvironmentDev  string = "dev"
	EnvironmentProd string = "prod"
)

type AppConfig struct {
	Environment             string
	MongodbDatabaseName     string
	MongodbConnectionString string
	RedisConnectionString   string
}

func LoadAppConfig() *AppConfig {
	appConfig := &AppConfig{
		Environment:             EnvironmentProd,
		MongodbConnectionString: "",
		RedisConnectionString:   "",
	}

	env := os.Getenv("ENVIRONMENT")

	if env == "dev" {
		appConfig.Environment = EnvironmentDev
	}

	appConfig.MongodbConnectionString = os.Getenv("MONGODB_CONNECTION_STRING")
	appConfig.MongodbDatabaseName = os.Getenv("MONGODB_DATABASE_NAME")
	appConfig.RedisConnectionString = os.Getenv("REDIS_CONNECTION_STRING")

	return appConfig
}
