package v2

import (
	"learning-golang/app/config"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.RouterGroup, client *config.Resource) {
	// User Routes Group
	userGroup := router.Group("/users")
	RegisterUserRoutes(userGroup, client)

	// Product Routes Group
	productGroup := router.Group("/products")
	RegisterProductRoutes(productGroup, client)
}
