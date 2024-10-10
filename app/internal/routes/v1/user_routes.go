package v1

import (
	"learning-golang/app/config"
	"learning-golang/app/internal/handlers"
	"learning-golang/app/internal/repository"
	"learning-golang/app/internal/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterUserRoutes(router *gin.RouterGroup, client *mongo.Client) {
	// Initialize repository, service, config, and handler
	cfg := config.LoadConfig()
	userRepo := repository.NewUserRepository(client, cfg.DBName)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router.POST("/", userHandler.CreateUser)
	router.GET("/all", userHandler.GetUsers)
	router.GET("/:id", userHandler.GetUserByID)
	router.PUT("/:id", userHandler.UpdateUser)
	router.DELETE("/:id", userHandler.DeleteUser)
}
