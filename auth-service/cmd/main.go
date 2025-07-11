package main

import (
	"log"

	"github.com/adslmgrv/mycourses-backend/auth-service/internal/config"
	"github.com/adslmgrv/mycourses-backend/auth-service/internal/repo"
	"github.com/adslmgrv/mycourses-backend/common/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Failed to load .env file, using environmental variables instead")
	}

	config := config.LoadAppConfig()
	mongo, err := database.NewMongoClient(config.MongodbConnectionString)

	if err != nil {
		log.Printf("Failed to connect to MongoDB, cause: %s", err)
		return
	}

	_ = repo.NewMongoUserRepo(mongo)
}
