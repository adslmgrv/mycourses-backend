package main

import (
	"context"
	"log"

	"github.com/adslmgrv/mycourses-backend/auth-service/internal/config"
	"github.com/adslmgrv/mycourses-backend/auth-service/internal/repository"
	"github.com/adslmgrv/mycourses-backend/common/pkg/database"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Failed to load .env file, using environmental variables instead")
	}

	config := config.LoadAppConfig()
	mongo, err := database.NewMongoClient(ctx, config.MongodbConnectionString)

	if err != nil {
		log.Printf("Failed to connect to MongoDB, cause: %s", err)
		return
	}

	mongodb := mongo.Database(config.MongodbDatabaseName)

	_ = repository.NewUserMongoRepository(mongodb)
}
