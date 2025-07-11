package config

import (
	"fmt"
	"log"
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
	SmtpHost                string
	SmtpPort                uint16
	SmtpUsername            string
	SmtpPassword            string
	SmtpFrom                string
}

func LoadAppConfig() *AppConfig {
	appConfig := &AppConfig{
		Environment:             EnvironmentProd,
		MongodbConnectionString: "",
		RedisConnectionString:   "",
		SmtpHost:                "",
		SmtpPort:                534,
		SmtpPassword:            "",
		SmtpFrom:                "",
	}

	env := os.Getenv("ENVIRONMENT")

	if env == "dev" {
		appConfig.Environment = EnvironmentDev
	}

	appConfig.MongodbConnectionString = os.Getenv("MONGODB_CONNECTION_STRING")
	appConfig.MongodbDatabaseName = os.Getenv("MONGODB_DATABASE_NAME")
	appConfig.RedisConnectionString = os.Getenv("REDIS_CONNECTION_STRING")
	appConfig.MongodbConnectionString = os.Getenv("MONGODB_CONNECTION_STRING")
	appConfig.MongodbDatabaseName = os.Getenv("MONGODB_DATABASE_NAME")
	appConfig.RedisConnectionString = os.Getenv("REDIS_CONNECTION_STRING")
	appConfig.SmtpHost = os.Getenv("SMTP_HOST")

	if port := os.Getenv("SMTP_PORT"); port != "" {
		var smtpPort uint16
		_, err := fmt.Sscan(port, &smtpPort)

		if err == nil {
			appConfig.SmtpPort = smtpPort
		} else {
			log.Printf("Failed to parse smtp port: %s", port)
		}
	}

	appConfig.SmtpUsername = os.Getenv("SMTP_USERNAME")
	appConfig.SmtpPassword = os.Getenv("SMTP_PASSWORD")
	appConfig.SmtpFrom = os.Getenv("SMTP_FROM")

	return appConfig
}
