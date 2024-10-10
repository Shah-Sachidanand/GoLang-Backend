package v2

import (
	"learning-golang/config"
	"learning-golang/internal/handlers"
	"learning-golang/internal/repository"
	"learning-golang/internal/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterUserRoutes(router *gin.RouterGroup, client *mongo.Client) {
	// Initialize repository, service, config, and handler
	cfg := config.LoadConfig()
	userRepo := repository.NewUserRepository(client, cfg.DBName)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// @Summary      Create a new user
	// @Description  Create a new user with the provided information
	// @Tags         users
	// @Accept       json
	// @Produce      json
	// @Param        user  body      internal/models.User  true  "User information"
	// @Success      201   {object}  internal/models.User
	// @Failure      400   {object}  gin.H{"error": "Invalid input"}
	router.POST("/", userHandler.CreateUser)
	router.GET("/all", userHandler.GetUsers)
	router.GET("/:id", userHandler.GetUserByID)
	router.PUT("/:id", userHandler.UpdateUser)
	router.DELETE("/:id", userHandler.DeleteUser)
}
