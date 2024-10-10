package v2

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeRoutes(router *gin.RouterGroup, client *mongo.Client) {
	// User Routes Group
	userGroup := router.Group("/users")
	RegisterUserRoutes(userGroup, client)

	// Product Routes Group
	productGroup := router.Group("/products")
	RegisterProductRoutes(productGroup, client)
}
