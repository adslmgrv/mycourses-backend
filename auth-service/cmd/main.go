package main

import (
	"context"
	"log"

	"github.com/adslmgrv/mycourses-backend/auth-service/internal/config"
	"github.com/adslmgrv/mycourses-backend/auth-service/internal/controller/v1"
	"github.com/adslmgrv/mycourses-backend/auth-service/internal/repository"
	"github.com/adslmgrv/mycourses-backend/auth-service/internal/service"
	"github.com/adslmgrv/mycourses-backend/common/pkg/database"
	"github.com/gin-gonic/gin"
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

	redis, err := database.NewRedisClient(ctx, config.RedisConnectionString)

	if err != nil {
		log.Printf("Failed to connect to Redis, cause: %s", err)
	}

	userRepository := repository.NewUserMongoRepository(mongodb)
	mfaRepository := repository.NewMFARedisRepository(redis)

	smtpEmailService := service.NewSmtpEmailService(config.SmtpHost, config.SmtpPort, config.SmtpUsername, config.SmtpPassword, config.SmtpFrom)

	authService := service.NewAuthService(userRepository, mfaRepository, smtpEmailService)

	r := gin.Default()
	authControllerV1 := v1.NewAuthController(authService)
	r.POST("/api/v1/auth/users", authControllerV1.SignUp)
	r.Run()
}
