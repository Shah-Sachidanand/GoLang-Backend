package v2

import (
	"learning-golang/app/config"
	"learning-golang/app/internal/handlers"
	"learning-golang/app/internal/repository"
	"learning-golang/app/internal/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterProductRoutes(router *gin.RouterGroup, client *mongo.Client) {
	// Initialize repository, service, config and handler
	cfg := config.LoadConfig()
	productRepo := repository.NewProductRepository(client, cfg.DBName)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	router.POST("/", productHandler.CreateProduct)
	router.GET("/", productHandler.GetProducts)

}
