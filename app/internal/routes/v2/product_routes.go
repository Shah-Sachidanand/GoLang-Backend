package v2

import (
	"learning-golang/app/config"
	"learning-golang/app/internal/handlers"
	"learning-golang/app/internal/repository"
	"learning-golang/app/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.RouterGroup, client *config.Resource) {
	productRepo := repository.NewProductRepository(client)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	router.POST("/", productHandler.CreateProduct)
	router.GET("/", productHandler.GetProducts)

}
